package metrics

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	dto "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"
)

type VanguardMetricss struct {
	Peers    int64 `json:"peers"`
	HeadSlot int64 `json:"headSlot"`
}

func VanguardMetrics(w http.ResponseWriter, r *http.Request) {
	body, err1 := getMetrics("http://127.0.0.1:8080/metrics", w)
	if err1 != nil {
		handleError(err1, w)
		return
	}

	mf, err2 := parseMetricFamily(body)

	if err2 != nil {
		handleError(err2, w)
		return
	}

	peers := mf["p2p_peer_count"].GetMetric()
	headSlot := mf["beacon_head_slot"].GetMetric()

	if peers == nil || headSlot == nil {
		return
	}

	// TODO: proper error handling in case the structure of the metrics changes
	var response VanguardMetricss = VanguardMetricss{
		Peers:    int64(*peers[1].Gauge.Value),
		HeadSlot: int64(*headSlot[0].Gauge.Value),
	}

	err := setPeersOverTime(*peers[1].Gauge.Value, "vanguard")
	if err != nil {
		fmt.Println(err)
	}

	jsonString, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
		handleError(err, w)
		return
	}

	returnBody(jsonString, w)
}

func PandoraMetrics(w http.ResponseWriter, r *http.Request) {
	body, err := getMetrics("http://127.0.0.1:6060/debug/metrics", w)
	if err != nil {
		return
	}

	var pandoraMetrics map[string]float64

	json.Unmarshal(body, &pandoraMetrics)

	err2 := setPeersOverTime(pandoraMetrics["p2p/peers"], "pandora")
	if err2 != nil {
		fmt.Println(err2)
	}

	if nil != err {
		fmt.Println(err)
		handleError(err, w)
		return
	}

	returnBody(body, w)
}

func ValidatorMetrics(w http.ResponseWriter, r *http.Request) {
	body, err := getMetrics("http://127.0.0.1:8081/metrics", w)
	if err != nil {
		handleError(err, w)
		return
	}
	returnBody(body, w)
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func getMetrics(url string, w http.ResponseWriter) (body []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		handleError(err, w)
		return
	}

	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		handleError(err, w)
		return
	}

	return
}

func GetPandoraPeersOverTime(w http.ResponseWriter, r *http.Request) {
	metrics, err := getPeersOverTime("pandora")
	if err != nil {
		handleError(err, w)
		return
	}

	jsonString, _ := json.Marshal(metrics)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonString))
}

func GetVanguardPeersOverTime(w http.ResponseWriter, r *http.Request) {
	metrics, err := getPeersOverTime("vanguard")
	if err != nil {
		handleError(err, w)
		return
	}

	jsonString, _ := json.Marshal(metrics)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonString))
}

func handleError(err error, w http.ResponseWriter) {
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(err.Error()))
		return
	}
}

func returnBody(body []byte, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func decodeSettings(data []byte) (metrics map[int64]float64, err error) {
	err = json.Unmarshal(data, &metrics)
	if err != nil {
		return
	}
	return
}

func parseMetricFamily(text []byte) (map[string]*dto.MetricFamily, error) {
	var parser expfmt.TextParser
	mf, err := parser.TextToMetricFamilies(bytes.NewReader(text))
	if err != nil {
		return nil, err
	}
	return mf, nil
}
