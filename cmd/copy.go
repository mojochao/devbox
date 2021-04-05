package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mojochao/devbox-cli/internal/devbox"
)

// copyCmd represents the copy command
var copyCmd = &cobra.Command{
	Aliases: []string{"cp"},
	Use:   "copy SRC DST",
	Short: "Copy SRC files to devbox DST files",
	Long: `State begin life free of the custom config developers use daily to be their
most productive.

The copy command provides an easy way to copy files from the localhost to the
devbox container.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Ensure correct usage:
		if len(args) == 1 {
			exit(1, "missing DST argument")
		}
		if len(args) > 2 {
			exit(1, "extra arguments present")
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

		// Copy files to devbox.
		err = devbox.CopyFile(args[0], args[1])
		exitOnError(err, 1, fmt.Sprintf("cannot copy file to devbox %s", id))
	},
}

func init() {
	rootCmd.AddCommand(copyCmd)
	copyCmd.Flags().StringP("id", "i", "", "Devbox id")
}
