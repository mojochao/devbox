package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mojochao/devbox/internal/devbox"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Aliases: []string{"ls"},
	Use:     "list",
	Short:   "List devboxes in state",
	Long: `Multiple devboxes may be managed in application state.

This command displays their ids and devbox data.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Ensure correct usage.
		if len(args) > 0 {
			exit(1, "no arguments allowed")
		}

		// Load state.
		state, err := devbox.LoadState(stateFile)
		exitOnError(err, 1, fmt.Sprintf("cannot load state from %s", stateFile))

		// Display devboxes table.
		printBoxesTable(state.Boxes)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
