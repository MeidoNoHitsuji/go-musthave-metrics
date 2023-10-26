package main

import (
	"github.com/MeidoNoHitsuji/go-musthave-metrics/internal/flags"
	"github.com/MeidoNoHitsuji/go-musthave-metrics/internal/server"
	"log"
	"net/http"
)

func main() {
	err := flags.ParseFlags()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Сервер запущен")
	err = http.ListenAndServe(flags.Addr, server.Router())
	if err != nil {
		log.Fatalf(err.Error())
	}
}
