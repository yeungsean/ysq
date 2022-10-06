package ysq

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCastString(t *testing.T) {
	strSlice := FromSlice([]int64{1, 2, 3, 4, 5}).CastToStringBy(CastToString[int64]()).ToSlice()
	want := []string{"1", "2", "3", "4", "5"}
	if !reflect.DeepEqual(strSlice, want) {
		t.Errorf("want %v, actual: %v ", strSlice, want)
	}
}

func TestCastInt(t *testing.T) {
	intSlice := FromSlice([]int64{1, 2, 3, 4, 5}).CastToIntBy(func(i int64) int {
		return int(i)
	}).ToSlice()
	want := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(intSlice, want) {
		t.Errorf("want %v, actual: %v", intSlice, want)
	}
}

func TestCastInt64(t *testing.T) {
	intSlice := FromSlice([]int{1, 2, 3, 4, 5}).CastToInt64By(func(i int) int64 {
		return int64(i)
	}).ToSlice()
	want := []int64{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(intSlice, want) {
		t.Errorf("want %v, actual: %v", intSlice, want)
	}
}

func TestCastInt32(t *testing.T) {
	intSlice := FromSlice([]int{1, 2, 3, 4, 5}).CastToInt32By(func(i int) int32 {
		return int32(i)
	}).ToSlice()
	want := []int32{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(intSlice, want) {
		t.Errorf("want %v, actual: %v", intSlice, want)
	}
}

func TestCastInterface(t *testing.T) {
	int64ToInterfaceSlice := FromSlice([]int64{1, 2, 3, 4, 5}).CastToInterface().ToSlice()
	want := []interface{}{int64(1), int64(2), int64(3), int64(4), int64(5)}
	assert.EqualValues(t, want, int64ToInterfaceSlice)

	strToInterfaceSlice := FromSlice([]string{"ali", "tx", "meituan"}).CastToInterface().ToSlice()
	want2 := []interface{}{"ali", "tx", "meituan"}
	assert.EqualValues(t, want2, strToInterfaceSlice)
}
