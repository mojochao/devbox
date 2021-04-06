// Package devbox provides state management functionality to the Cobra CLI cmds.
package devbox

import (
	"fmt"
)

// Config contains configuration data required of a Devbox.
type Config struct {
	// Name of Docker container or Kubernetes pod running devbox image.
	Name string

	// Description of devbox.
	Description string

	// Image of devbox running in Docker container or Kubernetes pod.
	Image string

	// Shell to exec in devbox running in Docker container or Kubernetes pod.
	Shell string

	// Kubeconfig is path to kubeconfig of Kubernetes cluster hosting devbox pod.
	Kubeconfig string

	// Namespace is namespace of Kubernetes cluster hosting devbox pod.
	Namespace string
}

// DefaultConfig is a Config containing default configuration values.
var DefaultConfig = Config{
	Name:        fmt.Sprintf("devbox-%s", getCurrentUsername()),
	Description: "",
	Image:       "github.com/mojochao/devbox-base",
	Shell:       "sh",
	Kubeconfig:  "",
	Namespace:   "",
}

// Devbox contains information on a devbox.
type Devbox struct {
	// Name of Docker container or Kubernetes pod running devbox image.
	Name string `yaml:"name"`

	// Description of devbox.
	Description string `yaml:"description"`

	// Image of devbox running in Docker container or Kubernetes pod.
	Image string `yaml:"image"`

	// Shell to exec in devbox running in Docker container or Kubernetes pod.
	Shell string `yaml:"shell"`

	// Kubeconfig is path to kubeconfig of Kubernetes cluster hosting devbox pod.
	Kubeconfig string `yaml:"kubeconfig"`

	// Namespace is namespace of Kubernetes cluster hosting devbox pod.
	Namespace string `yaml:"namespace"`
}

// Start starts a Devbox.
func (box Devbox) Start() error {
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

// Stop stops a Devbox.
func (box Devbox) Stop() error {
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

// OpenShell opens a shell in a Devbox.
func (box Devbox) OpenShell(shellPath string) error {
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

// CopyFile copies a file to a Devbox.
func (box Devbox) CopyFile(src string, dst string) error {
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

// New returns a fully constructed Devbox.
func New(cfg *Config) Devbox {
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
	return Devbox{
		Name:        cfg.Name,
		Description: cfg.Description,
		Image:       cfg.Image,
		Shell:       cfg.Shell,
		Kubeconfig:  cfg.Kubeconfig,
		Namespace:   cfg.Namespace,
	}
}
