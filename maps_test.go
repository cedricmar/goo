package goo

import (
	"fmt"
	"sort"
	"strconv"
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

	// Sorting to help the testing
	sort.IntSlice(keys).Sort()

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

func TestReduceMap(t *testing.T) {
	m := map[string]int{"1": 1, "2": 2, "3": 3, "4": 4}

	sum := ReduceMap(m, func(mem float64, k string, el int) float64 {
		ik, _ := strconv.Atoi(k)
		return mem + float64(ik+el)
	})

	assert.Equal(t, 20., sum)
}

func TestFilterMap(t *testing.T) {
	m := map[string]int{"1": 1, "2": 2, "3": 3, "4": 4}
	Odd := func(k string, el int) bool {
		return el%2 != 0
	}
	m2 := FilterMap(m, Odd)

	assert.Len(t, m2, len(m)/2)
	for _, el := range m2 {
		assert.True(t, el%2 != 0)
	}
	assert.NotEqual(t, fmt.Sprintf("%p", m), fmt.Sprintf("%p", m2))
}

func TestCloneMap(t *testing.T) {
	m := map[string]bool{"one": true}
	m2 := CloneMap(m)

	assert.Len(t, m2, len(m))
	assert.Equal(t, m["one"], m2["one"])
	assert.NotEqual(t, fmt.Sprintf("%p", m), fmt.Sprintf("%p", m2))
}
