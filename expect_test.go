package ysq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpect(t *testing.T) {
	q1 := FromElement(2.0, 2.0, 2.1, 2.2, 2.3, 2.3, 2.4, 2.5)
	q2 := []float64{2.2}
	actual := q1.Except(q2).ToSlice(5)
	want := []float64{2, 2.1, 2.3, 2.4, 2.5}
	assert.Equal(t, want, actual)
}

func TestExpectBy(t *testing.T) {
	q1 := FromElement(2.0, 2.0, 2.1, 2.2, 2.3, 2.3, 2.4, 2.5)
	q2 := []float64{2.2}
	actual := q1.ExpectBy(q2, func(f float64) int64 {
		return int64(f * 10)
	}).ToSlice(5)
	want := []float64{2.0, 2.1, 2.3, 2.4, 2.5}
	assert.Equal(t, want, actual)
}
