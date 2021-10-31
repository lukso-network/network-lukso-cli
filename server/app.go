package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/fs"
	"log"
	"net/http"
)

//go:embed dist/apps/lukso-status
var static embed.FS

type App struct {
	r  *mux.Router
}

func (a *App) start() {
	webapp, err := fs.Sub(static, "dist/apps/lukso-status")
	if err != nil {
		fmt.Println(err)
	}
	a.r.PathPrefix("/").Handler(http.FileServer(http.FS(webapp)))
	log.Fatal(http.ListenAndServe(":8111", a.r))
}

func sendErr(w http.ResponseWriter, code int, message string) {
	resp, _ := json.Marshal(map[string]string{"error": message})
	http.Error(w, string(resp), code)
}