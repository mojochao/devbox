package cmd

import (
	"fmt"
	"github.com/rodaine/table"

	"github.com/spf13/cobra"

	"github.com/mojochao/devbox-cli/internal/devbox"
)

// contextCmd represents the context command
var contextCmd = &cobra.Command{
	Aliases: []string{"ctx"},
	Use:   "context [ID]",
	Short: "GetDevbox or set active devbox name context",
	Long: `An active devbox name context can be set to reduce the need to provide the
global --name flag to all commands. 

If an ID argument is not provided the current active context will be displayed.
If an ID argument is provided the current active context will be set to it.

If the global --verbose flag is provided, full details on the active devbox
will be displayed.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Ensure correct usage.
		if len(args) > 1 {
			exit(1, "only one ID argument may be provided")
		}

		// Load state.
		state, err := devbox.LoadState(stateFile)
		exitOnError(err, 1, fmt.Sprintf("cannot load state from %s", stateFile))

		// If no NAME arg provided and a current active context in state, just print it.
		if len(args) == 0 && state.Active != "" {
			// If not verbose, print only the devbox id.
			if !verbose {
				fmt.Println(state.Active)
				return
			}
			// Otherwise, load the devbox and print all its info.
			devbox, err := state.GetDevbox(state.Active)
			exitOnError(err, 1, fmt.Sprintf("devbox %s not found", state.Active))
			tbl := table.New("id", "name", "image", "shell", "kubeconfig", "namespace", "description")
			tbl.AddRow(state.Active, devbox.Name, devbox.Image, devbox.Shell, devbox.Kubeconfig, devbox.Namespace, devbox.Description)
			tbl.Print()
			return
		}

		// Otherwise, set the current active context in state and save it.
		// First ensure it exists.
		id := args[0]
		if !state.ContainsDevbox(id) {
			exit(1, fmt.Sprintf("devbox %s not found", state.Active))
		}
		state.Active = id
		err = state.Save()
		exitOnError(err, 1, fmt.Sprintf("cannot save state to %s", stateFile))
	},
}

func init() {
	rootCmd.AddCommand(contextCmd)
}
