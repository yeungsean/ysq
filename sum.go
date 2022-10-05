package ysq

import (
	"golang.org/x/exp/constraints"
)

// Sum ...
func Sum[T any, R constraints.Integer | constraints.Float](q *Query[T], selector func(T) R) R {
	var total R
	next := q.Next()
	for item, ok := next(); ok; item, ok = next() {
		if selector != nil {
			total += selector(item)
		} else {
			var box interface{} = item
			switch tmp := box.(type) {
			case *float64:
				total += (R)(*tmp)
			case float64:
				total += (R)(tmp)
			case *int:
				total += (R)(*tmp)
			case int:
				total += (R)(tmp)
			case *int32:
				total += (R)(*tmp)
			case int32:
				total += (R)(tmp)
			case *int64:
				total += (R)(*tmp)
			case int64:
				total += (R)(tmp)
			case *int16:
				total += (R)(*tmp)
			case int16:
				total += (R)(tmp)
			case *int8:
				total += (R)(*tmp)
			case int8:
				total += (R)(tmp)
			case *uint:
				total += (R)(*tmp)
			case uint:
				total += (R)(tmp)
			case *uint32:
				total += (R)(*tmp)
			case uint32:
				total += (R)(tmp)
			case *uint64:
				total += (R)(*tmp)
			case uint64:
				total += (R)(tmp)
			case *uint16:
				total += (R)(*tmp)
			case uint16:
				total += (R)(tmp)
			case *uint8:
				total += (R)(*tmp)
			case uint8:
				total += (R)(tmp)
			default:
				panic("unsupport")
			}
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
