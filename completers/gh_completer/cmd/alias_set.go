package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/rsteube/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var alias_setCmd = &cobra.Command{
	Use:   "set <alias> <expansion>",
	Short: "Create a shortcut for a gh command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(alias_setCmd).Standalone()

	alias_setCmd.Flags().BoolP("shell", "s", false, "Declare an alias to be passed through a shell interpreter")
	aliasCmd.AddCommand(alias_setCmd)

	carapace.Gen(alias_setCmd).PositionalCompletion(
		action.ActionAliases(),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if !alias_setCmd.Flag("shell").Changed {
				return bridge.ActionCarapaceBin("gh").Split()
			}
			return bridge.ActionCarapaceBin().SplitP()
		}),
	)
}
