package goo

import (
	"fmt"
	"strconv"
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

func TestReduceInt(t *testing.T) {
	is := []int{1, 2, 3, 4, 5}

	sum := Reduce(is, func(mem, el int) int {
		return mem + el
	})

	assert.Equal(t, 15, sum)
}

func TestReduceString(t *testing.T) {
	ss := []string{"a", "b", "c", "d"}

	str := Reduce(ss, func(mem, el string) string {
		return mem + el
	})

	assert.Equal(t, "abcd", str)
}

func TestReduceStringInt(t *testing.T) {
	ss := []string{"1", "2", "3", "4"}

	sum := Reduce(ss, func(mem int, el string) int {
		n, _ := strconv.Atoi(el)
		return mem + n
	})

	assert.Equal(t, 10, sum)
}

func TestClone(t *testing.T) {
	s := []string{"one", "two", "three"}

	cl := Clone(s)

	assert.Equal(t, s, cl)
	assert.NotEqual(t, fmt.Sprintf("%p", s), fmt.Sprintf("%p", cl))
}

func TestUnique(t *testing.T) {
	s := []int{1, 1, 2, 3, 1, 4, 3, 3, 4}
	u := Unique(s)

	assert.ElementsMatch(t, []int{1, 2, 3, 4}, u)
}

func TestFind(t *testing.T) {
	s := []int{2, 4, 6, 15, 3, 7}
	Odd := func(el int) bool {
		return el%2 != 0
	}
	f := Find(s, Odd)

	assert.Equal(t, 15, f)
}

func TestFilter(t *testing.T) {
	s := []int{2, 4, 6, 1, 3, 7}
	Even := func(el int) bool {
		return el%2 == 0
	}
	f := Filter(s, Even)

	assert.ElementsMatch(t, []int{2, 4, 6}, f)
}

func TestFoundIn(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6}
	fnd := FoundIn(s, 4)

	assert.True(t, fnd)
}

func TestFoundAt(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6}
	for k, v := range s {
		i, fnd := FoundAt(s, v)
		assert.Equal(t, k, i)
		assert.True(t, fnd)
	}
	i, fnd := FoundAt(s, 7)
	assert.Equal(t, -1, i)
	assert.False(t, fnd)
}

func TestIndex(t *testing.T) {
	s := []struct {
		id   int
		name string
	}{
		{id: 1, name: "one"},
		{id: 2, name: "two"},
		{id: 3, name: "three"},
	}

	m := Index(s, func(el struct {
		id   int
		name string
	}) int {
		return el.id
	})

	assert.Equal(t, KeysOrdered(m), []int{1, 2, 3})
}
