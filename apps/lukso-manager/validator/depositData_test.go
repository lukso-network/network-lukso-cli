package validator

import (
	"reflect"
	"testing"
)

func TestReadDepositData(t *testing.T) {
	type args struct {
		network string
	}
	tests := []struct {
		name string
		args args
		want DepositData
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReadDepositData(tt.args.network); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadDepositData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readDepositJsonFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want DepositData
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readDepositJsonFile(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readDepositJsonFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
