package main

import (
	"github.com/MeidoNoHitsuji/go-musthave-metrics/internal/flags"
	"github.com/MeidoNoHitsuji/go-musthave-metrics/internal/logger"
	"github.com/MeidoNoHitsuji/go-musthave-metrics/internal/server"
	"net/http"
)

func main() {
	err := flags.ParseFlags()
	log := logger.Instant()

	if err != nil {
		log.Fatal(err)
	}

	log.Infoln("Сервер запущен")
	err = http.ListenAndServe(flags.Addr, server.Router())
	if err != nil {
		log.Fatalf(err.Error())
	}
}
