package goo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEachMap(t *testing.T) {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
	}

	col := make(map[string]int)
	EachMap(m, func(el int, k string, m map[string]int) {
		col[k] = el
	})

	assert.Equal(t, len(m), len(col))
	for k, i := range m {
		assert.Equal(t, i, col[k])
	}
}

func TestEachMapOrdered(t *testing.T) {
	m := map[int]string{
		1: "one",
		4: "four",
		3: "three",
		2: "two",
	}

	keys := []int{}
	col := make(map[int]string)
	EachMapOrdered(m, func(el string, k int, m map[int]string) {
		col[k] = el
		keys = append(keys, k)
	})

	assert.Equal(t, len(m), len(col))
	for k, i := range m {
		assert.Equal(t, i, col[k])
	}

	v := 1
	for i := 0; i < len(m); i++ {
		assert.Equal(t, v, keys[i])
		v++
	}
}

func TestKeys(t *testing.T) {
	m := map[int]string{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
	}

	keys := Keys(m)

	assert.Equal(t, len(m), len(keys))
	v := 1
	for i := 0; i < len(m); i++ {
		assert.Equal(t, v, keys[i])
		v++
	}
}

func TestKeysOrdered(t *testing.T) {
	m := map[int]string{
		1: "one",
		4: "four",
		3: "three",
		2: "two",
	}

	keys := KeysOrdered(m)

	assert.Equal(t, len(m), len(keys))
	v := 1
	for i := 0; i < len(m); i++ {
		assert.Equal(t, v, keys[i])
		v++
	}
}
