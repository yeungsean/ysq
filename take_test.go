package ysq

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTakeWhile(t *testing.T) {
	func() {
		actual := FromSequence[int64](1, 101).TakeWhile(func(i int64) bool {
			return i <= 10
		}).ToSlice(10)
		want := FromSequence[int64](1, 11).ToSlice(10)
		if !reflect.DeepEqual(actual, want) {
			t.Errorf("actual(%v) != want(%v)", actual, want)
		}
	}()

	func() {
		actual := FromSequence[int64](1, 11).TakeWhile(func(i int64) bool {
			return i <= 10
		}).ToSlice(10)
		want := FromSequence[int64](1, 11).ToSlice(10)
		assert.Equal(t, want, actual)
	}()
}

func TestTakeWhileN(t *testing.T) {
	func() {
		actual := FromElement(
			"apple", "passionfruit", "banana", "mango",
			"orange", "blueberry", "grape", "strawberry").TakeWhileN(
			func(fruit string, index int) bool {
				return len(fruit) >= index
			}).ToSlice(5)
		want := []string{"apple", "passionfruit", "banana", "mango",
			"orange", "blueberry"}
		assert.Equal(t, actual, want)
	}()

	func() {
		actual := FromElement(
			"apple", "passionfruit", "banana", "mango",
			"orange", "blueberry", "grape", "strawberry").TakeWhileN(
			func(fruit string, index int) bool {
				return index < 0
			}).ToSlice(5)
		want := []string{}
		assert.Equal(t, actual, want)
	}()

	func() {
		actual := FromElement(
			"apple", "passionfruit", "banana", "mango",
			"orange", "blueberry", "grape", "strawberry").TakeWhileN(
			func(fruit string, index int) bool {
				return index < 10000
			}).ToSlice(8)
		want := []string{
			"apple", "passionfruit", "banana", "mango",
			"orange", "blueberry", "grape", "strawberry",
		}
		assert.Equal(t, want, actual)
	}()
}

func TestTake(t *testing.T) {
	assert.Panics(t, func() {
		FromSequence(1, 10).Take(0)
	})

	func() {
		result := FromSequence[int64](1, 100).Where(func(i int64) bool {
			return i < 20
		}).Take(5).ToSlice()
		if !reflect.DeepEqual(result, []int64{1, 2, 3, 4, 5}) {
			t.Error("not equal sequence")
		}
	}()

	func() {
		result := FromElement[int64](1, 2, 3, 4, 5).Take(3).ToSlice()
		if !reflect.DeepEqual(result, []int64{1, 2, 3}) {
			t.Error("not equal element")
		}
	}()

	func() {
		q := FromElement[int64](1, 2, 3, 4, 5).Take(3)
		if v := q.FirstOr(0); v != 1 {
			t.Error("v != 1")
		}

		result := q.Where(func(i int64) bool { return i > 2 }).ToSlice(1)
		if !reflect.DeepEqual(result, []int64{3}) {
			t.Error("not equal element 1")
		}

		result = q.Where(func(i int64) bool { return i <= 2 }).ToSlice(2)
		if !reflect.DeepEqual(result, []int64{1, 2}) {
			t.Error("not equal element 2")
		}

		result = q.Where(func(i int64) bool { return i <= 2 }).ToSlice()
		if !reflect.DeepEqual(result, []int64{1, 2}) {
			t.Error("not equal element 3")
		}
	}()
}
