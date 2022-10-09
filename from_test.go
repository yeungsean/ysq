package ysq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkFromElementB(b *testing.B) {
	b.Run("", func(b *testing.B) {
		FromElement(1, 2, 3, 4, 5).ToSlice()
	})
}

func TestFromString(t *testing.T) {
	FromString(`good good boy`)
}

func TestFromSequenceChan(t *testing.T) {
	slice := FromSequenceChan(1, 11).ToSlice(10)
	assert.Equal(t, slice, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	assert.Panics(t, func() {
		FromSequenceChan(1, -1)
	})

	assert.Panics(t, func() {
		FromSequenceChan(1, 1)
	})

	func() {
		actual := FromSequenceChan(1, 10, 2).ToSlice()
		want := []int{1, 3, 5, 7, 9}
		assert.Equal(t, want, actual)
	}()
}

func TestFromSequence(t *testing.T) {
	assert.Panics(t, func() {
		FromSequence(-1, -2)
	})

	assert.Panics(t, func() {
		FromSequence(1, 1)
	})

	func() {
		actual := FromSequence(1, 10, 2).ToSlice(5)
		want := []int{1, 3, 5, 7, 9}
		assert.Equal(t, want, actual)
	}()

	func() {
		actual := FromSequence(1, 10).ToSlice(10)
		want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		assert.Equal(t, want, actual)
	}()
}

func TestFromMap(t *testing.T) {
	m := map[int]string{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
	}
	func() {
		q := FromMap(m)
		keys := Select(q, func(kvp KeyValuePair[int, string]) int {
			return kvp.Key
		}).ToSlice(4)
		want := map[int]struct{}{
			1: {}, 2: {}, 3: {}, 4: {},
		}
		for _, key := range keys {
			_, ok := want[key]
			assert.True(t, ok)
		}
	}()

	func() {
		q := FromMap(m)
		vals := Select(q, func(kvp KeyValuePair[int, string]) string {
			return kvp.Value
		}).ToSlice(4)
		want := map[string]struct{}{
			"one":   {},
			"two":   {},
			"three": {},
			"four":  {},
		}
		for _, val := range vals {
			_, ok := want[val]
			assert.True(t, ok)
		}
	}()
}
