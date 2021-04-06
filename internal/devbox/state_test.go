package devbox

import (
	"os"
	"reflect"
	"testing"
)

var state = State{
	Active: "minimal",
	Boxes: Boxes{
		"minimal": {
			Image: "example.com/image",
			Name:  "minimal",
		},
		"docker": {
			Name:        "docker",
			Description: "docker devbox",
			Image:       "example.com/image:latest",
			Shell:       "bash",
		},
		"minikube": {
			Name:        "minikube",
			Description: "minikube devbox",
			Image:       "example.com/image:1.0.0",
			Shell:       "zsh",
			Kubeconfig:  "~/.kube/minikube",
			Namespace:   "default",
		},
		"eks": {
			Name:        "eks",
			Description: "eks devbox",
			Image:       "example.com/image:1.0.0",
			Shell:       "zsh",
			Kubeconfig:  "~/.kube/eks",
			Namespace:   "devbox",
		},
	},
}

var saveStateFile = "save.state.yaml"
var testStateFile = "test.state.yaml"

func Test_loadState(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    State
		wantErr bool
	}{
		{
			name: "test happy path",
			args: args{path: testStateFile},
			want: state,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := loadState(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadState() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadState() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_saveState(t *testing.T) {
	type args struct {
		path  string
		state State
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test happy path",
			args: args{
				path:  saveStateFile,
				state: state,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := saveState(tt.args.path, tt.args.state); (err != nil) != tt.wantErr {
				t.Errorf("saveState() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
		testState, err := loadState(testStateFile)
		if err != nil {
			t.Errorf("failed to load %s", testStateFile)
		}
		saveState, err := loadState(saveStateFile)
		if err != nil {
			t.Errorf("failed to load %s", saveStateFile)
		}
		if !reflect.DeepEqual(testState, saveState) {
			t.Error("testState and saveState are not equal")
		}
		os.Remove(saveStateFile)
	}
}

func TestState_Contains(t *testing.T) {
	type fields struct {
		Active    string
		Available Boxes
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "test happy path",
			fields: fields{
				Active:    state.Active,
				Available: state.Boxes,
			},
			args: args{name: "minimal"},
			want: true,
		},
		{
			name: "test crappy path",
			fields: fields{
				Active:    state.Active,
				Available: state.Boxes,
			},
			args: args{name: "nonesuch"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			state := State{
				Active: tt.fields.Active,
				Boxes:  tt.fields.Available,
			}
			if got := state.ContainsDevbox(tt.args.name); got != tt.want {
				t.Errorf("ContainsDevbox() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestState_Get(t *testing.T) {
	type fields struct {
		Active    string
		Available Boxes
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Box
		wantErr bool
	}{
		{
			name: "test happy path",
			fields: fields{
				Active:    state.Active,
				Available: state.Boxes,
			},
			args: args{name: "minimal"},
			want: state.Boxes["minimal"],
		},
		{
			name: "test crappy path",
			fields: fields{
				Active:    state.Active,
				Available: state.Boxes,
			},
			args:    args{name: "nonesuch"},
			want:    Box{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			state := State{
				Active: tt.fields.Active,
				Boxes:  tt.fields.Available,
			}
			got, err := state.GetDevbox(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDevbox() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDevbox() got = %v, want %v", got, tt.want)
			}
		})
	}
}
