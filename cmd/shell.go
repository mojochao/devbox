package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/mojochao/devbox/internal/devbox"
)

// shellCmd represents the shell command
var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "Open interactive shell in devbox",
	Long: `A devbox does nothing unless you use it. This is done by opening a
shell session to a devbox that has been previously started.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Load state.
		state, err := devbox.LoadState(stateFile)
		exitOnError(err, 1, "cannot load boxes")

		// Ensure we have a devbox id.
		id, _ := cmd.Flags().GetString("id")
		id = ensureDevboxID(state, id)

		// Load devbox by id.
		devbox, err := state.GetDevbox(id)
		exitOnError(err, 1, fmt.Sprintf("devbox %s not found", id))

		// Open shell on devbox.
		shell := devbox.Shell
		if len(args) > 0 {
			shell = strings.Join(args, " ")
		}
		err = devbox.OpenShell(shell)
		exitOnError(err, 1, "cannot open shell")
	},
}

func init() {
	rootCmd.AddCommand(shellCmd)
	shellCmd.Flags().StringP("id", "i", "", "Devbox id")
}
