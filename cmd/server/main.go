package main

import (
	"github.com/MeidoNoHitsuji/go-musthave-metrics/cmd/server/flags"
	"github.com/MeidoNoHitsuji/go-musthave-metrics/cmd/server/handlers"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func ServerRouter() chi.Router {
	r := chi.NewRouter()

	r.Get("/", handlers.GetMetrics)
	r.Post("/update/{type}/{key}/{value}", handlers.AddMetric)
	r.Get("/value/{type}/{key}", handlers.GetMetric)

	return r
}

func main() {
	flags.ParseFlags()

	log.Printf("Сервер запущен")
	err := http.ListenAndServe(flags.Addr, ServerRouter())
	if err != nil {
		log.Fatalf(err.Error())
	}
}
