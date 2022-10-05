package ysq

// Except 生成两个序列的差集
func (q *Query[T]) Except(other []T) *Query[T] {
	otherQ := FromSlice(other)
	return q.ExceptQ(otherQ)
}

// ExceptQ 生成两个序列的差集
func (q *Query[T]) ExceptQ(otherQ *Query[T]) *Query[T] {
	return &Query[T]{
		Next: func() Iterator[T] {
			qNext := q.Next()
			otherNext := otherQ.Next()
			set := make(map[interface{}]struct{})
			diff := make(map[interface{}]struct{})
			for item, ok := otherNext(); ok; item, ok = otherNext() {
				set[item] = struct{}{}
			}

			return func() (item T, ok bool) {
				for item, ok = qNext(); ok; item, ok = qNext() {
					if _, exists := set[item]; exists {
						continue
					}
					if _, exists := diff[item]; !exists {
						diff[item] = struct{}{}
						return
					}
				}

				return
			}
		},
	}
}

// ExpectBy 生成两个序列的差集
func (q *Query[T]) ExpectBy(other []T, getter GetHashCoder[T]) *Query[T] {
	otherQ := FromSlice(other)
	return &Query[T]{
		Next: func() Iterator[T] {
			qNext := q.Next()
			otherNext := otherQ.Next()
			set := make(map[int64]struct{})
			diff := make(map[int64]struct{})
			for item, ok := otherNext(); ok; item, ok = otherNext() {
				set[getter(item)] = struct{}{}
			}

			return func() (item T, ok bool) {
				for item, ok = qNext(); ok; item, ok = qNext() {
					hashVal := getter(item)
					if _, exists := set[hashVal]; exists {
						continue
					}
					if _, exists := diff[hashVal]; !exists {
						diff[hashVal] = struct{}{}
						return
					}
				}

				return
			}
		},
	}
}
