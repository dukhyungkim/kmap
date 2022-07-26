package kmap

import (
	"math/rand"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_sortedMap_Put(t *testing.T) {
	type args struct {
		key   string
		value int
	}
	tests := []struct {
		name string
		m    Map[string, int]
		args args
		want int
	}{
		{
			name: "no duplicated keys",
			m:    NewSortedMap[string, int](),
			args: args{
				key:   uuid.NewString(),
				value: rand.Intn(100),
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.Put(tt.args.key, tt.args.value)
			tt.m.Put(tt.args.key, tt.args.value)
			assert.Equal(t, tt.m.Size(), tt.want)
		})
	}
}

func Test_sortedMap_Keys(t *testing.T) {
	type args struct {
		keys []string
	}
	tests := []struct {
		name string
		m    Map[string, int]
		args args
		want []string
	}{
		{
			name: "put unsorted keys",
			m:    NewSortedMap[string, int](),
			args: args{
				keys: []string{"F", "D", "C", "B", "A"},
			},
			want: []string{"A", "B", "C", "D", "F"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, key := range tt.args.keys {
				tt.m.Put(key, rand.Intn(100))
			}

			for i, key := range tt.m.Keys() {
				assert.Equal(t, tt.want[i], key)
			}
		})
	}
}
