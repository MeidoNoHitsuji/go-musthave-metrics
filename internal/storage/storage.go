package storage

import (
	"fmt"
)

type Storage struct {
	MGauge   map[string]string
	MFloat64 map[string]float64
	MCounter map[string]int64
	MInt64   map[string]int64
}

func New() *Storage {
	return &Storage{
		MGauge:   make(map[string]string),
		MFloat64: make(map[string]float64),
		MCounter: make(map[string]int64),
		MInt64:   make(map[string]int64),
	}
}

func (s *Storage) AddGauge(k string, v interface{}) {
	s.MGauge[k] = fmt.Sprintf("%v", v)
}

func (s *Storage) AddFloat(k string, v float64) {
	s.MFloat64[k] = v
}

func (s *Storage) AddCounter(k string, v int64) {
	if _, exist := s.MCounter[k]; !exist {
		s.MCounter[k] = 0
	}

	s.MCounter[k] += v
}

func (s *Storage) AddInt(k string, v int64) {
	if _, exist := s.MInt64[k]; !exist {
		s.MInt64[k] = 0
	}

	s.MInt64[k] += v
}
