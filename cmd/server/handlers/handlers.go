package handlers

import (
	"fmt"
	"github.com/MeidoNoHitsuji/go-musthave-metrics/cmd/server/storage"
	"github.com/go-chi/chi/v5"
	"io"
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

func GetMetrics(res http.ResponseWriter, req *http.Request) {
	list := make([]string, 0)

	for s, f := range storage.Store.MGauge {
		list = append(list, fmt.Sprintf("%s = %s", s, f))
	}

	for s, f := range storage.Store.MFloat64 {
		list = append(list, fmt.Sprintf("%s = %f", s, f))
	}

	for s, f := range storage.Store.MCounter {
		list = append(list, fmt.Sprintf("%s = %d", s, f))
	}

	for s, f := range storage.Store.MInt64 {
		list = append(list, fmt.Sprintf("%s = %d", s, f))
	}

	io.WriteString(res, strings.Join(list, "\n"))
}

func AddGauge(res http.ResponseWriter, req *http.Request) {
	key := chi.URLParam(req, "key")
	value := chi.URLParam(req, "value")

	storage.Store.AddGauge(key, value)
}

func GetGauge(res http.ResponseWriter, req *http.Request) {
	key := chi.URLParam(req, "key")

	if v, exists := storage.Store.MGauge[key]; !exists {
		res.WriteHeader(http.StatusNotFound)
	} else {
		io.WriteString(res, fmt.Sprintf("%s", v))
	}
}

func AddFloat(res http.ResponseWriter, req *http.Request) {
	key := chi.URLParam(req, "key")
	value := chi.URLParam(req, "value")

	v, err := strconv.ParseFloat(value, 64)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		log.Printf("Ошибка параметра в Float: %s", err.Error())
		return
	}
	storage.Store.AddFloat(key, v)
}

func GetFloat(res http.ResponseWriter, req *http.Request) {
	key := chi.URLParam(req, "key")

	if v, exists := storage.Store.MFloat64[key]; !exists {
		res.WriteHeader(http.StatusNotFound)
	} else {
		io.WriteString(res, fmt.Sprintf("%f", v))
	}
}

func AddCounter(res http.ResponseWriter, req *http.Request) {
	key := chi.URLParam(req, "key")
	value := chi.URLParam(req, "value")

	v, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		log.Printf("Ошибка параметра в Int: %s", err.Error())
		return
	}
	storage.Store.AddCounter(key, v)
}

func GetCounter(res http.ResponseWriter, req *http.Request) {
	key := chi.URLParam(req, "key")

	if v, exists := storage.Store.MCounter[key]; !exists {
		res.WriteHeader(http.StatusNotFound)
	} else {
		io.WriteString(res, fmt.Sprintf("%d", v))
	}
}

func AddInt(res http.ResponseWriter, req *http.Request) {
	key := chi.URLParam(req, "key")
	value := chi.URLParam(req, "value")

	v, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		log.Printf("Ошибка параметра в Float: %s", err.Error())
		return
	}
	storage.Store.AddInt(key, v)
}

func GetInt(res http.ResponseWriter, req *http.Request) {
	key := chi.URLParam(req, "key")

	if v, exists := storage.Store.MCounter[key]; !exists {
		res.WriteHeader(http.StatusNotFound)
	} else {
		io.WriteString(res, fmt.Sprintf("%d", v))
	}
}
