// Package slice ...
package slice

import "github.com/yeungsean/ysq/pkg/delegate"

// In 判断切片里包含待比较的数值?
func In[T comparable](slice []T, want T) bool {
	for _, v := range slice {
		if v == want {
			return true
		}
	}
	return false
}

// InBy 通过回调函数,判断切片里包含待比较的值?
func InBy[T any](slice []T, fn delegate.FuncTBool[T]) bool {
	for _, v := range slice {
		if fn(v) {
			return true
		}
	}
	return false
}

// All 判断切片里所有元素等于比较数值?
func All[T comparable](slice []T, want T) bool {
	for _, v := range slice {
		if v != want {
			return false
		}
	}
	return true
}

// AllBy 通过回调函数,判断切片里所有元素满足条件
func AllBy[T any](slice []T, fn delegate.FuncTBool[T]) bool {
	for _, v := range slice {
		if !fn(v) {
			return false
		}
	}
	return true
}
