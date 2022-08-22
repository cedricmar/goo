package goo

import (
	"fmt"
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

func TestMap(t *testing.T) {
	is := []int{1, 2, 3, 4}

	Map(is, func(el int) int {
		return el * 2
	})

	for i := 0; i < len(is); i++ {
		assert.Equal(t, (i+1)*2, is[i])
	}
}

func TestClone(t *testing.T) {
	s := []string{"one", "two", "three"}

	cl := Clone(s)

	assert.Equal(t, s, cl)
	assert.NotEqual(t, fmt.Sprintf("%p", s), fmt.Sprintf("%p", cl))
}
