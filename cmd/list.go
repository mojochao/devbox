package cmd

import (
	"fmt"

	"github.com/rodaine/table"
	"github.com/spf13/cobra"

	"github.com/mojochao/devbox-cli/internal/devbox"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Aliases: []string{"ls"},
	Use:     "list",
	Short:   "List available devboxes",
	Long: `Multiple devboxes may be managed. The list command displays their names and
descriptions.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Ensure correct usage.
		if len(args) > 0 {
			exit(1, "no arguments allowed")
		}

		// Load state.
		state, err := devbox.LoadState(stateFile)
		exitOnError(err, 1, fmt.Sprintf("cannot load state from %s", stateFile))

		// If no boxes in state, we're done.
		if len(state.Boxes) == 0 {
			return
		}

		// Display table of boxes in state.
		tbl := table.New("id", "name", "image", "shell", "kubeconfig", "namespace", "description")
		for id, devbox := range state.Boxes {
			tbl.AddRow(id, devbox.Name, devbox.Image, devbox.Shell, devbox.Kubeconfig, devbox.Namespace, devbox.Description)
		}
		tbl.Print()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
