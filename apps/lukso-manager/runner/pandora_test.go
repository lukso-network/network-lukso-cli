package runner

import (
	"lukso/apps/lukso-manager/settings"
	"os/exec"
	"reflect"
	"testing"
)

func Test_startPandora(t *testing.T) {
	type args struct {
		version   string
		network   string
		settings  settings.Settings
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
			gotCmd, err := startPandora(tt.args.version, tt.args.network, tt.args.settings, tt.args.config, tt.args.timestamp)
			if (err != nil) != tt.wantErr {
				t.Errorf("startPandora() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCmd, tt.wantCmd) {
				t.Errorf("startPandora() = %v, want %v", gotCmd, tt.wantCmd)
			}
		})
	}
}
