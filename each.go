package ysq

import (
	"github.com/yeungsean/ysq/pkg/delegate"
)

// ForEach 遍历
func (q *Query[T]) ForEach(action delegate.Action1[T]) {
	next := q.Next()
	for item, ok := next(); ok; item, ok = next() {
		action(item)
	}
}

// ForEachN 遍历，带数字
func (q *Query[T]) ForEachN(action delegate.Action2[T, int]) {
	next := q.Next()
	idx := 0
	for item, ok := next(); ok; item, ok = next() {
		action(item, idx)
		idx++
	}
}

// ForEachx 可中断的遍历
func (q *Query[T]) ForEachx(action delegate.FuncTBool[T]) {
	next := q.Next()
	for item, ok := next(); ok; item, ok = next() {
		if !action(item) {
			break
		}
	}
}

// ForEachxN 可中断的遍历，带数字
func (q *Query[T]) ForEachxN(action delegate.FuncTIntBool[T]) {
	next := q.Next()
	idx := 0
	for item, ok := next(); ok; item, ok = next() {
		if !action(item, idx) {
			break
		}
		idx++
	}
}

// ForEachE 遍历，返回error就中断
func (q *Query[T]) ForEachE(f delegate.Func1[T, error]) {
	next := q.Next()
	for item, ok := next(); ok; item, ok = next() {
		if err := f(item); err != nil {
			break
		}
	}
}

// ForEachEN 遍历，带数字，返回error就中断
func (q *Query[T]) ForEachEN(f delegate.Func2[T, int, error]) {
	next := q.Next()
	idx := 0
	for item, ok := next(); ok; item, ok = next() {
		if err := f(item, idx); err != nil {
			break
		}
		idx++
	}
}
