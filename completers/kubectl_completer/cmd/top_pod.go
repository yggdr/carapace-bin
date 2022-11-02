package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var top_podCmd = &cobra.Command{
	Use:     "pod",
	Short:   "Display resource (CPU/memory) usage of pods",
	Aliases: []string{"pods", "po"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(top_podCmd).Standalone()
	top_podCmd.Flags().BoolP("all-namespaces", "A", false, "If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace.")
	top_podCmd.Flags().Bool("containers", false, "If present, print usage of containers within a pod.")
	top_podCmd.Flags().String("field-selector", "", "Selector (field query) to filter on, supports '=', '==', and '!='.(e.g. --field-selector key1=value1,key2=value2). The server only supports a limited number of field queries per type.")
	top_podCmd.Flags().Bool("no-headers", false, "If present, print output without headers.")
	top_podCmd.Flags().StringP("selector", "l", "", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2). Matching objects must satisfy all of the specified label constraints.")
	top_podCmd.Flags().String("sort-by", "", "If non-empty, sort pods list using specified field. The field can be either 'cpu' or 'memory'.")
	top_podCmd.Flags().Bool("sum", false, "Print the sum of the resource usage")
	top_podCmd.Flags().Bool("use-protocol-buffers", true, "Enables using protocol-buffers to access Metrics API.")
	topCmd.AddCommand(top_podCmd)

	carapace.Gen(top_podCmd).FlagCompletion(carapace.ActionMap{
		"sort-by": carapace.ActionValues("cpu", "memory"),
	})

	carapace.Gen(top_podCmd).PositionalCompletion(
		action.ActionResources("", "pods"),
	)
}