package main

import (
	"github.com/MeidoNoHitsuji/go-musthave-metrics/internal/agent"
	"github.com/MeidoNoHitsuji/go-musthave-metrics/internal/flags"
	"github.com/MeidoNoHitsuji/go-musthave-metrics/internal/storage"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	err := flags.ParseFlags()
	if err != nil {
		log.Fatal(err)
	}

	store := storage.New()

	go func() {
		for {
			time.Sleep(time.Duration(flags.RollInterval) * time.Second)
			agent.LoadMetric(store)
			log.Printf("Метрики собраны")
		}
	}()

	go func() {
		for {
			time.Sleep(time.Duration(flags.ReportInterval) * time.Second)
			agent.SendMetrics(store)
			log.Printf("Метрики отправлены")
		}
	}()

	log.Printf("Сбор метрик запущен")
	Hold()
	log.Printf("Сбор метрик остановлен")
}

func Hold() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGUSR1)
	<-sigChan
}
