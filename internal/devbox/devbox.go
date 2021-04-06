// Package devbox provides state management functionality to the Cobra CLI cmds.
package devbox

import (
	"fmt"
)

// Config contains configuration data required of a Box.
type Config struct {
	// Image of devbox running in Docker container or Kubernetes pod.
	Image string

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
}

// Start starts a Box.
func (box Box) Start() error {
	var command, message string
	if box.Namespace == "" {
		command = fmt.Sprintf("docker run --detach --name %s --rm %s", box.Name, box.Image)
		message = fmt.Sprintf("starting devbox %s in docker", box.Name)
	} else {
		command = fmt.Sprintf("kubectl --kubeconfig %s run --image %s %s -n %s", box.Kubeconfig, box.Image, box.Name, box.Namespace)
		message = fmt.Sprintf("starting devbox %s in cluster with %s kubeconfig", box.Name, box.Kubeconfig)
	}
	return execCommand(command, message)
}

// Stop stops a Box.
func (box Box) Stop() error {
	var command, message string
	if box.Namespace == "" {
		command = fmt.Sprintf("docker stop %s", box.Name)
		message = fmt.Sprintf("stopping devbox %s in docker", box.Name)
	} else {
		command = fmt.Sprintf("kubectl --kubeconfig %s delete pod %s -n %s", box.Kubeconfig, box.Name, box.Namespace)
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
		command = fmt.Sprintf("kubectl --kubeconfig %s exec -it %s -n %s -- %s", box.Kubeconfig, box.Name, box.Namespace, shellPath)
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
		command = fmt.Sprintf("kubectl --kubeconfig %s cp  %s %s:%s -n %s", box.Kubeconfig, src, box.Name, dst, box.Namespace)
		message = fmt.Sprintf("copying %s to %s in devbox %s in %s namespace in cluster with %s kubeconfig", src, dst, box.Name, box.Namespace, box.Kubeconfig)
	}
	return execCommand(command, message)
}

// New returns a fully constructed Box.
func New(cfg *Config) Box {
	if cfg == nil {
		cfg = &DefaultConfig
	} else {
		if cfg.Name == "" {
			cfg.Name = DefaultConfig.Name
		}
		if cfg.Description == "" {
			cfg.Description = DefaultConfig.Description
		}
		if cfg.Image == "" {
			cfg.Image = DefaultConfig.Image
		}
		if cfg.Shell == "" {
			cfg.Shell = DefaultConfig.Shell
		}
	}
	return Box{
		Image:       cfg.Image,
		Shell:       cfg.Shell,
		Name:        cfg.Name,
		Namespace:   cfg.Namespace,
		Kubeconfig:  cfg.Kubeconfig,
		Description: cfg.Description,
	}
}
