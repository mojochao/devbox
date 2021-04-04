// Package cmd contains the Cobra commands making up the devbox CLI.
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/mojochao/devbox-cli/internal/config"
	"github.com/mojochao/devbox-cli/internal/devbox"
)

// These variables are set by persistent flags on the rootCmd.
var (
	dryRun    bool
	stateFile string
	verbose   bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "devbox",
	Short: "Manage and operate devboxes",
	Long: `Interactively running and debugging in containers can be hell. For such an
environment, a development host running inside the container is useful.

This application manages use of terminal-based development environment devboxes.
A devbox is defined in terms of:

- name of devbox container or pod running the devbox image
- description of devbox usage
- image name of the devbox to run in a container or pod
- shell name or path to run in the container or pod
- kubeconfig of Kubernetes cluster to run devbox pods (optional, Kubernetes only)
- namespace of Kubernetes cluster to run devbox pods  (optional, Kubernetes only

Note that a devbox is intended to be more "pet" than "cattle", more persistent
than ephemeral.  Any files copied to the devbox will be lost once stopped.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&stateFile, "state", devbox.DefaultStateFile, "state file")
	rootCmd.PersistentFlags().BoolVar(&dryRun, "dry-run", false, "preview commands")
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "show verbose output")
}

// initConfig sets the global config.
func initConfig() {
	config.StateFile = stateFile
	config.DryRun = dryRun
	config.Verbose = verbose
}
