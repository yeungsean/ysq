package ysq

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	func() {
		actual := FromSequence[int64](1, 6).Reverse().ToSlice(5)
		want := []int64{5, 4, 3, 2, 1}
		if !reflect.DeepEqual(actual, want) {
			t.Errorf("actual(%v) != want(%v)", actual, want)
		}
	}()
}
