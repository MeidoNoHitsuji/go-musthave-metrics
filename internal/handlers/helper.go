package handlers

import (
	"errors"
	"github.com/MeidoNoHitsuji/go-musthave-metrics/internal/agent"
	"net/http"
	"strconv"
)

func (h *Handler) StoreMetric(metric agent.Metrics) (int, error) {
	switch metric.MType {
	case string(GAUGE):
		if metric.Value == nil {
			return http.StatusBadRequest, errors.New("Отсутствует значение в параметре Value в типе gauge")
		}
		h.store.AddGauge(metric.ID, *metric.Value)
	case string(FLOAT):
		if metric.Value == nil {
			return http.StatusBadRequest, errors.New("Отсутствует значение в параметре Value в типе float")
		}
		h.store.AddFloat(metric.ID, *metric.Value)
	case string(COUNTER):
		if metric.Delta == nil {
			return http.StatusBadRequest, errors.New("Отсутствует значение в параметре Delta в типе counter")
		}
		h.store.AddCounter(metric.ID, *metric.Delta)
	case string(INT):
		if metric.Delta == nil {
			return http.StatusBadRequest, errors.New("Отсутствует значение в параметре Delta в типе int64")
		}
		h.store.AddInt(metric.ID, *metric.Delta)
	}

	return http.StatusOK, nil
}

func (h *Handler) GetMetricByStruct(metric agent.Metrics) (agent.Metrics, int, error) {
	switch metric.MType {
	case string(GAUGE):
		v, ok := h.store.MGauge[metric.ID]
		if !ok {
			return metric, http.StatusNotFound, nil
		}

		vv, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return metric, http.StatusBadRequest, err
		}
		metric.Value = &vv
	case string(FLOAT):
		v, ok := h.store.MFloat64[metric.ID]
		if !ok {
			return metric, http.StatusNotFound, nil
		}
		metric.Value = &v
	case string(COUNTER):
		v, ok := h.store.MCounter[metric.ID]
		if !ok {
			return metric, http.StatusNotFound, nil
		}
		metric.Delta = &v
	case string(INT):
		v, ok := h.store.MInt64[metric.ID]
		if !ok {
			return metric, http.StatusNotFound, nil
		}
		metric.Delta = &v
	default:
		return metric, http.StatusBadRequest, nil
	}

	return metric, http.StatusOK, nil
}
