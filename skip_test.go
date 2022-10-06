package ysq

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSkip(t *testing.T) {
	func() {
		actual := FromSequence(1, 21).Skip(10).ToSlice(10)
		want := FromSequence(11, 21).ToSlice(10)
		if !reflect.DeepEqual(actual, want) {
			t.Errorf("10, %v != %v", actual, want)
		}
	}()

	func() {
		actual := FromSequence(1, 21).Skip(10).Take(5).ToSlice(5)
		want := FromSequence(11, 16).ToSlice(5)
		if !reflect.DeepEqual(actual, want) {
			t.Errorf("5, %v != %v", actual, want)
		}
	}()

	func() {
		actual := FromSequence(1, 21).Skip(30).ToSlice(20)
		want := []int{}
		assert.Equal(t, want, actual)
	}()
}

func TestSkipWhile(t *testing.T) {
	func() {
		actual := FromSequence(1, 20).SkipWhile(func(i int) bool {
			return i > 10
		}).ToSlice(10)
		want := FromSequence(11, 20).ToSlice(10)
		assert.Equal(t, want, actual)
	}()

	func() {
		actual := FromSequence(1, 20).SkipWhile(func(i int) bool {
			return i > 20
		}).ToSlice(5)
		want := []int{}
		assert.Equal(t, want, actual)
	}()
}
