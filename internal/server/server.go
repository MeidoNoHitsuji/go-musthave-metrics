package server

import (
	"github.com/MeidoNoHitsuji/go-musthave-metrics/internal/handlers"
	"github.com/MeidoNoHitsuji/go-musthave-metrics/internal/middlewares"
	"github.com/go-chi/chi/v5"
)

func Router() chi.Router {
	r := chi.NewRouter()
	handler := handlers.New()

	r.Use(middlewares.WithLogging, middlewares.GzipHandle, middlewares.UnzipHandle)

	r.Get("/", handler.GetMetrics)
	r.Post("/update/", handler.AddMetricByJson)
	r.Post("/update/{type}/{key}/{value}", handler.AddMetric)
	r.Get("/value/{type}/{key}", handler.GetMetric)
	r.Post("/value/", handler.GetMetricByJson)

	return r
}
