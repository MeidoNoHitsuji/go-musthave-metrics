package main

import (
	"fmt"
	"github.com/MeidoNoHitsuji/go-musthave-metrics/cmd/server/handlers"
	"github.com/MeidoNoHitsuji/go-musthave-metrics/cmd/server/router"
	"github.com/MeidoNoHitsuji/go-musthave-metrics/cmd/server/storage"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testRequest(t *testing.T, ts *httptest.Server, method,
	path string, body interface{}) *resty.Response {

	req := resty.New().R()
	req.Method = method
	req.URL = ts.URL + path
	if body != nil {
		req.Body = body
	}

	resp, err := req.Send()
	assert.NoError(t, err, "error making HTTP request")

	return resp
}

func emptyStorage() {
	storage.Store = &storage.MemStorage{
		MGauge:   make(map[string]string),
		MFloat64: make(map[string]float64),
		MCounter: make(map[string]int64),
		MInt64:   make(map[string]int64),
	}
}

func TestGetMetrics(t *testing.T) {
	srv := httptest.NewServer(router.ServerRouter())
	defer srv.Close()

	type args struct {
		method string
		path   string
		body   interface{}
	}

	type want struct {
		statusCode int
		body       string
	}

	tests := []struct {
		name     string
		requests []args
		want     want
	}{
		{
			name:     "Пустой ответ",
			requests: make([]args, 0),
			want: want{
				statusCode: 200,
				body:       ``,
			},
		},
		{
			name: "Наполненные параметры",
			requests: []args{
				{
					method: http.MethodPost,
					path:   fmt.Sprintf("/update/%s/testParam1/123.00", handlers.GAUGE),
					body:   nil,
				},
				{
					method: http.MethodPost,
					path:   fmt.Sprintf("/update/%s/testParam2/124.00", handlers.FLOAT),
					body:   nil,
				},
				{
					method: http.MethodPost,
					path:   fmt.Sprintf("/update/%s/testParam3/1", handlers.COUNTER),
					body:   nil,
				},
				{
					method: http.MethodPost,
					path:   fmt.Sprintf("/update/%s/testParam4/2", handlers.INT),
					body:   nil,
				},
			},
			want: want{
				statusCode: 200,
				body: `testParam1 = 123.00
testParam2 = 124.000000
testParam3 = 1
testParam4 = 2`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			emptyStorage()
			for _, request := range tt.requests {
				testRequest(t, srv, request.method, request.path, request.body)
			}
			resp := testRequest(t, srv, http.MethodGet, "", nil)
			assert.Equal(t, resp.StatusCode(), tt.want.statusCode, "Статус-код не соответствует ответу")
			assert.Equal(t, resp.String(), tt.want.body, "Тело ответа не соответствует ответу")
		})
	}
}

func TestAddAndGetMetric(t *testing.T) {
	srv := httptest.NewServer(router.ServerRouter())
	defer srv.Close()

	type args struct {
		method     string
		typeMetric string
		key        string
		value      string
	}

	type want struct {
		statusCode int
		typeMetric string
		key        string
		body       string
	}

	tests := []struct {
		name     string
		requests []args
		want     want
	}{
		{
			name: "Проверка параметра GAUGE",
			requests: []args{
				{
					method:     http.MethodPost,
					typeMetric: string(handlers.GAUGE),
					key:        "testKey",
					value:      "123.00",
				},
			},
			want: want{
				statusCode: 200,
				typeMetric: string(handlers.GAUGE),
				key:        "testKey",
				body:       "123.00",
			},
		},
		{
			name: "Проверка параметра FLOAT",
			requests: []args{
				{
					method:     http.MethodPost,
					typeMetric: string(handlers.FLOAT),
					key:        "testKey",
					value:      "123.00",
				},
			},
			want: want{
				statusCode: 200,
				typeMetric: string(handlers.FLOAT),
				key:        "testKey",
				body:       "123.000000",
			},
		},
		{
			name: "Проверка параметра COUNTER",
			requests: []args{
				{
					method:     http.MethodPost,
					typeMetric: string(handlers.COUNTER),
					key:        "testKey",
					value:      "1",
				},
				{
					method:     http.MethodPost,
					typeMetric: string(handlers.COUNTER),
					key:        "testKey",
					value:      "5",
				},
			},
			want: want{
				statusCode: 200,
				typeMetric: string(handlers.COUNTER),
				key:        "testKey",
				body:       "6",
			},
		},
		{
			name: "Проверка параметра INT",
			requests: []args{
				{
					method:     http.MethodPost,
					typeMetric: string(handlers.INT),
					key:        "testKey",
					value:      "1",
				},
				{
					method:     http.MethodPost,
					typeMetric: string(handlers.INT),
					key:        "testKey",
					value:      "5",
				},
			},
			want: want{
				statusCode: 200,
				typeMetric: string(handlers.INT),
				key:        "testKey",
				body:       "6",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			emptyStorage()
			for _, request := range tt.requests {
				testRequest(t, srv, request.method, fmt.Sprintf("/update/%s/%s/%s", request.typeMetric, request.key, request.value), nil)
			}
			resp := testRequest(t, srv, http.MethodGet, fmt.Sprintf("/value/%s/%s", tt.want.typeMetric, tt.want.key), nil)
			assert.Equal(t, resp.StatusCode(), tt.want.statusCode, "Статус-код не соответствует ответу")
			assert.Equal(t, resp.String(), tt.want.body, "Тело ответа не соответствует ответу")
		})
	}
}

//func TestGetMetric(t *testing.T) {
//	type args struct {
//		res http.ResponseWriter
//		req *http.Request
//	}
//	tests := []struct {
//		name string
//		args args
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			GetMetric(tt.args.res, tt.args.req)
//		})
//	}
//}
