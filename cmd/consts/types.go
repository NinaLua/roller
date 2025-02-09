package consts

type HubData struct {
	Environment   string    `toml:"environment"     json:"environment"`
	ApiUrl        string    `toml:"api_url"         json:"apiUrl"`
	ID            string    `toml:"id"              json:"id"            yaml:"id"`
	RpcUrl        string    `toml:"rpc_url"         json:"rpcUrl"        yaml:"rpc_url"`
	ArchiveRpcUrl string    `toml:"archive_rpc_url" json:"archiveRpcUrl" yaml:"archive_rpc_url"`
	GasPrice      string    `toml:"gas_price"       json:"gasPrice"      yaml:"gas_price"`
	DaNetwork     DaNetwork `toml:"da_network"      json:"daNetwork"     yaml:"da_network"`
}

type RollappData = struct {
	ID       string `toml:"id"        yaml:"id"`
	ApiUrl   string `toml:"api_url"`
	RpcUrl   string `toml:"rpc_url"`
	GasPrice string `toml:"gas_price"`
	Denom    string `toml:"denom"`
}

type DaData = struct {
	Backend          DAType    `toml:"backend"`
	ID               DaNetwork `toml:"id"`
	ApiUrl           string    `toml:"api_url"`
	RpcUrl           string    `toml:"rpc_url"`
	CurrentStateNode string    `toml:"current_state_node"`
	StateNodes       []string  `toml:"state_nodes"`
	GasPrice         string    `toml:"gas_price"`
}
