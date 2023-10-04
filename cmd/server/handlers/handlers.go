package handlers

import (
	"github.com/MeidoNoHitsuji/go-musthave-metrics/cmd/server/storage"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type MetricType string

const (
	GAUGE   = MetricType("gauge")
	FLOAT   = MetricType("float64")
	COUNTER = MetricType("counter")
	INT     = MetricType("int64")
)

func AddMetric(res http.ResponseWriter, req *http.Request) {
	keys := strings.Split(strings.Trim(req.URL.String(), "/"), "/")
	keys = keys[1:]
	//res.Write(make([]byte, 0))

	if len(keys) == 1 {
		res.WriteHeader(http.StatusNotFound)
		return
	} else if len(keys) < 3 {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	typeMetric := keys[0]
	key := keys[1]
	value := keys[2]

	switch typeMetric {
	case string(GAUGE):
		v, err := strconv.ParseFloat(value, 64)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			log.Printf("Ошибка параметра в Float: %s", err.Error())
			return
		}
		storage.Store.AddGauge(key, v)
	case string(FLOAT):
		v, err := strconv.ParseFloat(value, 64)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			log.Printf("Ошибка параметра в Float: %s", err.Error())
			return
		}
		storage.Store.AddFloat(key, v)
	case string(COUNTER):
		v, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			log.Printf("Ошибка параметра в Int: %s", err.Error())
			return
		}
		storage.Store.AddCounter(key, v)
	case string(INT):
		v, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			log.Printf("Ошибка параметра в Float: %s", err.Error())
			return
		}
		storage.Store.AddInt(key, v)
	default:
		res.WriteHeader(http.StatusBadRequest)
		return
	}
}
