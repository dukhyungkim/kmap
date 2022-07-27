package kmap

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func Test_Map_Put(t *testing.T) {
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
			name: "no duplicated keys for hashMap",
			m:    NewHashMap[string, int](),
			args: args{
				key:   uuid.NewString(),
				value: rand.Intn(100),
			},
			want: 1,
		},
		{
			name: "no duplicated keys for sortedMap",
			m:    NewSortedMap[string, int](),
			args: args{
				key:   uuid.NewString(),
				value: rand.Intn(100),
			},
			want: 1,
		},
		{
			name: "no duplicated keys for linkedHashMap",
			m:    NewLinkedHashMap[string, int](),
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

func Test_Map_PutGet(t *testing.T) {
	testValue := rand.Intn(100)

	type args struct {
		key   string
		value int
	}

	tests := []struct {
		name string
		maps Map[string, int]
		args args
		want int
	}{
		{
			name: "hashMap",
			maps: NewHashMap[string, int](),
			args: args{
				key:   uuid.NewString(),
				value: testValue,
			},
			want: testValue,
		},
		{
			name: "sortedMap",
			maps: NewSortedMap[string, int](),
			args: args{
				key:   uuid.NewString(),
				value: testValue,
			},
			want: testValue,
		},
		{
			name: "linkedHashMap",
			maps: NewLinkedHashMap[string, int](),
			args: args{
				key:   uuid.NewString(),
				value: testValue,
			},
			want: testValue,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.maps.Put(tt.args.key, tt.args.value)
			got := tt.maps.Get(tt.args.key)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_Map_GetOrDefault(t *testing.T) {
	testValue := rand.Intn(100)
	testDefaultValue := rand.Intn(100)

	type args struct {
		key          string
		value        int
		defaultValue int
	}

	tests := []struct {
		name             string
		maps             Map[string, int]
		args             args
		wantGetOrDefault int
		wantGet          int
	}{
		{
			name: "hashMap",
			maps: NewHashMap[string, int](),
			args: args{
				key:          uuid.NewString(),
				value:        testValue,
				defaultValue: testDefaultValue,
			},
			wantGetOrDefault: testDefaultValue,
			wantGet:          testValue,
		},
		{
			name: "sortedMap",
			maps: NewSortedMap[string, int](),
			args: args{
				key:          uuid.NewString(),
				value:        testValue,
				defaultValue: testDefaultValue,
			},
			wantGetOrDefault: testDefaultValue,
			wantGet:          testValue,
		},
		{
			name: "linkedHashMap",
			maps: NewLinkedHashMap[string, int](),
			args: args{
				key:          uuid.NewString(),
				value:        testValue,
				defaultValue: testDefaultValue,
			},
			wantGetOrDefault: testDefaultValue,
			wantGet:          testValue,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.maps.GetOrDefault(tt.args.key, tt.args.defaultValue)
			assert.Equal(t, tt.wantGetOrDefault, got)

			tt.maps.Put(tt.args.key, tt.args.value)
			got = tt.maps.GetOrDefault(tt.args.key, tt.args.defaultValue)
			assert.Equal(t, tt.wantGet, got)
		})
	}
}

func Test_Map_SizeClear(t *testing.T) {
	size := rand.Intn(10)

	type args struct {
		size int
	}

	tests := []struct {
		name     string
		maps     Map[string, int]
		args     args
		wantSize int
	}{
		{
			name: "hashMap",
			maps: NewHashMap[string, int](),
			args: args{
				size: size,
			},
			wantSize: size,
		},
		{
			name: "sortedMap",
			maps: NewSortedMap[string, int](),
			args: args{
				size: size,
			},
			wantSize: size,
		},
		{
			name: "linkedHashMap",
			maps: NewLinkedHashMap[string, int](),
			args: args{
				size: size,
			},
			wantSize: size,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < size; i++ {
				tt.maps.Put(uuid.NewString(), rand.Intn(100))
			}
			assert.Equal(t, tt.maps.Size(), size)
			tt.maps.Clear()
			assert.Equal(t, tt.maps.Size(), 0)
		})
	}
}

func Test_Map_ContainsKeyValue(t *testing.T) {
	key := uuid.NewString()
	value := rand.Intn(100)

	type args struct {
		key   string
		value int
	}

	tests := []struct {
		name string
		maps Map[string, int]
		args args
	}{
		{
			name: "hashMap",
			maps: NewHashMap[string, int](),
			args: args{
				key:   key,
				value: value,
			},
		},
		{
			name: "sortedMap",
			maps: NewSortedMap[string, int](),
			args: args{
				key:   key,
				value: value,
			},
		},
		{
			name: "linkedHashMap",
			maps: NewLinkedHashMap[string, int](),
			args: args{
				key:   key,
				value: value,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.maps.ContainsKey(tt.args.key)
			assert.False(t, got)

			got = tt.maps.ContainsValue(tt.args.value)
			assert.False(t, got)

			tt.maps.Put(tt.args.key, tt.args.value)
			got = tt.maps.ContainsKey(tt.args.key)
			assert.True(t, got)

			got = tt.maps.ContainsValue(tt.args.value)
			assert.True(t, got)

			got = tt.maps.ContainsKey(uuid.NewString())
			assert.False(t, got)

			got = tt.maps.ContainsValue(rand.Intn(100) + 100)
			assert.False(t, got)
		})
	}
}

func Test_Map_KeysValues(t *testing.T) {
	size := rand.Intn(100)
	keys := make([]string, 0, size)
	values := make([]int, 0, size)

	for i := 0; i < size; i++ {
		keys = append(keys, fmt.Sprint(i))
		values = append(values, i)
	}

	type args struct {
		keys   []string
		values []int
	}

	tests := []struct {
		name       string
		maps       Map[string, int]
		args       args
		wantKeys   []string
		wantValues []int
	}{
		{
			name: "hashMap",
			maps: NewHashMap[string, int](),
			args: args{
				keys:   keys,
				values: values,
			},
			wantKeys:   keys,
			wantValues: values,
		},
		{
			name: "sortedMap",
			maps: NewSortedMap[string, int](),
			args: args{
				keys:   keys,
				values: values,
			},
			wantKeys:   keys,
			wantValues: values,
		},
		{
			name: "linkedHashMap",
			maps: NewLinkedHashMap[string, int](),
			args: args{
				keys:   keys,
				values: values,
			},
			wantKeys:   keys,
			wantValues: values,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotKeys := tt.maps.Keys()
			assert.ElementsMatch(t, gotKeys, []string{})

			gotValues := tt.maps.Values()
			assert.ElementsMatch(t, gotValues, []int{})

			for i, key := range tt.args.keys {
				tt.maps.Put(key, tt.args.values[i])
			}

			gotKeys = tt.maps.Keys()
			assert.ElementsMatch(t, gotKeys, tt.wantKeys)

			gotValues = tt.maps.Values()
			assert.ElementsMatch(t, gotValues, tt.wantValues)
		})
	}
}

func Test_Map_RemoveIsEmpty(t *testing.T) {
	key := uuid.NewString()
	value := rand.Intn(100)

	type args struct {
		key   string
		value int
	}

	tests := []struct {
		name string
		maps Map[string, int]
		args args
	}{
		{
			name: "hashMap",
			maps: NewHashMap[string, int](),
			args: args{
				key:   key,
				value: value,
			},
		},
		{
			name: "sortedMap",
			maps: NewSortedMap[string, int](),
			args: args{
				key:   key,
				value: value,
			},
		},
		{
			name: "linkedHashMap",
			maps: NewLinkedHashMap[string, int](),
			args: args{
				key:   key,
				value: value,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.maps.IsEmpty()
			assert.True(t, got)

			tt.maps.Put(tt.args.key, tt.args.value)
			got = tt.maps.IsEmpty()
			assert.False(t, got)

			tt.maps.Remove(tt.args.key)

			got = tt.maps.IsEmpty()
			assert.True(t, got)
		})
	}
}
