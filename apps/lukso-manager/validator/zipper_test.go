package validator

import "testing"

func Test_zipFolder(t *testing.T) {
	type args struct {
		network string
		folder  string
	}
	tests := []struct {
		name         string
		args         args
		wantFilePath string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFilePath := zipFolder(tt.args.network, tt.args.folder); gotFilePath != tt.wantFilePath {
				t.Errorf("zipFolder() = %v, want %v", gotFilePath, tt.wantFilePath)
			}
		})
	}
}
