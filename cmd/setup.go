package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mojochao/devbox/internal/devbox"
	"github.com/mojochao/devbox/internal/util"
)

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup [ID]",
	Short: "Setup devbox",
	Long: `Set ups a new started devbox with common developer configuration copied from
the local user's home directory by a manifest type.

If no --include flags are provided, all manifest types are included by default.
If any --include flags are provided, only those manifest types will be setup.

If any --exclude flags are provided, those manifest types will not be setup.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Load state.
		state, err := devbox.LoadState(stateFile)
		exitOnError(err, 1, fmt.Sprintf("cannot load state from %s", stateFile))

		// Set ids of devboxes to start.
		ids := args
		if len(ids) == 0 && state.Active != "" {
			ids = []string{state.Active}
		}
		for _, id := range ids {
			id = ensureDevboxID(state, id)
		}

		// Ensure valid includes and excludes.
		manifestTypes := devbox.ManifestTypes
		includes, _ := cmd.Flags().GetStringSlice("include")
		if len(includes) > 0 {
			for _, include := range includes {
				if !util.ContainsString(manifestTypes, include) {
					exit(1, fmt.Sprintf("invalid manifest type %s in --include flag", include))
				}
			}
		} else {
			includes = devbox.ManifestTypes
		}
		excludes, _ := cmd.Flags().GetStringSlice("exclude")

		// Setup devboxes.
		for _, id := range ids {
			box, err := state.GetDevbox(id)
			exitOnError(err, 1, fmt.Sprintf("devbox %s not found", id))

			for _, manifestType := range manifestTypes {
				if len(includes) > 0 && !util.ContainsString(includes, manifestType) {
					continue
				}
				if len(excludes) > 0 && util.ContainsString(excludes, manifestType) {
					continue
				}

				fmt.Printf("setting up devbox %s with %s config\n", id, manifestType)
				err = box.Setup(manifestType, includes, excludes)
				exitOnError(err, 1, fmt.Sprintf("cannot setup devbox %s with %s config", id, manifestType))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
	setupCmd.Flags().StringSliceP("include", "i", devbox.ManifestTypes, "Manifest types to include in setup")
	setupCmd.Flags().StringSliceP("exclude", "e", []string{}, "Manifest types to exclude in setup")
}
