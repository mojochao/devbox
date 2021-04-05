package cmd

import (
	"fmt"
	"github.com/spf13/cobra"

	"github.com/mojochao/devbox-cli/internal/devbox"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize devbox state file",
	Long: `Data on managed devboxes is persisted in a state file. This command
creates that file.

By default, its path is ~/.devbox.state.yaml, but this can be overridden with the
global --state flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Ensure correct usage. An existing boxes file can only be re-initialized
		// with the local --force flag.
		var forceInit bool
		if fileExists(stateFile) {
			force, _ := cmd.Flags().GetBool("force")
			if !force {
				exit(1, fmt.Sprintf("state already initialized in %s", stateFile))
			}
			forceInit = true
		}

		// Create new state and save it to disk.
		state := devbox.NewState(stateFile)
		err := state.Save()
		exitOnError(err, 1, fmt.Sprintf("cannot save state to %s", stateFile))

		// Success!
		msg := fmt.Sprintf("initialized state in %s", stateFile)
		if forceInit {
			msg = "re-" + msg
		}
		fmt.Println(msg)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().BoolP("force", "f", false, "Force initialize existing state file")
}
