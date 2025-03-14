package rollapp

import (
	"github.com/spf13/cobra"

	"github.com/dymensionxyz/roller/cmd/rollapp/config"
	"github.com/dymensionxyz/roller/cmd/rollapp/drs"
	initrollapp "github.com/dymensionxyz/roller/cmd/rollapp/init"
	"github.com/dymensionxyz/roller/cmd/rollapp/keys"
	"github.com/dymensionxyz/roller/cmd/rollapp/migrate"
	"github.com/dymensionxyz/roller/cmd/rollapp/sequencer"
	"github.com/dymensionxyz/roller/cmd/rollapp/setup"
	"github.com/dymensionxyz/roller/cmd/rollapp/snapshot"
	"github.com/dymensionxyz/roller/cmd/rollapp/start"
	"github.com/dymensionxyz/roller/cmd/rollapp/status"
	"github.com/dymensionxyz/roller/cmd/services"
	loadservices "github.com/dymensionxyz/roller/cmd/services/load"
	restartservices "github.com/dymensionxyz/roller/cmd/services/restart"
	startservices "github.com/dymensionxyz/roller/cmd/services/start"
	stopservices "github.com/dymensionxyz/roller/cmd/services/stop"
)

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rollapp [command]",
		Short: "Commands to initialize and run a RollApp",
	}

	cmd.AddCommand(initrollapp.Cmd())
	cmd.AddCommand(status.Cmd())
	cmd.AddCommand(start.Cmd())
	cmd.AddCommand(config.Cmd())
	cmd.AddCommand(setup.Cmd())
	cmd.AddCommand(sequencer.Cmd())
	cmd.AddCommand(keys.Cmd())
	cmd.AddCommand(migrate.Cmd())
	cmd.AddCommand(drs.Cmd())
	cmd.AddCommand(snapshot.Cmd())

	sl := []string{"rollapp"}
	cmd.AddCommand(
		services.Cmd(
			loadservices.Cmd(sl, "rollapp"),
			startservices.RollappCmd(),
			restartservices.Cmd(sl),
			stopservices.Cmd(sl),
		),
	)

	return cmd
}
