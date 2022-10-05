package ysq

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelect(t *testing.T) {
	func() {
		intSlice := FromSlice([]int64{1, 2, 3, 4, 5}).Select(func(v int64) int64 {
			return v * v
		}).ToSlice()
		want := []int64{1, 4, 9, 16, 25}
		assert.Equal(t, intSlice, want)
	}()

	func() {
		fe := FromElement(1, 2, 3, 4, 5, 6)
		result := Select(fe, func(v int) string {
			return fmt.Sprint(v)
		}).ToSlice()
		want := []string{"1", "2", "3", "4", "5", "6"}
		assert.Equal(t, result, want)
	}()
}

type testPetOwner struct {
	ID   int
	Name string
	Pets []string
}

func TestSelectMany(t *testing.T) {
	func() {
		slice := FromElement(
			testPetOwner{Name: "cat", Pets: []string{"niki", "maoki"}},
			testPetOwner{Name: "dog", Pets: []string{"rober", "sb"}},
		)
		afterSlice := SelectMany(slice, func(p testPetOwner) *Query[string] {
			return FromSlice(p.Pets)
		}).Where(func(s string) bool {
			return len(s) > 2
		})
		actual := afterSlice.ToSlice()
		want := []string{"niki", "maoki", "rober"}
		assert.Equal(t, want, actual)
		// assert.Equal(t, 3, afterSlice.length)
	}()
}
