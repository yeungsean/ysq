package ysq

import (
	"github.com/yeungsean/ysq/pkg/delegate"
)

// FirstOr 返回序列中的第一个元素；如果未找到该元素，则返回默认值
func (q *Query[T]) FirstOr(v T) T {
	return q.FirstOrBy(v, nil)
}

// FirstOrBy 返回序列中满足条件的第一个元素；如果未找到这样的元素，则返回默认值
func (q *Query[T]) FirstOrBy(v T, predicate delegate.FuncTBool[T]) T {
	back, err := q.FirstBy(predicate)
	if err != nil {
		return v
	}
	return back
}

// First 返回序列中第一个元素
func (q *Query[T]) First() (T, error) {
	return q.FirstBy(nil)
}

// FirstBy 返回序列中满足指定条件的第一个元素
func (q *Query[T]) FirstBy(predicate delegate.FuncTBool[T]) (T, error) {
	next := q.Next()
	for item, ok := next(); ok; item, ok = next() {
		switch {
		case predicate != nil:
			if predicate(item) {
				return item, nil
			}
		default:
			return item, nil
		}
	}
	var zero T
	return zero, ErrDataNotfound
}
