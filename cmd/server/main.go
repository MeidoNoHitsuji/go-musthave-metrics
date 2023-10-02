package main

import "net/http"

var Store MemStorage

func main() {
	Store = MemStorage{
		mGauge:   make(map[string]float64),
		mFloat64: make(map[string]float64),
		mCounter: make(map[string]int64),
		mInt64:   make(map[string]int64),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/update/", AddMetric)

	err := http.ListenAndServe(":8080", OnlyPOTS(mux))
	if err != nil {
		return
	}
}
