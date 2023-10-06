package main

import (
	"flag"
	"github.com/MeidoNoHitsuji/go-musthave-metrics/cmd/server/handlers"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

var (
	addr = flag.String("a", "localhost:8080", "Адрес сервера формата host:port")
)

func main() {
	flag.Parse()

	r := chi.NewRouter()

	r.Get("/", handlers.GetMetrics)
	r.Post("/update/{type}/{key}/{value}", handlers.AddMetric)
	r.Get("/value/{type}/{key}", handlers.GetMetric)

	log.Printf("Сервер запущен")
	err := http.ListenAndServe(*addr, r)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
