package ysq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersect(t *testing.T) {
	func() {
		q := FromElement(2, 3, 6)
		other := []int{1, 2, 3, 4, 5, 6}
		actual := q.Intersect(other).ToSlice(3)
		want := []int{2, 3, 6}
		assert.Equal(t, actual, want)
	}()

	func() {
		q := FromElement(1, 2, 3, 4, 5, 6)
		other := []int{2, 3, 6}
		actual := q.Intersect(other).ToSlice(3)
		want := []int{2, 3, 6}
		assert.Equal(t, actual, want)
	}()

	func() {
		q := FromElement(1, 2, 2, 3, 4, 5, 6)
		other := []int{2, 3, 3, 6}
		actual := q.Intersect(other).ToSlice(3)
		want := []int{2, 3, 6}
		assert.Equal(t, actual, want)
	}()
}

func TestIntersectBy(t *testing.T) {
	func() {
		niki := testPetOwner{ID: 1, Name: "Niki"}
		maotai := testPetOwner{ID: 2, Name: "Maotai"}
		huanhuan := testPetOwner{ID: 3, Name: "Huanhuan"}
		zhazha := testPetOwner{ID: 4, Name: "zhazha"}
		doudou := testPetOwner{ID: 5, Name: "doudou"}
		q1 := FromElement(niki, maotai, huanhuan, zhazha)
		l2 := []testPetOwner{niki, doudou}
		actual := q1.IntersectBy(l2, func(po testPetOwner) int64 {
			return int64(po.ID)
		}).ToSlice(1)
		want := []testPetOwner{niki}
		assert.Equal(t, actual, want)
	}()

	func() {
		niki := testPetOwner{ID: 1, Name: "Niki"}
		maotai := testPetOwner{ID: 2, Name: "Maotai"}
		huanhuan := testPetOwner{ID: 3, Name: "Huanhuan"}
		zhazha := testPetOwner{ID: 4, Name: "zhazha"}
		doudou := testPetOwner{ID: 5, Name: "doudou"}
		q1 := FromElement(niki, maotai, maotai, huanhuan, zhazha)
		l2 := []testPetOwner{niki, niki, maotai, doudou}
		actual := q1.IntersectBy(l2, func(po testPetOwner) int64 {
			return int64(po.ID)
		}).ToSlice(1)
		want := []testPetOwner{niki, maotai}
		assert.Equal(t, actual, want)
	}()
}
