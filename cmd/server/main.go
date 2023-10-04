package main

import (
	"github.com/MeidoNoHitsuji/go-musthave-metrics/cmd/server/handlers"
	"github.com/MeidoNoHitsuji/go-musthave-metrics/cmd/server/middlewares"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/update/", handlers.AddMetric)

	log.Printf("Сервер запущен")
	err := http.ListenAndServe(":8080", middlewares.OnlyPOTS(mux))
	if err != nil {
		log.Fatalf(err.Error())
	}
}
