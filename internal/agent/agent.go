package agent

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"github.com/MeidoNoHitsuji/go-musthave-metrics/internal/flags"
	"github.com/MeidoNoHitsuji/go-musthave-metrics/internal/logger"
	"github.com/MeidoNoHitsuji/go-musthave-metrics/internal/storage"
	"github.com/go-resty/resty/v2"
	"math/rand"
	"runtime"
	"strconv"
)

type (
	MetricType   string
	FormatType   string
	CompressType string
)

var (
	RStats runtime.MemStats
	err    error
)

const (
	GAUGE   = MetricType("gauge")
	FLOAT   = MetricType("float64")
	COUNTER = MetricType("counter")
	INT     = MetricType("int64")

	JSON = FormatType("json")
	URL  = FormatType("url")

	NoneCompress = CompressType("none")
	GZIP         = CompressType("gzip")
)

type Agent struct {
	t            FormatType
	store        *storage.Storage
	compressType CompressType
}

type Metrics struct {
	ID    string   `json:"id"`              // имя метрики
	MType string   `json:"type"`            // параметр, принимающий значение gauge или counter
	Delta *int64   `json:"delta,omitempty"` // значение метрики в случае передачи counter
	Value *float64 `json:"value,omitempty"` // значение метрики в случае передачи gauge
}

func New(store *storage.Storage, t FormatType) *Agent {
	return &Agent{
		t:            t,
		store:        store,
		compressType: NoneCompress,
	}
}

func (a *Agent) SetCompress(compressType CompressType) {
	a.compressType = compressType
}

func (a *Agent) SendMetric(m MetricType, name string, value string) (interface{}, error) {
	client := resty.New()
	request := client.R()
	log := logger.Instant()

	switch a.t {
	case URL:
		_, err = request.
			Post(fmt.Sprintf("http://%s/update/%s/%s/%s", flags.Addr, m, name, value))
	case JSON:
		obj := Metrics{
			ID:    name,
			MType: string(m),
		}

		switch m {
		case GAUGE:
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				log.Errorln("Ошибка параметра в Float: %s", err.Error())
				break
			}
			obj.Value = &v
		case COUNTER:
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				log.Errorln("Ошибка параметра в COUNTER: %s", err.Error())
				break
			}
			obj.Delta = &v
		}

		objByte, err := json.Marshal(obj)
		if err != nil {
			log.Errorln("Ошибка Marshal: %s", err.Error())
			break
		}

		if a.compressType == GZIP {
			request.SetHeader("Content-Encoding", "gzip")
			var buf bytes.Buffer

			gz, err := gzip.NewWriterLevel(&buf, gzip.BestSpeed)
			if err != nil {
				return nil, err
			}

			_, err = gz.Write(objByte)
			if err != nil {
				return nil, err
			}

			err = gz.Close()
			if err != nil {
				return nil, err
			}

		}

		_, err = request.
			SetBody(objByte).
			SetHeader("Content-Type", "application/json").
			Post(fmt.Sprintf("http://%s/update", flags.Addr))
	}

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (a *Agent) LoadMetric() {
	runtime.ReadMemStats(&RStats)
	a.store.AddGauge("Alloc", RStats.Alloc)
	a.store.AddGauge("BuckHashSys", RStats.BuckHashSys)
	a.store.AddGauge("Frees", RStats.Frees)
	a.store.AddGauge("GCCPUFraction", RStats.GCCPUFraction)
	a.store.AddGauge("GCSys", RStats.GCSys)
	a.store.AddGauge("HeapAlloc", RStats.HeapAlloc)
	a.store.AddGauge("HeapIdle", RStats.HeapIdle)
	a.store.AddGauge("HeapInuse", RStats.HeapInuse)
	a.store.AddGauge("HeapObjects", RStats.HeapObjects)
	a.store.AddGauge("HeapReleased", RStats.HeapReleased)
	a.store.AddGauge("HeapSys", RStats.HeapSys)
	a.store.AddGauge("LastGC", RStats.LastGC)
	a.store.AddGauge("Lookups", RStats.Lookups)
	a.store.AddGauge("MCacheInuse", RStats.MCacheInuse)
	a.store.AddGauge("MCacheSys", RStats.MCacheSys)
	a.store.AddGauge("MSpanInuse", RStats.MSpanInuse)
	a.store.AddGauge("MSpanSys", RStats.MSpanSys)
	a.store.AddGauge("Mallocs", RStats.Mallocs)
	a.store.AddGauge("NextGC", RStats.NextGC)
	a.store.AddGauge("NumForcedGC", RStats.NumForcedGC)
	a.store.AddGauge("NumGC", RStats.NumGC)
	a.store.AddGauge("OtherSys", RStats.OtherSys)
	a.store.AddGauge("PauseTotalNs", RStats.PauseTotalNs)
	a.store.AddGauge("StackInuse", RStats.StackInuse)
	a.store.AddGauge("StackSys", RStats.StackSys)
	a.store.AddGauge("Sys", RStats.Sys)
	a.store.AddGauge("TotalAlloc", RStats.TotalAlloc)
	a.store.AddGauge("RandomValue", rand.Float64())
	a.store.AddCounter("PollCount", 1)
}

func (a *Agent) SendMetrics() {
	log := logger.Instant()
	for k, v := range a.store.MGauge {
		_, err := a.SendMetric(GAUGE, k, v)
		if err != nil {
			log.Errorln("Ошибка при отправке метрики в Gauge: %s", err.Error())
		}
	}

	a.store.MGauge = make(map[string]string)

	for k, v := range a.store.MCounter {
		_, err := a.SendMetric(COUNTER, k, strconv.FormatInt(v, 10))
		if err != nil {
			log.Errorln("Ошибка при отправке метрики в Counter: %s", err.Error())
		}
	}

	a.store.MCounter = make(map[string]int64)
}
