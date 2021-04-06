package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mojochao/devbox/internal/devbox"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop [ID...]",
	Short: "Stop devboxes",
	Long: `Once a devbox is no longer needed, it should be stopped. This command stops
devboxes in Docker containers or a Kubernetes pods.

If no ID arguments are provided, any set in the active devbox context will be used.

Once stopped, any files copied over to that devbox will be lost.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Load state.
		state, err := devbox.LoadState(stateFile)
		exitOnError(err, 1, fmt.Sprintf("cannot load state from %s", stateFile))

		// Set ids of devboxes to stop.
		if len(args) == 0 && state.Active != "" {
			args = []string{state.Active}
		}
		for _, id := range args {
			id = ensureDevboxID(state, id)
		}

		// Stop devboxes.
		for _, id := range args {
			box, err := state.GetDevbox(id)
			exitOnError(err, 1, fmt.Sprintf("devbox %s not found", id))

			err = box.Stop()
			exitOnError(err, 1, fmt.Sprintf("cannot stop devbox %s", id))

			fmt.Println(fmt.Sprintf("devbox %s stopped", id))
		}
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
