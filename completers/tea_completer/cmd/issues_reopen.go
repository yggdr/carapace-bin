package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/tea_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

var issues_reopenCmd = &cobra.Command{
	Use:     "reopen",
	Short:   "Change state of one or more issues to 'open'",
	Aliases: []string{"open"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issues_reopenCmd).Standalone()

	issues_reopenCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	issues_reopenCmd.Flags().StringP("output", "o", "", "Output format. (simple, table, csv, tsv, yaml, json)")
	issues_reopenCmd.Flags().StringP("remote", "R", "", "Discover Gitea login from remote. Optional")
	issues_reopenCmd.Flags().StringP("repo", "r", "", "Override local repository path or gitea repository slug to interact with. Optional")
	issuesCmd.AddCommand(issues_reopenCmd)

	// TODO completion
	carapace.Gen(issues_reopenCmd).FlagCompletion(carapace.ActionMap{
		"output": tea.ActionOutputFormats(),
		"remote": git.ActionRemotes(),
	})

	carapace.Gen(issues_reopenCmd).PositionalCompletion(
		action.ActionIssues(issues_reopenCmd, false),
	)
}
