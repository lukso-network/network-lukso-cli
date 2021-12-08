package runner

import (
	"os/exec"
	"reflect"
	"testing"
)

func Test_startOrchestrator(t *testing.T) {
	type args struct {
		version string
		network string
	}
	tests := []struct {
		name    string
		args    args
		wantCmd *exec.Cmd
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCmd, err := startOrchestrator(tt.args.version, tt.args.network)
			if (err != nil) != tt.wantErr {
				t.Errorf("startOrchestrator() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCmd, tt.wantCmd) {
				t.Errorf("startOrchestrator() = %v, want %v", gotCmd, tt.wantCmd)
			}
		})
	}
}
