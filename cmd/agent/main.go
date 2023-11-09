package main

import (
	"github.com/MeidoNoHitsuji/go-musthave-metrics/internal/agent"
	"github.com/MeidoNoHitsuji/go-musthave-metrics/internal/flags"
	"github.com/MeidoNoHitsuji/go-musthave-metrics/internal/logger"
	"github.com/MeidoNoHitsuji/go-musthave-metrics/internal/storage"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	err := flags.ParseFlags()
	log := logger.Instant()
	if err != nil {
		log.Fatal(err)
	}

	store := storage.New()
	ag := agent.New(store, agent.URL)

	go func() {
		for {
			time.Sleep(time.Duration(flags.RollInterval) * time.Second)
			ag.LoadMetric()
			log.Infoln("Метрики собраны")
		}
	}()

	go func() {
		for {
			time.Sleep(time.Duration(flags.ReportInterval) * time.Second)
			ag.SendMetrics()
			log.Infoln("Метрики отправлены")
		}
	}()

	log.Infoln("Сбор метрик запущен")
	Hold()
	log.Infoln("Сбор метрик остановлен")
}

func Hold() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGUSR1)
	<-sigChan
}
