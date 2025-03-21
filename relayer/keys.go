package relayer

import (
	cosmossdkmath "cosmossdk.io/math"

	"github.com/dymensionxyz/roller/cmd/consts"
	"github.com/dymensionxyz/roller/utils/keys"
)

var oneDayRelayPrice, _ = cosmossdkmath.NewIntFromString(
	"2000000000000000000",
) // 2000000000000000000 = 2dym

func VerifyRelayerBalances(home string, hd consts.HubData) error {
	insufficientBalances, err := getRelayerInsufficientBalances(home, hd)
	if err != nil {
		return err
	}

	if len(insufficientBalances) != 0 {
		err = keys.PrintInsufficientBalancesIfAny(insufficientBalances)
		if err != nil {
			return err
		}
	}

	return nil
}

func getRelayerInsufficientBalances(
	home string,
	hd consts.HubData,
) ([]keys.NotFundedAddressData, error) {
	var insufficientBalances []keys.NotFundedAddressData

	accData, err := GetRelayerAccountsData(home, hd)
	if err != nil {
		return nil, err
	}

	// consts.Denoms.Hub is used here because as of @202409 we no longer require rollapp
	// relayer account funding to establish IBC connection.
	for _, acc := range accData {
		if acc.Balance.Amount.IsNegative() {
			insufficientBalances = append(
				insufficientBalances, keys.NotFundedAddressData{
					KeyName:         consts.KeysIds.HubRelayer,
					Address:         acc.Address,
					CurrentBalance:  acc.Balance.Amount.BigInt(),
					RequiredBalance: oneDayRelayPrice.BigInt(),
					Denom:           consts.Denoms.Hub,
					Network:         hd.ID,
				},
			)
		}
	}

	return insufficientBalances, nil
}

func GetRelayerAccountsData(
	home string,
	hd consts.HubData,
) ([]keys.AccountData, error) {
	var data []keys.AccountData

	// rollappRlyAcc, err := getRolRlyAccData(cfg)
	// if err != nil {
	// 	return nil, err
	// }
	// data = append(data, *rollappRlyAcc)

	hubRlyAcc, err := getHubRlyAccData(home, hd)
	if err != nil {
		return nil, err
	}

	data = append(data, *hubRlyAcc)
	return data, nil
}

func getHubRlyAccData(home string, hd consts.HubData) (*keys.AccountData, error) {
	HubRlyAddr, err := keys.GetRelayerAddress(home, hd.ID)
	if err != nil {
		return nil, err
	}

	HubRlyBalance, err := keys.QueryBalance(
		keys.ChainQueryConfig{
			RPC:    hd.RpcUrl,
			Denom:  consts.Denoms.Hub,
			Binary: consts.Executables.Dymension,
		}, HubRlyAddr,
	)
	if err != nil {
		return nil, err
	}

	return &keys.AccountData{
		Address: HubRlyAddr,
		Balance: *HubRlyBalance,
	}, nil
}
