// Package slice ...
package slice

import (
	"github.com/yeungsean/ysq"
	"github.com/yeungsean/ysq/pkg/delegate"
)

// In 判断切片里包含待比较的数值?
func In[T comparable](slice []T, want T) bool {
	return InBy(slice, func(t T) bool {
		return want == t
	})
}

// InBy 通过回调函数,判断切片里包含待比较的值?
func InBy[T any](slice []T, fn delegate.FuncTBool[T]) bool {
	return ysq.FromSlice(slice).In(fn)
}

// All 判断切片里所有元素等于比较数值?
func All[T comparable](slice []T, want T) bool {
	return AllBy(slice, func(t T) bool {
		return t == want
	})
}

// AllBy 通过回调函数,判断切片里所有元素满足条件
func AllBy[T any](slice []T, fn delegate.FuncTBool[T]) bool {
	return ysq.FromSlice(slice).All(fn)
}

// CastToInterface 转interface列表
func CastToInterface[T any](slice []T) []interface{} {
	return ysq.FromSlice(slice).CastToInterface().ToSlice()
}
