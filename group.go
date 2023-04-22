package ysq

func groupByGetSet[TKey comparable, T any](q *Query[T], keySelector func(T) TKey) map[TKey][]T {
	next := q.Next()
	set := make(map[TKey][]T)
	for item, ok := next(); ok; item, ok = next() {
		key := keySelector(item)
		set[key] = append(set[key], item)
	}
	return set
}

// GroupBy 分组归类,不统计
func GroupBy[T any, TKey comparable](
	q *Query[T],
	keySelector func(T) TKey,
) *Query[KeyListPair[TKey, T]] {
	return &Query[KeyListPair[TKey, T]]{
		Next: func() Iterator[KeyListPair[TKey, T]] {
			set := groupByGetSet(q, keySelector)
			length, idx, index := len(set), 0, 0
			groups := make([]KeyListPair[TKey, T], length)
			for k, v := range set {
				groups[idx] = KeyListPair[TKey, T]{Key: k, List: v}
				idx++
			}

			return func() (item KeyListPair[TKey, T], ok bool) {
				if ok = index < length; ok {
					item = groups[index]
					index++
				}
				return
			}
		},
	}
}

// GroupByWithC 根据自定义逻辑分组统计
func GroupByWithC[T any, TKey comparable, TResult any](
	q *Query[T],
	keySelector func(T) TKey,
	resultSelector func(TKey, []T) TResult,
) *Query[KeyValuePair[TKey, TResult]] {
	return &Query[KeyValuePair[TKey, TResult]]{
		Next: func() Iterator[KeyValuePair[TKey, TResult]] {
			set := groupByGetSet(q, keySelector)
			length, idx, index := len(set), 0, 0
			groups := make([]KeyValuePair[TKey, TResult], length)
			for k, v := range set {
				groups[idx] = KeyValuePair[TKey, TResult]{
					Key:   k,
					Value: resultSelector(k, v),
				}
				idx++
			}

			return func() (item KeyValuePair[TKey, TResult], ok bool) {
				if ok = index < length; ok {
					item = groups[index]
					index++
				}
				return
			}
		},
	}
}
