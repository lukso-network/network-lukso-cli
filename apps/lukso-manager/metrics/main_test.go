package metrics

import (
	"lukso/apps/lukso-manager/shared"
	"net/http"
	"reflect"
	"testing"

	dto "github.com/prometheus/client_model/go"
)

func TestVanguardMetrics(t *testing.T) {
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
			VanguardMetrics(tt.args.w, tt.args.r)
		})
	}
}

func TestPandoraMetrics(t *testing.T) {
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
			PandoraMetrics(tt.args.w, tt.args.r)
		})
	}
}

func TestValidatorMetrics(t *testing.T) {
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
			ValidatorMetrics(tt.args.w, tt.args.r)
		})
	}
}

func TestHealth(t *testing.T) {
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
			Health(tt.args.w, tt.args.r)
		})
	}
}

func Test_getMetrics(t *testing.T) {
	type args struct {
		url string
		w   http.ResponseWriter
	}
	tests := []struct {
		name     string
		args     args
		wantBody []byte
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBody, err := getMetrics(tt.args.url, tt.args.w)
			if (err != nil) != tt.wantErr {
				t.Errorf("getMetrics() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBody, tt.wantBody) {
				t.Errorf("getMetrics() = %v, want %v", gotBody, tt.wantBody)
			}
		})
	}
}

func TestGetPandoraPeersOverTime(t *testing.T) {
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
			GetPandoraPeersOverTime(tt.args.w, tt.args.r)
		})
	}
}

func TestGetVanguardPeersOverTime(t *testing.T) {
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
			GetVanguardPeersOverTime(tt.args.w, tt.args.r)
		})
	}
}

func Test_handleError(t *testing.T) {
	type args struct {
		err error
		w   http.ResponseWriter
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shared.HandleError(tt.args.err, tt.args.w)
		})
	}
}

func Test_returnBody(t *testing.T) {
	type args struct {
		body []byte
		w    http.ResponseWriter
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			returnBody(tt.args.body, tt.args.w)
		})
	}
}

func Test_decodeSettings(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name        string
		args        args
		wantMetrics map[int64]float64
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMetrics, err := decodeSettings(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("decodeSettings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotMetrics, tt.wantMetrics) {
				t.Errorf("decodeSettings() = %v, want %v", gotMetrics, tt.wantMetrics)
			}
		})
	}
}

func Test_parseMetricFamily(t *testing.T) {
	type args struct {
		text []byte
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]*dto.MetricFamily
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseMetricFamily(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseMetricFamily() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseMetricFamily() = %v, want %v", got, tt.want)
			}
		})
	}
}
