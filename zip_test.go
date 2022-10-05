package ysq

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZip(t *testing.T) {
	func() {
		qInt := FromSequence(1, 4)
		strs := []string{"one", "two", "three"}
		slice := Zip(qInt, strs, func(first int, second string) string {
			return fmt.Sprintf("%d %s", first, second)
		}).ToSlice(3)
		want := []string{"1 one", "2 two", "3 three"}
		assert.Equal(t, slice, want)
	}()

	func() {
		qInt := FromSequence(1, 4)
		strs := []string{"one", "two", "three", "four"}
		slice := Zip(qInt, strs, func(first int, second string) string {
			return fmt.Sprintf("%d %s", first, second)
		}).ToSlice(4)
		want := []string{"1 one", "2 two", "3 three", "4 four"}
		assert.Equal(t, slice, want)
	}()
}

func TestZipQ(t *testing.T) {
	qInt := FromSequence(1, 4)
	qStr := FromElement("one", "two", "three")
	slice := ZipQ(qInt, qStr, func(first int, second string) string {
		return fmt.Sprintf("%d %s", first, second)
	}).ToSlice(3)
	want := []string{"1 one", "2 two", "3 three"}
	assert.Equal(t, slice, want)
}
