package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testGetSet() *Set[int, int] {
	keySelector := func(i int) int {
		return i
	}
	return MakeSet(keySelector)
}

func TestMakeSet(t *testing.T) {
	s := testGetSet()
	assert.NotNil(t, s.keySelector)
	assert.NotNil(t, s.m)
}

func TestSetAddAndLen(t *testing.T) {
	s := testGetSet()
	s.Add(1, 2, 3, 4, 1)
	assert.Equal(t, s.Len(), 4)
}

func TestSetDelete(t *testing.T) {
	s := testGetSet()
	s.Add(1, 2, 3, 4, 1)
	s.Delete(1)
	assert.Equal(t, s.Len(), 3)
}

func TestSetContains(t *testing.T) {
	s := testGetSet()
	s.Add(1, 2, 3, 4, 1)
	assert.True(t, s.Contains(1))
	assert.False(t, s.Contains(10))
}

func TestSetIterate(t *testing.T) {
	s := testGetSet()
	s.Add(1, 2, 3, 4)
	res := map[int]bool{
		1: false,
		2: false,
		3: false,
		4: false,
	}
	s.Iterate(func(i int) {
		res[i] = true
	})
	for _, v := range res {
		assert.True(t, v)
	}
}

func TestToSlice(t *testing.T) {
	s := testGetSet()
	s.Add(1, 2, 3, 4, 1)
	slice := s.ToSlice()
	assert.Equal(t, 4, len(slice))
	res := map[int]struct{}{
		1: {},
		2: {},
		3: {},
		4: {},
	}
	for _, v := range slice {
		_, ok := res[v]
		assert.True(t, ok)
	}
}
