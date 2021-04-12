// Package devbox provides state management functionality to the Cobra CLI cmds.
package devbox

import (
	"fmt"
	"os"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/mojochao/devbox/internal/util"
)

// Config contains configuration data required of a Box.
type Config struct {
	// Image of devbox running in Docker container or Kubernetes pod.
	Image string

	// User in docker container or Kubernetes pod running devbox image.
	User string

	// Shell to exec in devbox running in Docker container or Kubernetes pod.
	Shell string

	// Name of Docker container or Kubernetes pod running devbox image.
	Name string

	// Namespace is namespace of Kubernetes cluster hosting devbox pod.
	Namespace string

	// Kubeconfig is path to kubeconfig of Kubernetes cluster hosting devbox pod.
	Kubeconfig string

	// Description of devbox.
	Description string
}

// DefaultConfig is a Config containing default configuration values.
var DefaultConfig = Config{
	Image:       "github.com/mojochao/devbox-base",
	User:        "developer",
	Shell:       "sh",
	Name:        fmt.Sprintf("devbox-%s", getCurrentUsername()),
	Namespace:   "",
	Kubeconfig:  "",
	Description: "",
}

// Box contains information on a devbox.
type Box struct {
	// Image of devbox running in Docker container or Kubernetes pod.
	Image string `yaml:"image"`

	// User in docker container or Kubernetes pod running devbox image.
	User string `yaml:"user"`

	// Shell to exec in devbox running in Docker container or Kubernetes pod.
	Shell string `yaml:"shell"`

	// Name of Docker container or Kubernetes pod running devbox image.
	Name string `yaml:"name"`

	// Namespace is namespace of Kubernetes cluster hosting devbox pod.
	Namespace string `yaml:"namespace"`

	// Kubeconfig is path to kubeconfig of Kubernetes cluster hosting devbox pod.
	Kubeconfig string `yaml:"kubeconfig"`

	// Description of devbox.
	Description string `yaml:"description"`

	// Manifest of devbox.
	Manifest Manifest `yaml:"defaultManifest"`
}

// New returns a fully constructed Box.
func New(cfg *Config) Box {
	if cfg == nil {
		cfg = &DefaultConfig
	} else {
		if cfg.Image == "" {
			cfg.Image = DefaultConfig.Image
		}
		if cfg.User == "" {
			cfg.User = DefaultConfig.User
		}
		if cfg.Shell == "" {
			cfg.Shell = DefaultConfig.Shell
		}
		if cfg.Name == "" {
			cfg.Name = DefaultConfig.Name
		}
		if cfg.Namespace == "" {
			cfg.Namespace = DefaultConfig.Namespace
		}
		if cfg.Kubeconfig == "" {
			cfg.Kubeconfig = DefaultConfig.Kubeconfig
		}
		if cfg.Description == "" {
			cfg.Description = DefaultConfig.Description
		}
	}
	return Box{
		Image:       cfg.Image,
		User:        cfg.User,
		Shell:       cfg.Shell,
		Name:        cfg.Name,
		Namespace:   cfg.Namespace,
		Kubeconfig:  cfg.Kubeconfig,
		Description: cfg.Description,
		Manifest:    defaultManifest,
	}
}

// HomeDir returns the home directory of the devbox user.
func (box Box) HomeDir() string {
	return fmt.Sprintf("/home/%s", box.User)
}

// Start starts a Box.
func (box Box) Start() error {
	var command, message string
	if box.Namespace == "" {
		command = fmt.Sprintf("docker run --detach --name %s --rm --ulimit nofile=90000:90000 %s", box.Name, box.Image)
		message = fmt.Sprintf("starting devbox %s in docker", box.Name)
	} else {
		kubeconfig, _ := homedir.Expand(box.Kubeconfig)
		command = fmt.Sprintf("kubectl --kubeconfig %s run --image %s %s -n %s", kubeconfig, box.Image, box.Name, box.Namespace)
		message = fmt.Sprintf("starting devbox %s in cluster with %s kubeconfig", box.Name, box.Kubeconfig)
	}
	return execCommand(command, message)
}

