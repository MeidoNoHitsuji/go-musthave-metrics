package storage

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestMemStorage_AddCounter(t *testing.T) {
	type fields struct {
		MGauge   map[string]string
		MFloat64 map[string]float64
		MCounter map[string]int64
		MInt64   map[string]int64
	}
	type args struct {
		k string
		v int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   fields
	}{
		{
			name: "Проверка добавления",
			fields: fields{
				MGauge:   make(map[string]string),
				MFloat64: make(map[string]float64),
				MCounter: make(map[string]int64),
				MInt64:   make(map[string]int64),
			},
			args: args{
				k: "key",
				v: 123,
			},
			want: fields{
				MGauge:   make(map[string]string),
				MFloat64: make(map[string]float64),
				MCounter: map[string]int64{
					"key": 123,
				},
				MInt64: make(map[string]int64),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				MGauge:   tt.fields.MGauge,
				MFloat64: tt.fields.MFloat64,
				MCounter: tt.fields.MCounter,
				MInt64:   tt.fields.MInt64,
			}
			s.AddCounter(tt.args.k, tt.args.v)
			assert.True(t, reflect.DeepEqual(s.MCounter, tt.want.MCounter))
		})
	}
}

func TestMemStorage_AddFloat(t *testing.T) {
	type fields struct {
		MGauge   map[string]string
		MFloat64 map[string]float64
		MCounter map[string]int64
		MInt64   map[string]int64
	}
	type args struct {
		k string
		v float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   fields
	}{
		{
			name: "Проверка добавления",
			fields: fields{
				MGauge:   make(map[string]string),
				MFloat64: make(map[string]float64),
				MCounter: make(map[string]int64),
				MInt64:   make(map[string]int64),
			},
			args: args{
				k: "key",
				v: 123,
			},
			want: fields{
				MGauge: make(map[string]string),
				MFloat64: map[string]float64{
					"key": 123,
				},
				MCounter: make(map[string]int64),
				MInt64:   make(map[string]int64),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				MGauge:   tt.fields.MGauge,
				MFloat64: tt.fields.MFloat64,
				MCounter: tt.fields.MCounter,
				MInt64:   tt.fields.MInt64,
			}
			s.AddFloat(tt.args.k, tt.args.v)
			assert.True(t, reflect.DeepEqual(s.MFloat64, tt.want.MFloat64))
		})
	}
}

func TestMemStorage_AddGauge(t *testing.T) {
	type fields struct {
		MGauge   map[string]string
		MFloat64 map[string]float64
		MCounter map[string]int64
		MInt64   map[string]int64
	}
	type args struct {
		k string
		v interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   fields
	}{
		{
			name: "Проверка добавления",
			fields: fields{
				MGauge:   make(map[string]string),
				MFloat64: make(map[string]float64),
				MCounter: make(map[string]int64),
				MInt64:   make(map[string]int64),
			},
			args: args{
				k: "key",
				v: "key",
			},
			want: fields{
				MGauge: map[string]string{
					"key": "key",
				},
				MFloat64: make(map[string]float64),
				MCounter: make(map[string]int64),
				MInt64:   make(map[string]int64),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				MGauge:   tt.fields.MGauge,
				MFloat64: tt.fields.MFloat64,
				MCounter: tt.fields.MCounter,
				MInt64:   tt.fields.MInt64,
			}
			s.AddGauge(tt.args.k, tt.args.v)
			assert.True(t, reflect.DeepEqual(s.MGauge, tt.want.MGauge))
		})
	}
}

func TestMemStorage_AddInt(t *testing.T) {
	type fields struct {
		MGauge   map[string]string
		MFloat64 map[string]float64
		MCounter map[string]int64
		MInt64   map[string]int64
	}
	type args struct {
		k string
		v int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   fields
	}{
		{
			name: "Проверка добавления",
			fields: fields{
				MGauge:   make(map[string]string),
				MFloat64: make(map[string]float64),
				MCounter: make(map[string]int64),
				MInt64:   make(map[string]int64),
			},
			args: args{
				k: "key",
				v: 123,
			},
			want: fields{
				MGauge:   make(map[string]string),
				MFloat64: make(map[string]float64),
				MCounter: make(map[string]int64),
				MInt64: map[string]int64{
					"key": 123,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				MGauge:   tt.fields.MGauge,
				MFloat64: tt.fields.MFloat64,
				MCounter: tt.fields.MCounter,
				MInt64:   tt.fields.MInt64,
			}
			s.AddInt(tt.args.k, tt.args.v)
			assert.True(t, reflect.DeepEqual(s.MInt64, tt.want.MInt64))
		})
	}
}
