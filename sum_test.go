package ysq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumToInt64(t *testing.T) {
	sumInt64 := FromSequence[int64](1, 101).SumToInt64(func(i int64) int64 {
		return i
	})
	assert.Equal(t, sumInt64, int64(5050))
}

func TestSumToInt32(t *testing.T) {
	sumInt32 := FromSequence[int32](1, 101).SumToInt32(func(i int32) int32 {
		return i
	})
	assert.Equal(t, sumInt32, int32(5050))
}

func TestSumToInt(t *testing.T) {
	sumInt := FromSequence(1, 101).SumToInt(func(i int) int {
		return i
	})
	assert.Equal(t, sumInt, 5050)
}

func TestSumToFloat64(t *testing.T) {
	q := FromElement(1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0)
	ret := q.SumToFloat64(nil)
	assert.Equal(t, 55, int(ret))
}

func TestSumPointer(t *testing.T) {
	func() {
		ptr1, ptr2 := uintptr(1), uintptr(2)
		q := FromElement(ptr1, ptr2)
		assert.Panics(t, func() { Sum[uintptr, uintptr](q, nil) })
	}()

	func() {
		one, two, three, four := int8(1), int8(2), int8(3), int8(4)
		iq := []*int8{&one, &two, &three, &four}
		q := FromSlice(iq)
		ret := Sum[*int8, int8](q, nil)
		assert.Equal(t, int8(10), ret)
	}()

	func() {
		one, two, three, four := int16(1), int16(2), int16(3), int16(4)
		iq := []*int16{&one, &two, &three, &four}
		q := FromSlice(iq)
		ret := Sum[*int16, int16](q, nil)
		assert.Equal(t, int16(10), ret)
	}()

	func() {
		one, two, three, four := 1, 2, 3, 4
		iq := []*int{&one, &two, &three, &four}
		q := FromSlice(iq)
		ret := Sum[*int, int](q, nil)
		assert.Equal(t, 10, ret)
	}()

	func() {
		one, two, three, four := int32(1), int32(2), int32(3), int32(4)
		iq := []*int32{&one, &two, &three, &four}
		q := FromSlice(iq)
		ret := Sum[*int32, int32](q, nil)
		assert.Equal(t, int32(10), ret)
	}()

	func() {
		one, two, three, four := int64(1), int64(2), int64(3), int64(4)
		iq := []*int64{&one, &two, &three, &four}
		q := FromSlice(iq)
		ret := Sum[*int64, int64](q, nil)
		assert.Equal(t, int64(10), ret)
	}()

	func() {
		one, two, three, four := uint8(1), uint8(2), uint8(3), uint8(4)
		iq := []*uint8{&one, &two, &three, &four}
		q := FromSlice(iq)
		ret := Sum[*uint8, uint8](q, nil)
		assert.Equal(t, uint8(10), ret)
	}()

	func() {
		one, two, three, four := uint16(1), uint16(2), uint16(3), uint16(4)
		iq := []*uint16{&one, &two, &three, &four}
		q := FromSlice(iq)
		ret := Sum[*uint16, uint16](q, nil)
		assert.Equal(t, uint16(10), ret)
	}()

	func() {
		one, two, three, four := uint(1), uint(2), uint(3), uint(4)
		iq := []*uint{&one, &two, &three, &four}
		q := FromSlice(iq)
		ret := Sum[*uint, uint](q, nil)
		assert.Equal(t, uint(10), ret)
	}()

	func() {
		one, two, three, four := uint32(1), uint32(2), uint32(3), uint32(4)
		iq := []*uint32{&one, &two, &three, &four}
		q := FromSlice(iq)
		ret := Sum[*uint32, uint32](q, nil)
		assert.Equal(t, uint32(10), ret)
	}()

	func() {
		one, two, three, four := uint64(1), uint64(2), uint64(3), uint64(4)
		iq := []*uint64{&one, &two, &three, &four}
		q := FromSlice(iq)
		ret := Sum[*uint64, uint64](q, nil)
		assert.Equal(t, uint64(10), ret)
	}()
}

func TestSum(t *testing.T) {
	func() {
		q := FromSequence(1, 11)
		ret := Sum[int, int](q, nil)
		assert.Equal(t, 55, ret)
	}()

	func() {
		q := FromSequence(int32(1), int32(11))
		ret := Sum[int32, int32](q, nil)
		assert.Equal(t, int32(55), ret)
	}()

	func() {
		q := FromSequence(int64(1), int64(11))
		ret := Sum[int64, int64](q, nil)
		assert.Equal(t, int64(55), ret)
	}()

	func() {
		q := FromSequence(int16(1), int16(11))
		ret := Sum[int16, int16](q, nil)
		assert.Equal(t, int16(55), ret)
	}()

	func() {
		q := FromSequence(int8(1), int8(11))
		ret := Sum[int8, int8](q, nil)
		assert.Equal(t, int8(55), ret)
	}()

	func() {
		q := FromSequence(uint(1), uint(11))
		ret := Sum[uint, uint](q, nil)
		assert.Equal(t, uint(55), ret)
	}()

	func() {
		q := FromSequence(uint8(1), uint8(11))
		ret := Sum[uint8, uint8](q, nil)
		assert.Equal(t, uint8(55), ret)
	}()

	func() {
		q := FromSequence(uint16(1), uint16(11))
		ret := Sum[uint16, uint16](q, nil)
		assert.Equal(t, uint16(55), ret)
	}()

	func() {
		q := FromSequence(uint32(1), uint32(11))
		ret := Sum[uint32, uint32](q, nil)
		assert.Equal(t, uint32(55), ret)
	}()

	func() {
		q := FromSequence(uint64(1), uint64(11))
		ret := Sum[uint64, uint64](q, nil)
		assert.Equal(t, uint64(55), ret)
	}()

	func() {
		q := FromElement(1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0)
		ret := Sum[float64, float64](q, nil)
		assert.Equal(t, int(55), int(ret))
	}()

	func() {
		one, two, three, four := float64(1), float64(2), float64(3), float64(4)
		q := FromElement(&one, &two, &three, &four)
		ret := Sum[*float64, float64](q, nil)
		assert.Equal(t, int(10), int(ret))
	}()
}
