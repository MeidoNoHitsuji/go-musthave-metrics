package server

import (
	"github.com/MeidoNoHitsuji/go-musthave-metrics/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func Router() chi.Router {
	r := chi.NewRouter()
	handler := handlers.New()

	r.Get("/", handler.GetMetrics)
	r.Post("/update/{type}/{key}/{value}", handler.AddMetric)
	r.Get("/value/{type}/{key}", handler.GetMetric)

	return r
}
