package server

import (
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testRequest(t *testing.T, ts *httptest.Server, method,
	path string) *resty.Response {

	req := resty.New().R()
	req.Method = method
	req.URL = ts.URL + path

	resp, err := req.Send()
	assert.NoError(t, err, "error making HTTP request")

	return resp
}

func TestRouter(t *testing.T) {
	tests := []struct {
		name       string
		method     string
		path       string
		statusCode int
	}{
		{
			name:       "Проверка запроса получения метрик",
			method:     http.MethodGet,
			path:       "/",
			statusCode: 200,
		},
		{
			name:       "Проверка запроса добавить метрики",
			method:     http.MethodPost,
			path:       "/update/int64/testKey/123",
			statusCode: 200,
		},
		{
			name:       "Проверка запроса получения метрики",
			method:     http.MethodGet,
			path:       "/value/int64/testKey",
			statusCode: 404,
		},
	}
	for _, tt := range tests {
		srv := httptest.NewServer(Router())
		t.Run(tt.name, func(t *testing.T) {
			resp := testRequest(t, srv, tt.method, tt.path)
			assert.Equal(t, resp.StatusCode(), tt.statusCode, "Статус-код не соответствует ответу")
		})
		srv.Close()
	}
}
