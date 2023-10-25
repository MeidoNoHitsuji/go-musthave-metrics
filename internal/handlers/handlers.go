package handlers

import (
	"fmt"
	"github.com/MeidoNoHitsuji/go-musthave-metrics/internal/storage"
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

type Handler struct {
	store *storage.Storage
}

func New() *Handler {
	return &Handler{
		store: storage.New(),
	}
}

func (h *Handler) GetMetrics(res http.ResponseWriter, req *http.Request) {
	list := make([]string, 0)

	for s, f := range h.store.MGauge {
		list = append(list, fmt.Sprintf("%s = %s", s, f))
	}

	for s, f := range h.store.MFloat64 {
		list = append(list, fmt.Sprintf("%s = %f", s, f))
	}

	for s, f := range h.store.MCounter {
		list = append(list, fmt.Sprintf("%s = %d", s, f))
	}

	for s, f := range h.store.MInt64 {
		list = append(list, fmt.Sprintf("%s = %d", s, f))
	}

	io.WriteString(res, strings.Join(list, "\n"))
}

func (h *Handler) AddMetric(res http.ResponseWriter, req *http.Request) {
	typeMetric := chi.URLParam(req, "type")
	key := chi.URLParam(req, "key")
	value := chi.URLParam(req, "value")

	switch typeMetric {
	case string(GAUGE):
		_, err := strconv.ParseFloat(value, 64)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			log.Printf("Ошибка параметра в Float: %s", err.Error())
			return
		}
		h.store.AddGauge(key, value)
	case string(FLOAT):
		v, err := strconv.ParseFloat(value, 64)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			log.Printf("Ошибка параметра в Float: %s", err.Error())
			return
		}
		h.store.AddFloat(key, v)
	case string(COUNTER):
		v, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			log.Printf("Ошибка параметра в COUNTER: %s", err.Error())
			return
		}
		h.store.AddCounter(key, v)
	case string(INT):
		v, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			log.Printf("Ошибка параметра в INT: %s", err.Error())
			return
		}
		h.store.AddInt(key, v)
	default:
		res.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (h *Handler) GetMetric(res http.ResponseWriter, req *http.Request) {
	typeMetric := chi.URLParam(req, "type")
	key := chi.URLParam(req, "key")

	switch typeMetric {
	case string(GAUGE):
		if v, exists := h.store.MGauge[key]; !exists {
			res.WriteHeader(http.StatusNotFound)
		} else {
			io.WriteString(res, v)
		}
	case string(FLOAT):
		key := chi.URLParam(req, "key")

		if v, exists := h.store.MFloat64[key]; !exists {
			res.WriteHeader(http.StatusNotFound)
		} else {
			io.WriteString(res, fmt.Sprintf("%f", v))
		}
	case string(COUNTER):
		key := chi.URLParam(req, "key")

		if v, exists := h.store.MCounter[key]; !exists {
			res.WriteHeader(http.StatusNotFound)
		} else {
			io.WriteString(res, fmt.Sprintf("%d", v))
		}
	case string(INT):
		key := chi.URLParam(req, "key")

		if v, exists := h.store.MInt64[key]; !exists {
			res.WriteHeader(http.StatusNotFound)
		} else {
			io.WriteString(res, fmt.Sprintf("%d", v))
		}
	default:
		res.WriteHeader(http.StatusBadRequest)
		return
	}
}
