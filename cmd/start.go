package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mojochao/devbox/internal/devbox"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start [ID...]",
	Short: "Start devboxes",
	Long: `Devboxes need to be started before it can be used. This command starts devboxes
in Docker containers or Kubernetes pods.

If no ID arguments are provided, any set in the active devbox context will be used.

Once started, devboxes can be used by opening a shell session with the shell command.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Load state.
		state, err := devbox.LoadState(stateFile)
		exitOnError(err, 1, fmt.Sprintf("cannot load state from %s", stateFile))

		// Set ids of devboxes to start.
		if len(args) == 0 && state.Active != "" {
			args = []string{state.Active}
		}
		for _, id := range args {
			id = ensureDevboxID(state, id)
		}

		// Start devboxes.
		for _, id := range args {
			box, err := state.GetDevbox(id)
			exitOnError(err, 1, fmt.Sprintf("devbox %s not found", id))

			err = box.Start()
			exitOnError(err, 1, fmt.Sprintf("cannot start devbox %s", id))

			fmt.Println(fmt.Sprintf("devbox %s started", id))
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringP("id", "i", "", "Box id")
}
