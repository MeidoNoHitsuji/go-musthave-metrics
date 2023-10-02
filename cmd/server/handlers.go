package main

import (
	"net/http"
	"strconv"
	"strings"
)

const (
	GAUGE   = "gauge"
	FLOAT   = "float64"
	COUNTER = "counter"
	INT     = "int64"
)

func AddMetric(res http.ResponseWriter, req *http.Request) {
	keys := strings.Split(strings.Trim(req.URL.String(), "/"), "/")
	keys = keys[1:]

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
	case GAUGE:
		v, err := strconv.ParseFloat(value, 64)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		Store.addGauge(key, v)
	case FLOAT:
		v, err := strconv.ParseFloat(value, 64)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		Store.addFloat(key, v)
	case COUNTER:
		v, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		Store.addCounter(key, v)
	case INT:
		v, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		Store.addInt(key, v)
	default:
		res.WriteHeader(http.StatusBadRequest)
		return
	}
}
