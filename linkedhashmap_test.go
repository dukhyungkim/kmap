package kmap

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_linkedHashMap_Keys(t *testing.T) {
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
			m:    NewLinkedHashMap[string, int](),
			args: args{
				keys: []string{"D", "F", "A", "C", "B"},
			},
			want: []string{"D", "F", "A", "C", "B"},
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
