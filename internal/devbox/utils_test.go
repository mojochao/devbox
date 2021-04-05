package devbox

import (
	"reflect"
	"testing"
)

func Test_getCurrentUsername(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test happy path",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCurrentUsername(); got == "" {
				t.Error("getCurrentUsername() = nothing, want something")
			}
		})
	}
}

func Test_getCommandAndArgs(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 []string
	}{
		{
			name: "test happy path with args",
			args: args{s: "echo this is a test"},
			want: "echo",
			want1: []string{"this", "is", "a", "test"},
		},
		{
			name: "test happy path with no args",
			args: args{s: "ls"},
			want: "ls",
			want1: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getCommandAndArgs(tt.args.s)
			if got != tt.want {
				t.Errorf("getCommandAndArgs() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("getCommandAndArgs() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_execCommand(t *testing.T) {
	type args struct {
		command string
		message string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test happy path",
			args: args{
				command: "echo this is a test",
				message: "this is a passing test",
			},
			wantErr: false,
		},
		{
			name: "test crappy path",
			args: args{
				command: "ech this is a test",
				message: "this is a failing test",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := execCommand(tt.args.command, tt.args.message); (err != nil) != tt.wantErr {
				t.Errorf("execCommand() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

