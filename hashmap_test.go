package kmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const TestKey = "TEST_KEY"
const TestValue = 123

func Test_hashMap_PutGet(t *testing.T) {
	var hMap = NewHashMap[string, int]()
	hMap.Put(TestKey, TestValue)
	got := hMap.Get(TestKey)
	assert.Equal(t, TestValue, got)
}
