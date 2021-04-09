package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mojochao/devbox/internal/devbox"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add ID [flags]",
	Short: "Add devbox to state",
	Long:  `Devboxes must be added before they can be started and used.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Ensure correct usage.
		if len(args) < 1 {
			exit(1, "missing ID argument")
		}
		if len(args) > 1 {
			exit(1, "extra arguments found")
		}
		id := args[0]
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")
		image, _ := cmd.Flags().GetString("image")
		if image == "" {
			exit(1, "missing --image flag")
		}
		user, _ := cmd.Flags().GetString("user")
		shell, _ := cmd.Flags().GetString("shell")
		kubeconfig, _ := cmd.Flags().GetString("kubeconfig")
		namespace, _ := cmd.Flags().GetString("namespace")

		// Load state.
		state, err := devbox.LoadState(stateFile)
		exitOnError(err, 1, fmt.Sprintf("cannot load state from %s", stateFile))

		// AddDevbox devbox to state.
		box := devbox.New(&devbox.Config{
			Image:       image,
			User:        user,
			Shell:       shell,
			Name:        name,
			Namespace:   namespace,
			Kubeconfig:  kubeconfig,
			Description: description,
		})
		err = state.AddDevbox(id, box)
		exitOnError(err, 1, fmt.Sprintf("cannot add devbox %s", id))

		// Success!
		fmt.Printf("added devbox %s\n", id)
		fmt.Printf("set active devbox context to %s\n", id)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("image", "i", "", "Box docker image")
	addCmd.Flags().StringP("user", "u", "developer", "Box user name")
	addCmd.Flags().StringP("shell", "s", "zsh", "Box shell name or path")
	addCmd.Flags().StringP("name", "", "", "Box container or pod name")
	addCmd.Flags().StringP("namespace", "n", "", "Box pod namespace (Kubernetes devboxes only)")
	addCmd.Flags().StringP("kubeconfig", "k", "", "Box cluster kubeconfig (Kubernetes devboxes only)")
	addCmd.Flags().StringP("description", "d", "", "Box description")
}
