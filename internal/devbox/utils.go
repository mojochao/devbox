package devbox

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"

	"github.com/mojochao/devbox/internal/config"
)

func execCommand(command string, message string) error {
	if config.DryRun || config.Verbose {
		fmt.Printf("cmd: %s\n", command)
	}
	if config.DryRun {
		return nil
	}
	if config.Verbose {
		fmt.Printf("msg: %s\n", message)
	}
	cmd, args := getCommandAndArgs(command)
	execCmd := exec.Command(cmd, args...)
	execCmd.Stdout = os.Stdout
	execCmd.Stdin = os.Stdin
	execCmd.Stderr = os.Stderr
	return execCmd.Run()
}

func getCommandAndArgs(s string) (string, []string) {
	parts := strings.Split(s, " ")
	return parts[0], parts[1:]
}

func getCurrentUsername() string {
	currentUser, err := user.Current()
	if err != nil {
		panic(err)
	}
	return currentUser.Username
}

