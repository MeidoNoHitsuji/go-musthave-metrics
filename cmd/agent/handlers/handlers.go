package handlers

import (
	"fmt"
	"github.com/MeidoNoHitsuji/go-musthave-metrics/cmd/agent/flags"
	"github.com/MeidoNoHitsuji/go-musthave-metrics/internal/storage"
	"github.com/go-resty/resty/v2"
	"log"
	"math/rand"
	"runtime"
	"strconv"
)

type MetricType string

var (
	RStats runtime.MemStats
)

const (
	GAUGE   = MetricType("gauge")
	FLOAT   = MetricType("float64")
	COUNTER = MetricType("counter")
	INT     = MetricType("int64")
)

func SendMetric(m MetricType, name string, value string) (interface{}, error) {
	client := resty.New()

	_, err := client.R().
		Post(fmt.Sprintf("http://%s/update/%s/%s/%s", flags.Addr, m, name, value))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func LoadMetric() {
	runtime.ReadMemStats(&RStats)
	storage.Store.AddGauge("Alloc", RStats.Alloc)
	storage.Store.AddGauge("BuckHashSys", RStats.BuckHashSys)
	storage.Store.AddGauge("Frees", RStats.Frees)
	storage.Store.AddGauge("GCCPUFraction", RStats.GCCPUFraction)
	storage.Store.AddGauge("GCSys", RStats.GCSys)
	storage.Store.AddGauge("HeapAlloc", RStats.HeapAlloc)
	storage.Store.AddGauge("HeapIdle", RStats.HeapIdle)
	storage.Store.AddGauge("HeapInuse", RStats.HeapInuse)
	storage.Store.AddGauge("HeapObjects", RStats.HeapObjects)
	storage.Store.AddGauge("HeapReleased", RStats.HeapReleased)
	storage.Store.AddGauge("HeapSys", RStats.HeapSys)
	storage.Store.AddGauge("LastGC", RStats.LastGC)
	storage.Store.AddGauge("Lookups", RStats.Lookups)
	storage.Store.AddGauge("MCacheInuse", RStats.MCacheInuse)
	storage.Store.AddGauge("MCacheSys", RStats.MCacheSys)
	storage.Store.AddGauge("MSpanInuse", RStats.MSpanInuse)
	storage.Store.AddGauge("MSpanSys", RStats.MSpanSys)
	storage.Store.AddGauge("Mallocs", RStats.Mallocs)
	storage.Store.AddGauge("NextGC", RStats.NextGC)
	storage.Store.AddGauge("NumForcedGC", RStats.NumForcedGC)
	storage.Store.AddGauge("NumGC", RStats.NumGC)
	storage.Store.AddGauge("OtherSys", RStats.OtherSys)
	storage.Store.AddGauge("PauseTotalNs", RStats.PauseTotalNs)
	storage.Store.AddGauge("StackInuse", RStats.StackInuse)
	storage.Store.AddGauge("StackSys", RStats.StackSys)
	storage.Store.AddGauge("Sys", RStats.Sys)
	storage.Store.AddGauge("TotalAlloc", RStats.TotalAlloc)
	storage.Store.AddGauge("RandomValue", rand.Float64())
	storage.Store.AddCounter("PollCount", 1)
}

func SendMetrics() {
	for k, v := range storage.Store.MGauge {
		_, err := SendMetric(GAUGE, k, v)
		if err != nil {
			log.Printf("Ошибка при отправке метрики в Gauge: %s", err.Error())
		}
	}

	storage.Store.MGauge = make(map[string]string)

	for k, v := range storage.Store.MCounter {
		_, err := SendMetric(COUNTER, k, strconv.FormatInt(v, 10))
		if err != nil {
			log.Printf("Ошибка при отправке метрики в Counter: %s", err.Error())
		}
	}

	storage.Store.MCounter = make(map[string]int64)
}
