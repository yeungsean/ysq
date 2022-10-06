package ysq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnion(t *testing.T) {
	func() {
		q1 := FromSequence(1, 11)
		l2 := []int{1, 2, 9, 11, 21, 39}
		actual := q1.Union(l2).ToSlice()
		want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 21, 39}
		assert.Equal(t, actual, want)
	}()

	func() {
		q1 := FromSequence(1, 11)
		l2 := []int{1, 2, 2, 9, 11, 21, 39}
		actual := q1.Union(l2).ToSlice()
		want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 21, 39}
		assert.Equal(t, actual, want)
	}()
}

func TestUnionBy(t *testing.T) {
	niki := testPetOwner{ID: 1, Name: "Niki"}
	maotai := testPetOwner{ID: 2, Name: "Maotai"}
	huanhuan := testPetOwner{ID: 3, Name: "Huanhuan"}
	zhazha := testPetOwner{ID: 4, Name: "zhazha"}
	doudou := testPetOwner{ID: 5, Name: "doudou"}
	q1 := FromElement(
		niki, maotai, huanhuan, zhazha,
	)
	l2 := []testPetOwner{niki, doudou}

	actual := q1.UnionBy(l2, func(po testPetOwner) int64 {
		return int64(po.ID)
	}).ToSlice(5)
	want := []testPetOwner{niki, maotai, huanhuan, zhazha, doudou}
	assert.Equal(t, actual, want)
}
