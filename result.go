package ysq

import (
	"sync"
	"time"

	"github.com/yeungsean/ysq/pkg/delegate"
	"golang.org/x/exp/constraints"
)

func (q *Query[T]) getSliceLength(sliceCap ...int) int {
	length := 1
	if len(sliceCap) > 0 {
		length = sliceCap[0]
	}
	return length
}

// ToSet 返回去重后的序列
func (q *Query[T]) ToSet(sliceCap ...int) []T {
	length := q.getSliceLength(sliceCap...)
	tmp := q.Distinct()
	result := make([]T, 0, length)
	tmp.ForEach(func(t T) {
		result = append(result, t)
	})
	return result
}

// ToSlice 返回切片
func (q *Query[T]) ToSlice(sliceCap ...int) []T {
	length := q.getSliceLength(sliceCap...)
	result := make([]T, 0, length)
	q.ForEach(func(t T) {
		result = append(result, t)
	})
	return result
}

// ChanResult ...
type ChanResult[T any] struct {
	waitCloseRW sync.RWMutex
	closedRW    sync.RWMutex
	sync.Once
	Ch        chan T
	waitClose bool
	closed    bool
}

// GetWaitClose 获取关闭标识位
func (cr *ChanResult[T]) GetWaitClose() bool {
	cr.waitCloseRW.RLock()
	tmp := cr.waitClose
	cr.waitCloseRW.RUnlock()
	return tmp
}

// SetWaitClose 设置关闭标识位
func (cr *ChanResult[T]) SetWaitClose(v bool) {
	cr.waitCloseRW.Lock()
	cr.waitClose = v
	cr.waitCloseRW.Unlock()
}

// GetClosed 获取closed的标识位
func (cr *ChanResult[T]) GetClosed() bool {
	cr.closedRW.RLock()
	tmp := cr.closed
	cr.closedRW.RUnlock()
	return tmp
}

// setClosed 设置closed的标识位
func (cr *ChanResult[T]) setClosed(v bool) {
	cr.closedRW.Lock()
	cr.closed = v
	cr.closedRW.Unlock()
}

// Close 关闭channel
func (cr *ChanResult[T]) Close() {
	if cr.GetClosed() {
		return
	}
	cr.Once.Do(func() {
		cr.setClosed(true)
		close(cr.Ch)
	})
}

// CloseWithTimeout 带超时设定的关闭channel
func (cr *ChanResult[T]) CloseWithTimeout(ts ...time.Duration) {
	if cr.GetClosed() {
		return
	}
	var last time.Duration
	if len(ts) > 0 {
		last = ts[0]
	} else {
		last = time.Millisecond
	}
	select {
	case <-cr.Ch:
		cr.Close()
	case <-time.After(last):
		cr.Close()
	}
}

// ToChan ...
func (q *Query[T]) ToChan(cr *ChanResult[T]) {
	next := q.Next()
	q.Iter(next, func(t T) (v IterContinue) {
		defer func() {
			if err := recover(); err != nil {
				v = IterContinueNo
			}
		}()
		if cr.GetWaitClose() {
			return IterContinueNo
		}
		cr.Ch <- t
		return IterContinueYes
	})
	cr.Close()
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
		if comparer(max, item) == 1 {
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
		if comparer(min, item) == -1 {
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
