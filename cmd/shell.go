package cmd

import (
	"fmt"
	"github.com/spf13/cobra"

	"github.com/mojochao/devbox/internal/devbox"
)

// shellCmd represents the shell command
var shellCmd = &cobra.Command{
	Aliases: []string{"sh"},
	Use:     "shell [ID]",
	Short:   "Open interactive shell in devbox",
	Long: `A devbox does nothing unless you use it. This is done by opening a
shell session to a devbox that has been previously started.

If no ID argument is provided, any set in the active devbox context will be
used.

If the --shell flag is provided, any set in the active devbox context will be
used`,
	Run: func(cmd *cobra.Command, args []string) {
		// Ensure correct usage.
		if len(args) > 1 {
			exit(1, "only one ID argument allowed")
		}

		// Load state.
		state, err := devbox.LoadState(stateFile)
		exitOnError(err, 1, "cannot load boxes")

		// Ensure we have a devbox id.
		id := state.Active
		if len(args) == 1 {
			id = args[0]
		}
		id = ensureDevboxID(state, id)

		// Load devbox by id.
		box, err := state.GetDevbox(id)
		exitOnError(err, 1, fmt.Sprintf("devbox %s not found", id))

		// Open shell on devbox.
		shell, _ := cmd.Flags().GetString("shell")
		err = box.OpenShell(shell)
		exitOnError(err, 1, "cannot open shell")
	},
}

func init() {
	rootCmd.AddCommand(shellCmd)
	shellCmd.Flags().StringP("shell", "s", "", "shell name or path")
}
