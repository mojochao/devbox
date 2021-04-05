package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit the state file",
	Long: `The devbox application stores its state in a state file. This command opens the
state file in the editor configured in the EDITOR environment variable, or the
value of the --editor or -e flags if that environment variable is not set.`,
	Run: func(cmd *cobra.Command, args []string) {
		editor, ok := os.LookupEnv("EDITOR")
		if !ok {
			editor, _ = cmd.Flags().GetString("editor")
		}

		stateFile, _ = homedir.Expand(stateFile)
		if verbose {
			fmt.Printf("opening state in %s\n for editing", stateFile)
		}

		editCmd := exec.Command(editor, stateFile)
		editCmd.Stdout = os.Stdout
		editCmd.Stdin = os.Stdin
		editCmd.Stderr = os.Stderr
		err := editCmd.Run()
		exitOnError(err, 1, fmt.Sprintf("cannot open state in %s\n in %s", stateFile, editor))
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
	editCmd.Flags().StringP("editor", "e", "vim", "Editor to use")
}
