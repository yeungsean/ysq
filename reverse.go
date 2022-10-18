package ysq

// Reverse 反转
func (q *Query[T]) Reverse(sliceCap ...int) *Query[T] {
	return &Query[T]{
		Next: func() Iterator[T] {
			length := q.getSliceLength(sliceCap...)
			vals := make([]T, 0, length)
			q.ForEach(func(t T) {
				vals = append(vals, t)
			})
			index := len(vals) - 1
			return func() (item T, ok bool) {
				if index < 0 {
					return
				}

				item, ok = vals[index], true
				index--
				return
			}
		},
	}
}
