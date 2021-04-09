package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mojochao/devbox/internal/devbox"
)

// contextCmd represents the context command
var contextCmd = &cobra.Command{
	Aliases: []string{"ctx"},
	Use:     "context [ID]",
	Short:   "Get or set active devbox ID context",
	Long: `An active devbox ID context can be set to reduce the need to provide it to
commands requiring them. 

If an ID argument is provided the current active devbox context will be set to
it.

If an ID argument is not provided the current active devbox ID context will be
displayed.

If the global --verbose flag is provided, full details on the active devbox ID
will be displayed.

If the local --reset flag is provided, the current active devbox context will
be reset.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Ensure correct usage.
		if len(args) > 1 {
			exit(1, "only one ID argument may be provided")
		}

		// Load state.
		state, err := devbox.LoadState(stateFile)
		exitOnError(err, 1, fmt.Sprintf("cannot load state from %s", stateFile))

		if len(args) == 0 {
			// Handle reset context case.
			reset, _ := cmd.Flags().GetBool("reset")
			if reset {
				state.Active = ""
				err = state.Save()
				exitOnError(err, 1, fmt.Sprintf("cannot save state to %s", stateFile))
				return
			}

			// Handle get context case. If not verbose, print only the devbox id.
			if !verbose {
				if state.Active != "" {
					fmt.Println(state.Active)
				}
				return
			}

			// Otherwise, load the devbox and print all its info.
			box, err := state.GetDevbox(state.Active)
			exitOnError(err, 1, fmt.Sprintf("devbox %s not found", state.Active))
			boxes := devbox.Boxes{state.Active: box}
			printBoxesTable(boxes)
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
	contextCmd.Flags().BoolP("reset", "r", false, "Reset context")
}
