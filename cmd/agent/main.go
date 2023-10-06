package main

import (
	"flag"
	"github.com/MeidoNoHitsuji/go-musthave-metrics/cmd/agent/handlers"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	reportInterval = flag.Int("r", 10, "Частота отправки метрик на сервер")
	pollInterval   = flag.Int("p", 2, "Частота опроса метрик")
)

func main() {
	flag.Parse()

	go func() {
		for {
			time.Sleep(time.Duration(*pollInterval) * time.Second)
			handlers.LoadMetric()
			log.Printf("Метрики собраны")
		}
	}()

	go func() {
		for {
			time.Sleep(time.Duration(*reportInterval) * time.Second)
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
