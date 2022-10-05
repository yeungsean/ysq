package ysq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirst(t *testing.T) {
	first, err := FromSequence(2, 100).First()
	assert.Equal(t, err, nil)
	assert.Equal(t, first, 2)
}

func TestFirstOr(t *testing.T) {
	first := FromSequence(2, 100).FirstOr(-1)
	assert.Equal(t, first, 2)

	emptySlice := []int{}
	first = FromSlice(emptySlice).FirstOr(-1)
	assert.Equal(t, first, -1)
}

func TestFirstBy(t *testing.T) {
	first, err := FromSequence(2, 100).FirstBy(func(i int) bool {
		return i > 20
	})
	assert.Equal(t, err, nil)
	assert.Equal(t, first, 21)
}
