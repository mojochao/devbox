# DEVBOX

Interactively running and debugging in containers can be hell. For such an
environment, a development box running inside the container is useful.

This `devbox` CLI manages use of terminal-based development environment devboxes.
A devbox is defined in terms of:

- name of container or pod running the devbox image
- description of devbox usage
- image name of the devbox to run in a container or pod
- shell name or path to run in the container or pod
- kubeconfig of Kubernetes cluster to run devbox pods (optional, Kubernetes only)
- namespace of Kubernetes cluster to run devbox pods  (optional, Kubernetes only

Note that a devbox is intended to be a "pet" not "cattle", more persistent
than ephemeral.  Any files copied to the devbox will be lost once stopped.

This application provides the following functionality:

- managing devboxes with the `list`, `add` and `remove` commands
- operating devboxes with the `start`, `stop`, `shell` and `copy` commands
- providing version and other build metadata with the `version` command

This application persists its state in a state file, which by default is
`~/.devbox.state.yaml`.

This state includes:
- the current active devbox context by its ID
- the devboxes that have been added
- the file path to the state file

Application state may be queried for the current active context.

    devbox context

It can also be set if provided an ID of a managed devbox.

    devbox context demo-box

Application state may be queried for a list of the managed `boxes` as well.

    devbox list

## Installation

This application is easily installed with `go`.

    go install github.com/mojochao/devbox@latest

## Usage

To get details on usage run `devbox` with no commands, or the `-h` or `--help`
flags.

Application state must be initialized once before use.

    devbox init

After initialization, the list of managed devboxes can be displayed.

    devbox list

Initially, this will be empty. After adding devboxes, this will not be empty
and details of added devboxes will be displayed.

After initialization, a typical startup and provisioning workflow looks like
the following.

    devbox add
    devbox start
    # copy local files to devbox with 'docker cp' and 'kubectl cp' as desired

Once started and configuration, the devbox is used interactively in shells
running in the devbox container or pod until no longer needed.

    devbox shell

Once the started devbox is no longer needed, it should be stopped.

    devbox stop

Stopping a devbox removes all files copied to it.  If the devbox is restarted,
it will be necessary to recopy any files needed to the devbox.

Once the stopped devbox is no longer needed and likely never to be needed again,
it may be removed from devbox management.

    devbox remove

Once removed, devboxes will not show up in the output of the `devbox list`
command and cannot be started. If needed again, re-add the devbox with the
`devbox add` command and provision it again for use as desired.
