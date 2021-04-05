package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"

	"github.com/mojochao/devbox-cli/internal/devbox"
)

// ensureDevboxID ensures that a devbox name to operate on is available
// from the active context in devboxes or in name parameter.
func ensureDevboxID(state devbox.State, name string) string {
	// Any devbox name provided on the command line takes priority over any
	// defined in the boxes file.
	if name == "" {
		name = state.Active
	}
	if name == "" {
		exit(1, "start requires a devbox name, either in boxes or with --name flag")
	}
	return name
}

// exit exits the application with an exit code and a message.
func exit(exitCode int, msg string) {
	if exitCode != 0 {
		msg = fmt.Sprintf("error: %s", msg)
	}
	fmt.Println(msg)
	os.Exit(exitCode)
}

// exitOnError exits the application on error with an exit code and a message.
func exitOnError(err error, exitCode int, msg string) {
	if err == nil {
		return
	}
	exit(exitCode, fmt.Sprintf("%s: %v", msg, err))
}

func fileExists(path string) bool {
	path, _ = homedir.Expand(path)
	info, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return !info.IsDir()
}
