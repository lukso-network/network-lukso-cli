package shared

import "testing"

func TestGetDataDir(t *testing.T) {
	type args struct {
		network string
		client  string
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
			if got := GetDataDir(tt.args.network, tt.args.client); got != tt.want {
				t.Errorf("GetDataDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetNetworkDir(t *testing.T) {
	type args struct {
		network string
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
			if got := GetNetworkDir(tt.args.network); got != tt.want {
				t.Errorf("GetNetworkDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContains(t *testing.T) {
	type args struct {
		s []string
		e string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.s, tt.args.e); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
