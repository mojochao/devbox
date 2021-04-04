package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mojochao/devbox-cli/internal/devbox"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Aliases: []string{"rm"},
	Use:   "remove ID",
	Short: "RemoveDevbox devbox",
	Long:  `Once a devbox is no longer needed it should be removed. This command does that.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Ensure correct usage.
		if len(args) < 1 {
			exit(1, "missing ID argument")
		}
		if len(args) > 1 {
			exit(1, "extra arguments found")
		}
		id := args[0]

		// Load state.
		state, err := devbox.LoadState(stateFile)
		exitOnError(err, 1, fmt.Sprintf("cannot load state from %s", stateFile))

		// RemoveDevbox devbox from state.
		err = state.RemoveDevbox(id)
		exitOnError(err, 1, fmt.Sprintf("cannot remove devbox %s", id))

		// Success!
		fmt.Printf("removed devbox %s\n", id)
		if state.Active != "" {
			fmt.Println("reset active devbox context to nothing")
		}
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
