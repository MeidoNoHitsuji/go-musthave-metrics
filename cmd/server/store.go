package main

type MemStorage struct {
	mGauge   map[string]float64
	mFloat64 map[string]float64
	mCounter map[string]int64
	mInt64   map[string]int64
}

func (s MemStorage) addGauge(k string, v float64) {
	s.mGauge[k] = v
}

func (s MemStorage) addFloat(k string, v float64) {
	s.mFloat64[k] = v
}

func (s MemStorage) addCounter(k string, v int64) {
	if _, exist := s.mCounter[k]; !exist {
		s.mCounter[k] = 0
	}

	s.mCounter[k] += v
}

func (s MemStorage) addInt(k string, v int64) {
	if _, exist := s.mInt64[k]; !exist {
		s.mInt64[k] = 0
	}

	s.mInt64[k] += v
}