// Setup sets up a Box.
func (box Box) Setup(manifestType string) error {
	items := defaultManifest[manifestType]
	for _, item := range items {
		if item.Path == "" {
			continue
		}
		if strings.HasSuffix(item.Path, "/") && !util.DirExists(item.Path) {
			continue
		}
		if !strings.HasSuffix(item.Path, "/") && !util.FileExists(item.Path) {
			continue
		}
		if err := box.copyPath(item.Path); err != nil {
			return err
		}

		namespace := ""
		if box.Namespace != "" {
			namespace = fmt.Sprintf("--namespace %s", box.Namespace)
		}

		for _, command := range item.Commands {
			if command == breakCommand {
				return nil
			}

			command = fmt.Sprintf("exec %s %s -- %s", namespace, box.Name, command)
			if err := box.execCommand(strings.Split(command, " ")); err != nil {
				return err
			}
		}
	}
	return nil
}

// Stop stops a Box.
func (box Box) Stop() error {
	var command, message string
	if box.Namespace == "" {
		command = fmt.Sprintf("docker stop %s", box.Name)
		message = fmt.Sprintf("stopping devbox %s in docker", box.Name)
	} else {
		kubeconfig, _ := homedir.Expand(box.Kubeconfig)
		command = fmt.Sprintf("kubectl --kubeconfig %s delete pod %s -n %s", kubeconfig, box.Name, box.Namespace)
		message = fmt.Sprintf("stopping devbox %s in %s namespace in cluster with %s kubeconfig", box.Name, box.Namespace, box.Kubeconfig)
	}
	return execCommand(command, message)
}

// OpenShell opens a shell in a Box.
func (box Box) OpenShell(shellPath string) error {
	if shellPath == "" {
		shellPath = box.Shell
	}
	var command, message string
	if box.Namespace == "" {
		command = fmt.Sprintf("docker exec -it %s %s", box.Name, shellPath)
		message = fmt.Sprintf("opening %s shell in devbox %s in docker", shellPath, box.Name)
	} else {
		kubeconfig, _ := homedir.Expand(box.Kubeconfig)
		command = fmt.Sprintf("kubectl --kubeconfig %s exec -it %s -n %s -- %s", kubeconfig, box.Name, box.Namespace, shellPath)
		message = fmt.Sprintf("opening %s shell in devbox %s in %s namespace in cluster with %s kubeconfig", box.Shell, box.Name, box.Namespace, box.Kubeconfig)
	}
	return execCommand(command, message)
}

// CopyFile copies a file to a Box.
func (box Box) CopyFile(src string, dst string) error {
	var command, message string
	if box.Namespace == "" {
		command = fmt.Sprintf("docker cp %s %s:%s", src, dst, box.Name)
		message = fmt.Sprintf("copying %s to %s in devbox %s in docker", src, dst, box.Name)
	} else {
		kubeconfig, _ := homedir.Expand(box.Kubeconfig)
		command = fmt.Sprintf("kubectl --kubeconfig %s cp  %s %s:%s -n %s", kubeconfig, src, box.Name, dst, box.Namespace)
		message = fmt.Sprintf("copying %s to %s in devbox %s in %s namespace in cluster with %s kubeconfig", src, dst, box.Name, box.Namespace, box.Kubeconfig)
	}
	return execCommand(command, message)
}

func (box Box) copyPath(path string) error {
	src, _ := homedir.Expand(path)
	src, err := os.Readlink(src)
	if err != nil {
		_, ok := err.(*os.PathError)
		if !ok {
			return err
		}
		src, _ = homedir.Expand(path)
	}

	dst := fmt.Sprintf("%s:%s", box.Name, strings.Replace(path, "~", box.HomeDir(), 1))
	args := []string{"cp", src, dst}
	return box.execCommand(args)
}

func (box Box) execCommand(args []string) error {
	for i, arg := range args {
		arg = strings.Replace(arg, "{box.User}", box.User, -1)
		arg = strings.Replace(arg, "{box.Shell}", box.Shell, -1)
		arg = strings.Replace(arg, "{box.Name}", box.Name, -1)
		arg = strings.Replace(arg, "{box.Namespace}", box.Namespace, -1)
		args[i] = arg
	}
	if box.Namespace == "" {
		return util.ExecCommand("docker", args...)
	}
	kubeconfig, _ := homedir.Expand(box.Kubeconfig)
	kubeArgs := []string{
		fmt.Sprintf("--kubeconfig=%s", kubeconfig),
		fmt.Sprintf("--namespace=%s", box.Namespace),
	}
	kubeArgs = append(kubeArgs, args...)
	return util.ExecCommand("kubectl", kubeArgs...)
}
