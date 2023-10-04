package storage

import (
	"log"
)

type MemStorage struct {
	MGauge   map[string]float64
	MFloat64 map[string]float64
	MCounter map[string]int64
	MInt64   map[string]int64
}

var Store *MemStorage = &MemStorage{
	MGauge:   make(map[string]float64),
	MFloat64: make(map[string]float64),
	MCounter: make(map[string]int64),
	MInt64:   make(map[string]int64),
}

func (s MemStorage) AddGauge(k string, v float64) {
	s.MGauge[k] = v
	log.Printf("Gauge со значением %s перезаписан на %f", k, v)
}

func (s MemStorage) AddFloat(k string, v float64) {
	s.MFloat64[k] = v
	log.Printf("Float со значением %s перезаписан на %f", k, v)
}

func (s MemStorage) AddCounter(k string, v int64) {
	if _, exist := s.MCounter[k]; !exist {
		s.MCounter[k] = 0
	}

	s.MCounter[k] += v
	log.Printf("К Counter со значением %s добавлено на %d", k, v)
}

func (s MemStorage) AddInt(k string, v int64) {
	if _, exist := s.MInt64[k]; !exist {
		s.MInt64[k] = 0
	}

	s.MInt64[k] += v
	log.Printf("К Int64 со значением %s добавлено на %d", k, v)
}
