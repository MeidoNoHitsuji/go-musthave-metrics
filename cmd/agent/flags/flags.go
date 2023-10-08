package flags

import (
	"flag"
	"os"
	"strconv"
)

var (
	ReportInterval int
	RollInterval   int
	Addr           string
)

func ParseFlags() error {
	flag.IntVar(&ReportInterval, "r", 10, "Частота отправки метрик на сервер")
	flag.IntVar(&RollInterval, "p", 2, "Частота опроса метрик")
	flag.StringVar(&Addr, "a", "localhost:8080", "Адрес сервера формата host:port")
	flag.Parse()

	if envRunAddr, exists := os.LookupEnv("ADDRESS"); exists {
		Addr = envRunAddr
	}

	if envReportInterval, exists := os.LookupEnv("REPORT_INTERVAL"); exists {
		if v, err := strconv.Atoi(envReportInterval); err != nil {
			return err
		} else {
			ReportInterval = v
		}
	}

	if envPollInterval, exists := os.LookupEnv("POLL_INTERVAL"); exists {
		if v, err := strconv.Atoi(envPollInterval); err != nil {
			return err
		} else {
			RollInterval = v
		}
	}

	return nil
}
