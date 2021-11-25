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

type MetricsResponseData struct {
	Peers     int64 `json:"peers"`
	ChainData int64 `json:"chainData"`
}

func VanguardMetrics(w http.ResponseWriter, r *http.Request) {
	body, metricsError := getMetrics("http://127.0.0.1:8080/metrics", w)
	if metricsError != nil {
		handleError(metricsError, w)
		return
	}

	metricFamily, parsingMetricsError := parseMetricFamily(body)

	if parsingMetricsError != nil {
		handleError(parsingMetricsError, w)
		return
	}

	peers := metricFamily["p2p_peer_count"].GetMetric()
	chainData := metricFamily["beacon_head_slot"].GetMetric()

	if peers == nil || chainData == nil {
		return
	}

	// TODO: proper error handling in case the structure of the metrics changes
	var response MetricsResponseData = MetricsResponseData{
		Peers:     int64(*peers[1].Gauge.Value),
		ChainData: int64(*chainData[0].Gauge.Value),
	}

	peersOverTimeError := setPeersOverTime(*peers[1].Gauge.Value, "vanguard")
	if peersOverTimeError != nil {
		fmt.Println(peersOverTimeError)
		handleError(peersOverTimeError, w)
		return
	}

	jsonString, jsonMarshalError := json.Marshal(response)
	if jsonMarshalError != nil {
		fmt.Println(jsonMarshalError)
		handleError(jsonMarshalError, w)
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

	// TODO: proper error handling in case the structure of the metrics changes
	var response MetricsResponseData = MetricsResponseData{
		Peers:     int64(pandoraMetrics["p2p/peers"]),
		ChainData: int64(pandoraMetrics["chain/head/block"]),
	}

	peersOverTimeError := setPeersOverTime(pandoraMetrics["p2p/peers"], "pandora")
	if peersOverTimeError != nil {
		fmt.Println(peersOverTimeError)
		handleError(peersOverTimeError, w)
		return
	}

	jsonString, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
		handleError(err, w)
		return
	}

	returnBody(jsonString, w)
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
