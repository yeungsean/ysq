package ysq

import "errors"

type (
	// Comparer 大小比较器
	Comparer[T any] func(prev, current T) int

	// GetHashCoder 获取hash值
	GetHashCoder[T any] func(T) int64

	// Query 查询器
	Query[T any] struct {
		Next func() Iterator[T]
	}

	// KeyValuePair 键值对
	KeyValuePair[K, V any] struct {
		Key   K
		Value V
	}

	// KeyListPair 键列表对
	KeyListPair[K, V any] struct {
		Key  K
		List []V
	}

	// IterContinue 终止迭代标记
	IterContinue bool
)

const (
	// IterContinueYes 继续迭代
	IterContinueYes IterContinue = true
	// IterContinueNo 停止迭代
	IterContinueNo IterContinue = false
)

var (
	// ErrDataNotfound 数据未找到
	ErrDataNotfound = errors.New(`not found`)
)

// Iter 迭代
func (q *Query[T]) Iter(next Iterator[T], predicate func(T) IterContinue) (item T, ok bool) {
	for item, ok = next(); ok; item, ok = next() {
		if !predicate(item) {
			return item, ok
		}
	}
	return item, ok
}
