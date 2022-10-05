package ysq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	slice := FromSequence(1, 20).Filter(func(i int) bool {
		return i < 10
	}).ToSlice(10)
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	assert.Equal(t, slice, want)
}

func TestFilterN(t *testing.T) {
	slice := FromSequence(1, 20).FilterN(func(val, idx int) bool {
		return idx < 5 && val <= 2
	}).ToSlice(2)
	want := []int{1, 2}
	assert.Equal(t, slice, want)
}
