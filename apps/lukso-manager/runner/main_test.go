package runner

import (
	"net/http"
	"os/exec"
	"reflect"
	"testing"
)

func TestStartClients(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StartClients(tt.args.w, tt.args.r)
		})
	}
}

func TestStopClients(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StopClients(tt.args.w, tt.args.r)
		})
	}
}

func TestStartBinary(t *testing.T) {
	type args struct {
		client  string
		version string
		args    []string
	}
	tests := []struct {
		name    string
		args    args
		want    *exec.Cmd
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StartBinary(tt.args.client, tt.args.version, tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("StartBinary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StartBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
