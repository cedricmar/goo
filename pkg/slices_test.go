package goo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEachInt(t *testing.T) {
	is := []int{1, 2, 3, 4}

	col := []int{}
	Each(is, func(el int, i int, s []int) {
		col = append(col, el)
	})

	assert.Equal(t, is, col)
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
