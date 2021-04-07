package cmd

import (
	"fmt"
	"github.com/spf13/cobra"

	"github.com/mojochao/devbox/internal/devbox"
)

// shellCmd represents the shell command
var shellCmd = &cobra.Command{
	Use:   "shell [SHELL]",
	Short: "Open interactive shell in devbox",
	Long: `A devbox does nothing unless you use it. This is done by opening a
shell session to a devbox that has been previously started.

If no SHELL argument is provided, the shell configured in the devbox will be
used defaulting to 'zsh' if not configured in the devbox.

If a SHELL argument is provided, that shell will be used instead of the shell
configuration in the devbox`,
	Run: func(cmd *cobra.Command, args []string) {
		// Ensure correct usage.
		if len(args) > 1 {
			exit(1, "only one SHELL argument allowed")
		}

		// Load state.
		state, err := devbox.LoadState(stateFile)
		exitOnError(err, 1, "cannot load boxes")

		// Ensure we have a devbox id.
		id, _ := cmd.Flags().GetString("id")
		id = ensureDevboxID(state, id)

		// Load devbox by id.
		box, err := state.GetDevbox(id)
		exitOnError(err, 1, fmt.Sprintf("devbox %s not found", id))

		// Open shell on devbox.
		shell := box.Shell
		if len(args) == 1 {
			shell = args[0]
		}
		err = box.OpenShell(shell)
		exitOnError(err, 1, "cannot open shell")
	},
}

func init() {
	rootCmd.AddCommand(shellCmd)
	shellCmd.Flags().StringP("id", "i", "", "Box id")
}
