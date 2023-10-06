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

	r.Route("/update", func(r chi.Router) {
		r.Post("/gauge/{key}/{value}", handlers.AddGauge)
		r.Post("/float64/{key}/{value}", handlers.AddFloat)
		r.Post("/counter/{key}/{value}", handlers.AddCounter)
		r.Post("/int64/{key}/{value}", handlers.AddInt)
	})

	r.Route("/value", func(r chi.Router) {
		r.Get("/gauge/{key}", handlers.GetGauge)
		r.Get("/float64/{key}", handlers.GetFloat)
		r.Get("/counter/{key}", handlers.GetCounter)
		r.Get("/int64/{key}", handlers.GetInt)
	})

	log.Printf("Сервер запущен")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
