package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mojochao/devbox-cli/internal/build"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display version",
	Long: `Display application version.

If the --verbose flag is provided, full build metadata including the git
branch and commit, and the build timestamp will be displayed.
`,
	Run: func(cmd *cobra.Command, args []string) {
		if !verbose {
			fmt.Println(build.Version)
		}
		fmt.Printf("rt-cid-api %s (branch=%s, commit=%s, date=%s, state=%s, summary=%s)\n",
			build.Version, build.GitBranch, build.GitCommit, build.Date, build.GitState, build.GitSummary)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
