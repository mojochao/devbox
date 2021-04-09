package cmd

import (
	"fmt"
	"os"

	"github.com/rodaine/table"

	"github.com/mojochao/devbox/internal/devbox"
)

// ensureDevboxID ensures that a devbox ID to operate on is available
// from the active context in devboxes or in ID argument..
func ensureDevboxID(state devbox.State, id string) string {
	// Any devbox id provided on the command line takes priority over any
	// defined in the state.
	if id == "" {
		id = state.Active
	}
	if id == "" {
		exit(1, "missing devbox ID in arguments or active context")
	}
	return id
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

func printBoxesTable(boxes devbox.Boxes) {
	if len(boxes) == 0 {
		return
	}
	tbl := table.New("id", "image", "user", "shell", "name", "namespace", "kubeconfig", "description")
	for id, box := range boxes {
		tbl.AddRow(id, box.Image, box.User, box.Shell, box.Name, box.Namespace, box.Kubeconfig, box.Description)
	}
	tbl.Print()

}
