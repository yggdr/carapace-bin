package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var taintCmd = &cobra.Command{
	Use:   "taint",
	Short: "Update the taints on one or more nodes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(taintCmd).Standalone()
	taintCmd.Flags().Bool("all", false, "Select all nodes in the cluster")
	taintCmd.Flags().Bool("allow-missing-template-keys", true, "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.")
	taintCmd.Flags().String("dry-run", "none", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.")
	taintCmd.Flags().String("field-manager", "kubectl-taint", "Name of the manager used to track field ownership.")
	taintCmd.Flags().StringP("output", "o", "", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).")
	taintCmd.Flags().Bool("overwrite", false, "If true, allow taints to be overwritten, otherwise reject taint updates that overwrite existing taints.")
	taintCmd.Flags().StringP("selector", "l", "", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2). Matching objects must satisfy all of the specified label constraints.")
	taintCmd.Flags().Bool("show-managed-fields", false, "If true, keep the managedFields when printing objects in JSON or YAML format.")
	taintCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].")
	taintCmd.Flags().String("validate", "strict", "Must be one of: strict (or true), warn, ignore (or false).")
	taintCmd.Flag("dry-run").NoOptDefVal = "unchanged"
	taintCmd.Flag("validate").NoOptDefVal = "strict"
	rootCmd.AddCommand(taintCmd)

	carapace.Gen(taintCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":  action.ActionDryRunModes(),
		"output":   action.ActionOutputFormats(),
		"template": carapace.ActionFiles(),
		"validate": kubectl.ActionValidationModes(),
	})

	carapace.Gen(taintCmd).PositionalCompletion(
		carapace.ActionValues("nodes"),
		action.ActionResources("", "nodes"),
	)
}