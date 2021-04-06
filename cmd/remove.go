package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mojochao/devbox/internal/devbox"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Aliases: []string{"rm"},
	Use:   "remove [ID...]",
	Short: "Remove devboxes from state",
	Long:  `Once a devbox is no longer needed it should be removed.

If no ID arguments are provided, any set in the active devbox context will be used.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Load state.
		state, err := devbox.LoadState(stateFile)
		exitOnError(err, 1, fmt.Sprintf("cannot load state from %s", stateFile))

		// Set ids of devboxes to remove.
		if len(args) == 0 && state.Active != "" {
			args = []string{state.Active}
		}
		for _, id := range args {
			id = ensureDevboxID(state, id)
		}

		// Remove devboxes from state.
		for _, id := range args {
			err = state.RemoveDevbox(id)
			exitOnError(err, 1, fmt.Sprintf("cannot remove devbox %s", id))

			// Success!
			fmt.Printf("removed devbox %s\n", id)
			if state.Active != "" {
				fmt.Println("reset active devbox context to nothing")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
