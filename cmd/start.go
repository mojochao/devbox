package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mojochao/devbox-cli/internal/devbox"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start devbox",
	Long: `A devbox needs to be started before it can be used. This command starts the
devbox in a Docker container or a Kubernetes pod.

Once started, the devbox can be used by opening a shell session with the shell
command`,
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

		// Start devbox.
		err = devbox.Start()
		exitOnError(err, 1, fmt.Sprintf("cannot start devbox %s", id))

		// Success!
		fmt.Println(fmt.Sprintf("devbox %s started", id))
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringP("id", "i", "", "Devbox id")
}
