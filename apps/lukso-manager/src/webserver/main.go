package webserver

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//go:embed dist/apps/lukso-status
var static embed.FS

type App struct {
	Router *mux.Router
}

func (a *App) Start() {
	webapp, err := fs.Sub(static, "dist/apps/lukso-status")
	if err != nil {
		fmt.Println(err)
	}
	a.Router.PathPrefix("/").Handler(http.FileServer(http.FS(webapp)))
	log.Fatal(http.ListenAndServe(":8111", a.Router))
}

func sendErr(w http.ResponseWriter, code int, message string) {
	resp, _ := json.Marshal(map[string]string{"error": message})
	http.Error(w, string(resp), code)
}