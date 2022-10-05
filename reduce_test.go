package ysq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReduce(t *testing.T) {
	func() {
		q := FromSequence(1, 10)
		total := Reduce(q, 0, func(result, v int) int {
			return result + v
		})
		assert.Equal(t, 55, total)
	}()

	func() {
		q := FromSequence(1, 10)
		total := q.Reduce(0, func(result, i2 int) int {
			return result + i2
		})
		assert.Equal(t, 55, total)
	}()
}