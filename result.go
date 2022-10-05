package ysq

import (
	"github.com/yeungsean/ysq/pkg/delegate"
	"golang.org/x/exp/constraints"
)

func (q *Query[T]) getSliceLength(sliceCap ...uint) int {
	length := 1
	if len(sliceCap) > 0 {
		length = int(sliceCap[0])
	}
	return length
}

// ToSet 返回去重后的序列
func (q *Query[T]) ToSet(sliceCap ...uint) []T {
	length := q.getSliceLength(sliceCap...)
	tmp := q.Distinct()
	result := make([]T, 0, length)
	tmp.ForEach(func(t T) {
		result = append(result, t)
	})
	return result
}

// ToSlice ...
func (q *Query[T]) ToSlice(sliceCap ...uint) []T {
	length := q.getSliceLength(sliceCap...)
	result := make([]T, 0, length)
	q.ForEach(func(t T) {
		result = append(result, t)
	})
	return result
}

// ToChan ...
func (q *Query[T]) ToChan(ch chan<- T) {
	next := q.Next()
	q.Iter(next, func(t T) (v IterContinue) {
		defer func() {
			if err := recover(); err != nil {
				v = IterContinueNo
			}
		}()
		ch <- t
		return IterContinueYes
	})
}

// Count 返回序列中的元素数量
func (q *Query[T]) Count() int64 {
	return q.CountBy(nil)
}

// CountBy 返回表示在指定的序列中满足条件的元素数量的数字
func (q *Query[T]) CountBy(predicate delegate.FuncTBool[T]) int64 {
	next := q.Next()
	cnt := int64(0)
	for item, ok := next(); ok; item, ok = next() {
		if predicate != nil {
			if predicate(item) {
				cnt++
			}
		} else {
			cnt++
		}
	}
	return cnt
}

// ElementAt 返回序列中指定索引处的元素
func (q *Query[T]) ElementAt(index int) (T, error) {
	next := q.Next()
	idx := 0
	for item, ok := next(); ok; item, ok = next() {
		if idx == index {
			return item, nil
		}
		idx++
	}
	var zero T
	return zero, ErrDataNotfound
}

// ElementAtOr 返回序列中指定索引处的元素；如果索引超出范围，则返回默认值
func (q *Query[T]) ElementAtOr(index int, source T) T {
	v, err := q.ElementAt(index)
	if err != nil {
		return source
	}
	return v
}

// NumberComparer ...
func NumberComparer[T constraints.Integer | constraints.Float](prev, current T) int {
	switch {
	case prev > current:
		return -1
	case prev < current:
		return 1
	default:
		return 0
	}
}

// Max 最大值
func (q *Query[T]) Max(comparer Comparer[T]) T {
	next := q.Next()
	item, ok := next()
	if !ok {
		panic(ErrDataNotfound)
	}

	max := item
	for item, ok = next(); ok; item, ok = next() {
		switch comparer(max, item) {
		case 1:
			max = item
		}
	}
	return max
}

// Min 最小值
func (q *Query[T]) Min(comparer Comparer[T]) T {
	next := q.Next()
	item, ok := next()
	if !ok {
		panic(ErrDataNotfound)
	}

	min := item
	for item, ok = next(); ok; item, ok = next() {
		switch comparer(min, item) {
		case -1:
			min = item
		}
	}
	return min
}

// ToMap 从 Query<T> 创建一个 map[TKey]T
func ToMap[T any, TKey comparable](q *Query[T], keySelector delegate.Func1[T, TKey]) map[TKey]T {
	res := make(map[TKey]T)
	q.ForEach(func(t T) {
		res[keySelector(t)] = t
	})
	return res
}
