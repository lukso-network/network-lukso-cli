package webserver

import (
	"testing"

	"github.com/gorilla/mux"
)

func TestApp_Start(t *testing.T) {
	type fields struct {
		Router *mux.Router
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &App{
				Router: tt.fields.Router,
			}
			app.Start(":3000")
		})
	}
}
