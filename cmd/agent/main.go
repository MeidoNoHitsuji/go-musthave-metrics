package main

import (
	"github.com/MeidoNoHitsuji/go-musthave-metrics/cmd/agent/flags"
	"github.com/MeidoNoHitsuji/go-musthave-metrics/cmd/agent/handlers"
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

	go func() {
		for {
			time.Sleep(time.Duration(flags.RollInterval) * time.Second)
			handlers.LoadMetric()
			log.Printf("Метрики собраны")
		}
	}()

	go func() {
		for {
			time.Sleep(time.Duration(flags.ReportInterval) * time.Second)
			handlers.SendMetrics()
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
