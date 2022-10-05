package ysq

// Intersect 生成两个序列的交集
func (q *Query[T]) Intersect(other []T) *Query[T] {
	otherQ := FromSlice(other)
	return q.IntersectQ(otherQ)
}

// IntersectQ 生成两个序列的交集
func (q *Query[T]) IntersectQ(otherQ *Query[T]) *Query[T] {
	return &Query[T]{
		Next: func() Iterator[T] {
			firstNext := q.Next()
			otherNext := otherQ.Next()
			set := make(map[interface{}]struct{})
			diff := make(map[interface{}]struct{})
			for item, ok := firstNext(); ok; item, ok = firstNext() {
				set[item] = struct{}{}
			}

			return func() (item T, ok bool) {
				for item, ok = otherNext(); ok; item, ok = otherNext() {
					if _, exists := set[item]; !exists {
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

// IntersectBy 生成两个序列的交集
func (q *Query[T]) IntersectBy(other []T, getter GetHashCoder[T]) *Query[T] {
	otherQ := FromSlice(other)
	return &Query[T]{
		Next: func() Iterator[T] {
			firstNext := q.Next()
			otherNext := otherQ.Next()
			set := make(map[int64]struct{})
			diff := make(map[int64]struct{})
			for item, ok := firstNext(); ok; item, ok = firstNext() {
				set[getter(item)] = struct{}{}
			}

			return func() (item T, ok bool) {
				for item, ok = otherNext(); ok; item, ok = otherNext() {
					hashVal := getter(item)
					if _, exists := set[hashVal]; !exists {
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
