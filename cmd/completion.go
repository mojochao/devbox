package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish]",
	Short: "Generate completion script",
	Long: `To load completions:

Bash:

  $ source <(devbox completion bash)

  # To load completions for each session, execute once:
  # Linux:
  $ devbox completion bash > /etc/bash_completion.d/devbox
  # macOS:
  $ devbox completion bash > /usr/local/etc/bash_completion.d/devbox

Zsh:

  # If shell completion is not already enabled in your environment,
  # you will need to enable it.  You can execute the following once:

  $ echo "autoload -U compinit; compinit" >> ~/.zshrc

  # To load completions for each session, execute once:
  $ devbox completion zsh > "${fpath[1]}/_devbox"

  # You will need to start a new shell for this setup to take effect.

fish:

  $ devbox completion fish | source

  # To load completions for each session, execute once:
  $ devbox completion fish > ~/.config/fish/completions/devbox.fish
`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish"},
	Args:                  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			cmd.Root().GenFishCompletion(os.Stdout, true)
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
