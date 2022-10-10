package ysq

// Union 生成两个序列的并集
func (q *Query[T]) Union(other []T) *Query[T] {
	otherQ := FromSlice(other)
	return q.UnionQ(otherQ)
}

func unionQFn[T any](set map[interface{}]struct{}, next Iterator[T]) (item T, ok bool) {
	for item, ok = next(); ok; item, ok = next() {
		if _, exists := set[item]; !exists {
			set[item] = struct{}{}
			return
		}
	}
	return
}

// UnionQ 生成两个序列的并集
func (q *Query[T]) UnionQ(otherQ *Query[T]) *Query[T] {
	return &Query[T]{
		Next: func() Iterator[T] {
			next := q.Next()
			otherNext := otherQ.Next()
			set := make(map[interface{}]struct{})
			qUse := true
			return func() (item T, ok bool) {
				if qUse {
					if item, ok = unionQFn(set, next); ok {
						return
					}
					qUse = false
				}

				if item, ok = unionQFn(set, otherNext); ok {
					return
				}

				return
			}
		},
	}
}

func unionByFn[T any](set map[int64]struct{}, getter GetHashCoder[T], next Iterator[T]) (item T, ok bool) {
	for item, ok = next(); ok; item, ok = next() {
		hashVal := getter(item)
		if _, exists := set[hashVal]; !exists {
			set[hashVal] = struct{}{}
			return
		}
	}
	return
}

// UnionBy 生成两个序列的并集
func (q *Query[T]) UnionBy(other []T, getter GetHashCoder[T]) *Query[T] {
	otherQ := FromSlice(other)
	return &Query[T]{
		Next: func() Iterator[T] {
			next := q.Next()
			otherNext := otherQ.Next()
			set := make(map[int64]struct{})
			qUse := true
			return func() (item T, ok bool) {
				if qUse {
					if item, ok = unionByFn(set, getter, next); ok {
						return
					}
					qUse = false
				}

				if item, ok = unionByFn(set, getter, otherNext); ok {
					return
				}

				return
			}
		},
	}
}
