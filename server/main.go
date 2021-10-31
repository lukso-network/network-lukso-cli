package main

import (
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"log"
)

func vanguardMetrics(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:8080/metrics")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func pandoraMetrics(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:6060/debug/metrics")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func main() {
	app := App{
		r:  mux.NewRouter(),
	}

	app.r.
	Methods("GET").
	Path("/vanguard/metrics").
	HandlerFunc(vanguardMetrics);

	app.r.
	Methods("GET").
	Path("/pandora/debug/metrics").
	HandlerFunc(pandoraMetrics);

	app.start()
}