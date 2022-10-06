package ysq

import (
	"golang.org/x/exp/constraints"
)

func sumIntegerN[R constraints.Integer | constraints.Float](box interface{}) (R, bool) {
	switch tmp := box.(type) {
	case *int32:
		return (R)(*tmp), true
	case int32:
		return (R)(tmp), true
	case *int64:
		return (R)(*tmp), true
	case int64:
		return (R)(tmp), true
	case *int16:
		return (R)(*tmp), true
	case int16:
		return (R)(tmp), true
	case *int8:
		return (R)(*tmp), true
	case int8:
		return (R)(tmp), true
	case *uint32:
		return (R)(*tmp), true
	case uint32:
		return (R)(tmp), true
	case *uint64:
		return (R)(*tmp), true
	case uint64:
		return (R)(tmp), true
	case *uint16:
		return (R)(*tmp), true
	case uint16:
		return (R)(tmp), true
	case *uint8:
		return (R)(*tmp), true
	case uint8:
		return (R)(tmp), true
	}
	var zero R
	return zero, false
}

func sumNumber[T any, R constraints.Integer | constraints.Float](item T) R {
	var box interface{} = item
	switch tmp := box.(type) {
	case *float64:
		return (R)(*tmp)
	case float64:
		return (R)(tmp)
	case *uint:
		return (R)(*tmp)
	case uint:
		return (R)(tmp)
	case *int:
		return (R)(*tmp)
	case int:
		return (R)(tmp)
	}

	if res, ok := sumIntegerN[R](box); ok {
		return res
	}
	panic("unsupport")
}

// Sum ...
func Sum[T any, R constraints.Integer | constraints.Float](q *Query[T], selector func(T) R) R {
	var total R
	next := q.Next()
	for item, ok := next(); ok; item, ok = next() {
		if selector != nil {
			total += selector(item)
		} else {
			total += sumNumber[T, R](item)
		}
	}
	return total
}

// SumToInt64 ...
func (q *Query[T]) SumToInt64(selector func(T) int64) int64 {
	return Sum(q, selector)
}

// SumToInt32 ...
func (q *Query[T]) SumToInt32(selector func(T) int32) int32 {
	return Sum(q, selector)
}

// SumToInt ...
func (q *Query[T]) SumToInt(selector func(T) int) int {
	return Sum(q, selector)
}

// SumToFloat64 ...
func (q *Query[T]) SumToFloat64(selector func(T) float64) float64 {
	return Sum(q, selector)
}
