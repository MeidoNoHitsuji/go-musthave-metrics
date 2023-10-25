package handlers

import (
	"github.com/MeidoNoHitsuji/go-musthave-metrics/internal/storage"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func makeServer(f func(http.ResponseWriter, *http.Request)) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		return nil, err
	}

	rr := httptest.NewRecorder()
	handlerFunc := http.HandlerFunc(f)
	handlerFunc.ServeHTTP(rr, req)
	return rr, nil
}

func TestHandler_AddMetric(t *testing.T) {
	tests := []struct {
		name       string
		store      *storage.Storage
		want       string
		statusCode int
	}{
		{
			name: "Тест с пустым ответом",
			store: &storage.Storage{
				MGauge:   map[string]string{},
				MFloat64: map[string]float64{},
				MCounter: map[string]int64{},
				MInt64:   map[string]int64{},
			},
			want:       ``,
			statusCode: 400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				store: tt.store,
			}

			rr, err := makeServer(h.AddMetric)
			assert.NoError(t, err)
			assert.Equal(t, rr.Body.String(), tt.want)
			assert.Equal(t, rr.Code, tt.statusCode)
		})
	}
}

func TestHandler_GetMetric(t *testing.T) {
	tests := []struct {
		name       string
		store      *storage.Storage
		want       string
		statusCode int
	}{
		{
			name: "Тест с пустым ответом",
			store: &storage.Storage{
				MGauge:   map[string]string{},
				MFloat64: map[string]float64{},
				MCounter: map[string]int64{},
				MInt64:   map[string]int64{},
			},
			want:       ``,
			statusCode: 400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				store: tt.store,
			}

			rr, err := makeServer(h.GetMetric)
			assert.NoError(t, err)
			assert.Equal(t, rr.Body.String(), tt.want)
			assert.Equal(t, rr.Code, tt.statusCode)
		})
	}
}

func TestHandler_GetMetrics(t *testing.T) {
	tests := []struct {
		name  string
		store *storage.Storage
		want  string
	}{
		{
			name: "Тест с пустым ответом",
			store: &storage.Storage{
				MGauge:   map[string]string{},
				MFloat64: map[string]float64{},
				MCounter: map[string]int64{},
				MInt64:   map[string]int64{},
			},
			want: ``,
		},
		{
			name: "Тест с наполненными данными",
			store: &storage.Storage{
				MGauge: map[string]string{
					"testParam1": "123.00",
				},
				MFloat64: map[string]float64{
					"testParam2": 124.000000,
				},
				MCounter: map[string]int64{
					"testParam3": 1,
				},
				MInt64: map[string]int64{
					"testParam4": 2,
				},
			},
			want: `testParam1 = 123.00
testParam2 = 124.000000
testParam3 = 1
testParam4 = 2`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				store: tt.store,
			}

			rr, err := makeServer(h.GetMetrics)
			assert.NoError(t, err)
			assert.Equal(t, rr.Body.String(), tt.want)
		})
	}
}
