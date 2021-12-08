package webserver

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type App struct {
	Router *mux.Router
}

func (app *App) Start() {
	handler := cors.Default().Handler(app.Router)

	log.Fatal(http.ListenAndServe(":3000", handler))
}
