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
	err            error
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
		ReportInterval, err = strconv.Atoi(envReportInterval)
		if err != nil {
			return err
		}

	}

	if envPollInterval, exists := os.LookupEnv("POLL_INTERVAL"); exists {
		RollInterval, err = strconv.Atoi(envPollInterval)
		if err != nil {
			return err
		}
	}

	return nil
}
