package ysq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	func() {
		res := FromSequence(1, 100).Contains(func(i int) bool {
			return i%2 == 0
		})
		assert.True(t, res)
	}()

	func() {
		res := FromSequence(1, 100).Contains(func(i int) bool {
			return i > 100
		})
		assert.False(t, res)
	}()
}

func TestAll(t *testing.T) {
	func() {
		res := FromSequence(1, 100).All(func(i int) bool {
			return i < 1000
		})
		assert.True(t, res)
	}()

	func() {
		res := FromSequence(1, 100).All(func(i int) bool {
			return i < 10
		})
		assert.False(t, res)
	}()
}

func TestAny(t *testing.T) {
	res := FromSequence(1, 100).Any()
	assert.True(t, res)
}
