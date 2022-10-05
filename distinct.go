package ysq

// Distinct 去重
func (q *Query[T]) Distinct() *Query[T] {
	return &Query[T]{
		Next: func() Iterator[T] {
			next := q.Next()
			set := make(map[interface{}]struct{})
			return func() (item T, ok bool) {
				for item, ok = next(); ok; item, ok = next() {
					if _, exists := set[item]; !exists {
						set[item] = struct{}{}
						return
					}
				}
				return
			}
		},
	}
}

// DistinctBy 有条件去重
func (q *Query[T]) DistinctBy(getter GetHashCoder[T]) *Query[T] {
	return &Query[T]{
		Next: func() Iterator[T] {
			next := q.Next()
			set := make(map[int64]struct{})
			return func() (item T, ok bool) {
				for item, ok = next(); ok; item, ok = next() {
					hashCode := getter(item)
					if _, exists := set[hashCode]; !exists {
						set[hashCode] = struct{}{}
						return
					}
				}
				return
			}
		},
	}
}
