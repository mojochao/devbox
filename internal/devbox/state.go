package devbox

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/ghodss/yaml"
	"github.com/mitchellh/go-homedir"

	"github.com/mojochao/devbox/internal/config"
)

// DefaultStateFile defines the default location of the boxes configuration file.
const DefaultStateFile = "~/.devbox.state.yaml"

// NewState returns a new State ready for use.
func NewState(path string) State {
	return State{
		Active: "",
		Boxes:  make(map[string]Devbox),
		Path:   path,
	}
}

// LoadState returns State loaded from path.
func LoadState(path string) (State, error) {
	if config.Verbose {
		fmt.Printf("loading state from %s\n", path)
	}
	state, err := loadState(path)
	if err != nil {
		return State{}, err
	}
	state.Path = path
	return state, nil
}

// State contains devbox state.
type State struct {
	// Active is the ID of the active devbox context, if any.
	Active string `yaml:"active"`

	// Boxes is the map of added Devbox items by ID.
	Boxes map[string]Devbox `yaml:"boxes"`

	// Path is the path to the state file.
	Path string `yaml:"path"`
}

// AddDevbox adds a Devbox to State.
func (boxes State) AddDevbox(id string, box Devbox) error {
	if boxes.ContainsDevbox(id) {
		return errors.New("devbox with id found")
	}
	boxes.Boxes[id] = box
	boxes.Active = id
	return saveState(boxes.Path, boxes)
}

// RemoveDevbox removes a Devbox from State.
func (boxes State) RemoveDevbox(id string) error {
	if !boxes.ContainsDevbox(id) {
		return errors.New("devbox with id not found")
	}
	delete(boxes.Boxes, id)
	if boxes.Active == id {
		boxes.Active = ""
	}
	return saveState(boxes.Path, boxes)
}

// ContainsDevbox tests if a Devbox is in State.
func (boxes State) ContainsDevbox(id string) bool {
	_, ok := boxes.Boxes[id]
	return ok
}

// GetDevbox returns Devbox in State.
func (boxes State) GetDevbox(id string) (Devbox, error) {
	box, ok := boxes.Boxes[id]
	if !ok {
		return Devbox{}, errors.New("devbox with id not found")
	}
	return box, nil
}

// Save saves State. If no paths are provided, the path from which State
// was loaded will be used.
func (boxes State) Save(paths ...string) error {
	if len(paths) == 0 {
		paths = append(paths, boxes.Path)
	}
	for _, path := range paths {
		boxes.Path = path
		if err := saveState(path, boxes); err != nil {
			return err
		}
		if config.Verbose {
			fmt.Printf("saved boxes to %s\n", path)
		}
	}
	return nil
}

func loadState(path string) (State, error) {
	var state State
	path, err := homedir.Expand(path)
	if err != nil {
		return state, err
	}
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return state, err
	}
	err = yaml.Unmarshal(buf, &state)
	return state, err
}

func saveState(path string, state State) error {
	path, err := homedir.Expand(path)
	if err != nil {
		return err
	}
	buf, err := yaml.Marshal(state)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, buf, 0644)
}
