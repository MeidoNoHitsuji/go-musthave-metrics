package server

import (
	"github.com/MeidoNoHitsuji/go-musthave-metrics/internal/handlers"
	"github.com/MeidoNoHitsuji/go-musthave-metrics/internal/middlewares"
	"github.com/go-chi/chi/v5"
)

func Router() chi.Router {
	r := chi.NewRouter()
	handler := handlers.New()

	r.Use(middlewares.WithLogging)

	r.Get("/", handler.GetMetrics)
	r.Post("/update/{type}/{key}/{value}", handler.AddMetric)
	r.Get("/value/{type}/{key}", handler.GetMetric)

	return r
}
