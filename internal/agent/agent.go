package agent

import (
	"fmt"
	"github.com/MeidoNoHitsuji/go-musthave-metrics/internal/flags"
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

func LoadMetric(store *storage.Storage) {
	runtime.ReadMemStats(&RStats)
	store.AddGauge("Alloc", RStats.Alloc)
	store.AddGauge("BuckHashSys", RStats.BuckHashSys)
	store.AddGauge("Frees", RStats.Frees)
	store.AddGauge("GCCPUFraction", RStats.GCCPUFraction)
	store.AddGauge("GCSys", RStats.GCSys)
	store.AddGauge("HeapAlloc", RStats.HeapAlloc)
	store.AddGauge("HeapIdle", RStats.HeapIdle)
	store.AddGauge("HeapInuse", RStats.HeapInuse)
	store.AddGauge("HeapObjects", RStats.HeapObjects)
	store.AddGauge("HeapReleased", RStats.HeapReleased)
	store.AddGauge("HeapSys", RStats.HeapSys)
	store.AddGauge("LastGC", RStats.LastGC)
	store.AddGauge("Lookups", RStats.Lookups)
	store.AddGauge("MCacheInuse", RStats.MCacheInuse)
	store.AddGauge("MCacheSys", RStats.MCacheSys)
	store.AddGauge("MSpanInuse", RStats.MSpanInuse)
	store.AddGauge("MSpanSys", RStats.MSpanSys)
	store.AddGauge("Mallocs", RStats.Mallocs)
	store.AddGauge("NextGC", RStats.NextGC)
	store.AddGauge("NumForcedGC", RStats.NumForcedGC)
	store.AddGauge("NumGC", RStats.NumGC)
	store.AddGauge("OtherSys", RStats.OtherSys)
	store.AddGauge("PauseTotalNs", RStats.PauseTotalNs)
	store.AddGauge("StackInuse", RStats.StackInuse)
	store.AddGauge("StackSys", RStats.StackSys)
	store.AddGauge("Sys", RStats.Sys)
	store.AddGauge("TotalAlloc", RStats.TotalAlloc)
	store.AddGauge("RandomValue", rand.Float64())
	store.AddCounter("PollCount", 1)
}

func SendMetrics(store *storage.Storage) {
	for k, v := range store.MGauge {
		_, err := SendMetric(GAUGE, k, v)
		if err != nil {
			log.Printf("Ошибка при отправке метрики в Gauge: %s", err.Error())
		}
	}

	store.MGauge = make(map[string]string)

	for k, v := range store.MCounter {
		_, err := SendMetric(COUNTER, k, strconv.FormatInt(v, 10))
		if err != nil {
			log.Printf("Ошибка при отправке метрики в Counter: %s", err.Error())
		}
	}

	store.MCounter = make(map[string]int64)
}
