package agent

import (
	"github.com/MeidoNoHitsuji/go-musthave-metrics/internal/storage"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadMetric(t *testing.T) {
	tests := []struct {
		name  string
		store *storage.Storage
	}{
		{
			name:  "Загрузка метрик",
			store: storage.New(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ag := New(tt.store, URL)
			ag.LoadMetric()
			assert.NotEmpty(t, tt.store.MGauge)
			assert.NotEmpty(t, tt.store.MCounter)
		})
	}
}
