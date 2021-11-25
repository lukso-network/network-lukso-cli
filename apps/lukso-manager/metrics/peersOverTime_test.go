package metrics

import (
	"reflect"
	"testing"
)

func Test_setPeersOverTime(t *testing.T) {
	type args struct {
		amountOfPeers float64
		client        string
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
			if err := setPeersOverTime(tt.args.amountOfPeers, tt.args.client); (err != nil) != tt.wantErr {
				t.Errorf("setPeersOverTime() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getPeersOverTime(t *testing.T) {
	type args struct {
		client string
	}
	tests := []struct {
		name    string
		args    args
		want    map[int64]float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getPeersOverTime(tt.args.client)
			if (err != nil) != tt.wantErr {
				t.Errorf("getPeersOverTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPeersOverTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
