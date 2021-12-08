package settings

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/boltdb/bolt"
)

func TestSaveSettingsEndpoint(t *testing.T) {
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
			SaveSettingsEndpoint(tt.args.w, tt.args.r)
		})
	}
}

func TestSaveSettings(t *testing.T) {
	type args struct {
		db       *bolt.DB
		settings *Settings
		network  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SaveSettings(tt.args.db, tt.args.settings, tt.args.network); (err != nil) != tt.wantErr {
				t.Errorf("SaveSettings() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetSettingsEndpoint(t *testing.T) {
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
			GetSettingsEndpoint(tt.args.w, tt.args.r)
		})
	}
}

func Test_decodeSettings(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *Settings
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := decodeSettings(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("decodeSettings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decodeSettings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSettings(t *testing.T) {
	type args struct {
		db      *bolt.DB
		network string
	}
	tests := []struct {
		name    string
		args    args
		want    *Settings
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getSettings(tt.args.db, tt.args.network)
			if (err != nil) != tt.wantErr {
				t.Errorf("getSettings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSettings() = %v, want %v", got, tt.want)
			}
		})
	}
}
