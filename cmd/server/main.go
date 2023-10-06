package main

import (
	"github.com/MeidoNoHitsuji/go-musthave-metrics/cmd/server/handlers"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", handlers.GetMetrics)
	r.Post("/update/{type}/{key}/{value}", handlers.AddMetric)
	r.Get("/value/{type}/{key}", handlers.GetMetric)

	log.Printf("Сервер запущен")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
