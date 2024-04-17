package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var flake_prefetchCmd = &cobra.Command{
	Use:   "prefetch [flags] [flake-url]",
	Short: "download the source tree denoted by a flake reference into the nix store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(flake_prefetchCmd).Standalone()

	flake_prefetchCmd.Flags().Bool("json", false, "Produce output in JSON format")

	addEvaluationFlags(flake_prefetchCmd)
	addFlakeFlags(flake_prefetchCmd)
	addLoggingFlags(flake_prefetchCmd)

	carapace.Gen(flake_prefetchCmd).PositionalCompletion(nix.ActionFlakes())

	flakeCmd.AddCommand(flake_prefetchCmd)
}
