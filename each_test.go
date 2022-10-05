package ysq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForEachN(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6}
	q := FromSlice(slice)
	idx := 0
	q.ForEachN(func(val, i int) {
		assert.Equal(t, idx, i)
		assert.Equal(t, slice[i], val)
		idx++
	})
}

func TestForEachx(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6}
	func() {
		q := FromSlice(slice)
		idx := 0
		q.ForEachx(func(i int) bool {
			assert.Equal(t, slice[idx], i)
			idx++
			return true
		})
	}()

	func() {
		q := FromSlice(slice)
		idx := 0
		q.ForEachx(func(i int) bool {
			assert.Equal(t, slice[idx], i)
			if idx > 3 {
				return false
			}
			idx++
			return true
		})
		assert.Equal(t, 4, idx)
	}()
}

func TestForEachxN(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6}
	func() {
		q := FromSlice(slice)
		idx := 0
		q.ForEachxN(func(val, i int) bool {
			assert.Equal(t, slice[idx], val)
			assert.Equal(t, idx, i)
			idx++
			return true
		})
	}()

	func() {
		q := FromSlice(slice)
		idx := 0
		q.ForEachxN(func(val, i int) bool {
			assert.Equal(t, slice[idx], val)
			if idx > 3 {
				return false
			}
			idx++
			return true
		})
		assert.Equal(t, 4, idx)
	}()
}
