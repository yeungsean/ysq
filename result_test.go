package ysq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToSet(t *testing.T) {
	set := FromElement(1, 2, 3, 4, 5, 1, 1, 2, 3, 10, 5).ToSet(6)
	want := map[int]struct{}{
		1:  {},
		2:  {},
		3:  {},
		4:  {},
		5:  {},
		10: {},
	}
	for _, v := range set {
		if _, ok := want[v]; !ok {
			t.Errorf("not contain %v", v)
		}
	}
}

func TestCount(t *testing.T) {
	cnt := FromSequence(1, 11).Count()
	assert.Equal(t, int64(10), cnt)

	cnt = FromSlice([]int{}).Count()
	assert.Equal(t, int64(0), cnt)
}

func TestCountBy(t *testing.T) {
	pets := []testPet{
		{ID: 1, Name: "Niki"},
		{ID: 2, Name: "Maotai"},
		{ID: 3, Name: "Huanhuan"},
		{ID: 4, Name: "Dog"},
	}

	cnt := FromSlice(pets).CountBy(func(tp testPet) bool {
		return len(tp.Name) > 3
	})
	assert.Equal(t, int64(3), cnt)
}

func TestElementAt(t *testing.T) {
	res, err := FromSequence(1, 10).ElementAt(2)
	assert.Nil(t, err)
	assert.Equal(t, 3, res)
}

func TestElementAtOr(t *testing.T) {
	func() {
		res := FromSequence(1, 10).ElementAtOr(100, -1)
		assert.Equal(t, -1, res)
	}()

	func() {
		res := FromSequence(1, 10).ElementAtOr(1, -1)
		assert.Equal(t, 2, res)
	}()
}

func TestToChan(t *testing.T) {
	cr := &ChanResult[int]{
		Ch: make(chan int),
	}
	go func() {
		FromSequence(1, 11).ToChan(cr)
	}()
	i := 0
	for v := range cr.Ch {
		assert.Equal(t, i+1, v)
		i++
	}

	assert.Equal(t, true, cr.GetClosed())
	assert.Equal(t, 10, i)
}

func TestToChanBreak(t *testing.T) {
	cr := &ChanResult[int]{
		Ch: make(chan int),
	}
	go func() {
		FromSequence(1, 10).ToChan(cr)
	}()
	i := 0
	for v := range cr.Ch {
		assert.Equal(t, i+1, v)
		i++
		if v >= 5 {
			cr.SetWaitClose(true)
			break
		}
	}

	assert.Equal(t, false, cr.GetClosed())
	cr.CloseWithTimeout()
	assert.Equal(t, 5, i)
}

func TestToMap(t *testing.T) {
	pets := []testPet{
		{ID: 1, Name: "Niki"},
		{ID: 2, Name: "Maotai"},
		{ID: 3, Name: "Huanhuan"},
		{ID: 4, Name: "Dog"},
	}
	q := FromSlice(pets)
	res := ToMap(q, func(p testPet) int {
		return p.ID
	})
	assert.Equal(t,
		map[int]testPet{
			1: {ID: 1, Name: "Niki"},
			2: {ID: 2, Name: "Maotai"},
			3: {ID: 3, Name: "Huanhuan"},
			4: {ID: 4, Name: "Dog"},
		}, res)
}

func TestMax(t *testing.T) {
	func() {
		seQ := FromSequence(1, 101)
		res := seQ.Max(NumberComparer[int])
		assert.Equal(t, 100, res)
	}()

	func() {
		seQ := FromElement(100, 99, 1, 2, 100, 2, 3, 4, 5)
		res := seQ.Max(NumberComparer[int])
		assert.Equal(t, 100, res)
	}()

	func() {
		seQ := FromSlice([]int{})
		assert.PanicsWithError(t, ErrDataNotfound.Error(), func() {
			seQ.Max(NumberComparer[int])
		})
	}()
}

func TestMin(t *testing.T) {
	func() {
		seQ := FromSequence(1, 100)
		res := seQ.Min(NumberComparer[int])
		assert.Equal(t, 1, res)
	}()

	func() {
		seQ := FromElement(100, 99, 1, 2, 1, 3, 3, 4, 5)
		res := seQ.Min(NumberComparer[int])
		assert.Equal(t, 1, res)
	}()

	func() {
		seQ := FromSlice([]int{})
		assert.PanicsWithError(t, ErrDataNotfound.Error(), func() {
			seQ.Min(NumberComparer[int])
		})
	}()
}
