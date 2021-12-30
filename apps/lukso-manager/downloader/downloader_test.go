package downloader

import (
	"net/http"
	"testing"
)

func Test_downloadFile(t *testing.T) {
	type args struct {
		filepath string
		url      string
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
			if err := downloadFile(tt.args.filepath, tt.args.url); (err != nil) != tt.wantErr {
				t.Errorf("downloadFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetDownloadedVersions(t *testing.T) {
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
			GetDownloadedVersions(tt.args.w, tt.args.r)
		})
	}
}

func TestGetAvailableVersions(t *testing.T) {
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
			GetAvailableVersionsEndpoint(tt.args.w, tt.args.r)
		})
	}
}

func Test_getDownloadUrlFromAsset(t *testing.T) {
	type args struct {
		name   string
		assets []Assets
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDownloadUrlFromAsset(tt.args.name, tt.args.assets); got != tt.want {
				t.Errorf("getDownloadUrlFromAsset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDownloadClient(t *testing.T) {
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
			DownloadClientEndpoint(tt.args.w, tt.args.r)
		})
	}
}

func TestDownloadClientBinary(t *testing.T) {
	type args struct {
		client      string
		tag_version string
		url         string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DownloadClientBinary(tt.args.client, tt.args.tag_version, tt.args.url)
		})
	}
}

func Test_createDirIfNotExists(t *testing.T) {
	type args struct {
		folder string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createDirIfNotExists(tt.args.folder)
		})
	}
}

func TestDownloadConfigFiles(t *testing.T) {
	type args struct {
		network string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DownloadConfigFiles(tt.args.network)
		})
	}
}
