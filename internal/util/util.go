package util

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/mitchellh/go-homedir"

	"github.com/mojochao/devbox/internal/config"
)

// ContainsString tests if a string is in items.
func ContainsString(items []string, s string) bool {
	for _, item := range items {
		if item == s {
			return true
		}
	}
	return false
}

// ExecCommand executes a command.
func ExecCommand(name string, args ...string) error {
	if config.DryRun || config.Verbose {
		fmt.Printf("cmd: %s %s\n", name, strings.Join(args, " "))
		if config.DryRun {
			return nil
		}
	}

	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// FileExists tests if path is a file.
func FileExists(path string) bool {
	path, _ = homedir.Expand(path)
	info, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return !info.IsDir()
}

// DirExists tests if path is a directory.
func DirExists(path string) bool {
	path, _ = homedir.Expand(path)
	info, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return info.IsDir()
}
