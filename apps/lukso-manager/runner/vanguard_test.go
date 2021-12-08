package runner

import (
	"os/exec"
	"reflect"
	"testing"
)

func Test_startVanguard(t *testing.T) {
	type args struct {
		version   string
		network   string
		config    *NetworkConfig
		timestamp string
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
			gotCmd, err := startVanguard(tt.args.version, tt.args.network, tt.args.config, tt.args.timestamp)
			if (err != nil) != tt.wantErr {
				t.Errorf("startVanguard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCmd, tt.wantCmd) {
				t.Errorf("startVanguard() = %v, want %v", gotCmd, tt.wantCmd)
			}
		})
	}
}
