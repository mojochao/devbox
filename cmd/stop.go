package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mojochao/devbox-cli/internal/devbox"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop devbox",
	Long: `Once a devbox is no longer needed, it should be stopped. This command stops the
devbox in a Docker container or a Kubernetes pod.

Once stopped, any files copied over to it will be lost.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Ensure correct usage.
		if len(args) > 0 {
			exit(1, "no arguments allowed")
		}

		// Load state.
		state, err := devbox.LoadState(stateFile)
		exitOnError(err, 1, fmt.Sprintf("cannot load state from %s", stateFile))

		// Ensure we have a devbox id.
		id, _ := cmd.Flags().GetString("id")
		id = ensureDevboxID(state, id)

		// Load devbox by id.
		devbox, err := state.GetDevbox(id)
		exitOnError(err, 1, fmt.Sprintf("devbox %s not found", id))

		// Stop devbox.
		err = devbox.Stop()
		exitOnError(err, 1, fmt.Sprintf("cannot stop devbox %s", id))

		// Success!
		fmt.Println(fmt.Sprintf("devbox %s stopped", id))
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
	stopCmd.Flags().StringP("id", "i", "", "Devbox id")
}
