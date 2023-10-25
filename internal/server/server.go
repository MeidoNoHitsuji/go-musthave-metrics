package server

import (
	"github.com/MeidoNoHitsuji/go-musthave-metrics/cmd/server/handlers"
	"github.com/go-chi/chi/v5"
)

func ServerRouter() chi.Router {
	r := chi.NewRouter()

	r.Get("/", handlers.GetMetrics)
	r.Post("/update/{type}/{key}/{value}", handlers.AddMetric)
	r.Get("/value/{type}/{key}", handlers.GetMetric)

	return r
}
