package ysq

// Select 将序列中的每个元素投影到新表单
func Select[T, TResult any](q *Query[T], selector func(T) TResult) *Query[TResult] {
	qr := &Query[TResult]{
		Next: func() Iterator[TResult] {
			next := q.Next()
			return func() (item TResult, ok bool) {
				var tmp T
				if tmp, ok = next(); ok {
					item = selector(tmp)
				}
				return
			}
		},
	}
	return qr
}

func selectManyIterator[T, TResult any](q *Query[T], _ *Query[TResult],
	selector func(T) *Query[TResult]) func() Iterator[TResult] {
	return func() Iterator[TResult] {
		outernext := q.Next()
		var inner *T
		var innernext Iterator[TResult]
		return func() (item TResult, ok bool) {
			var tmp T
			for !ok {
				if inner == nil {
					if tmp, ok = outernext(); !ok {
						return
					}
					inner = &tmp
					innernext = selector(*inner).Next()
				}

				if item, ok = innernext(); !ok {
					inner = nil
				}
			}
			return
		}
	}
}

// SelectMany 将序列的每个元素投影并将结果序列合并为一个序列
func SelectMany[T, TResult any](q *Query[T], selector func(T) *Query[TResult]) *Query[TResult] {
	qr := Query[TResult]{}
	qr.Next = selectManyIterator(q, &qr, selector)
	return &qr
}

// Select 将序列中的每个元素投影到新表单
func (q *Query[T]) Select(selector func(T) T) *Query[T] {
	return Select(q, selector)
}
