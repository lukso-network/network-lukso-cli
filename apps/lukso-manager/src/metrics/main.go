package metrics

import (
	"io/ioutil"
	"log"
	"net/http"
)

func VanguardMetrics(w http.ResponseWriter, r *http.Request) {
	getMetrics("http://127.0.0.1:8080/metrics", w)
}

func PandoraMetrics(w http.ResponseWriter, r *http.Request) {
	getMetrics("http://127.0.0.1:6060/debug/metrics", w)
}

func ValidatorMetrics(w http.ResponseWriter, r *http.Request) {
	getMetrics("http://127.0.0.1:8081/metrics", w)
}

func getMetrics(url string, w http.ResponseWriter) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
