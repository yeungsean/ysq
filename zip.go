package ysq

// Zip 将指定函数应用于两个序列的对应元素，以生成结果序列
func Zip[T, TSecond, TResult any](
	q *Query[T],
	list []TSecond,
	resultSelector func(T, TSecond) TResult,
) *Query[TResult] {
	return &Query[TResult]{
		Next: func() Iterator[TResult] {
			next := q.Next()
			idx := 0
			return func() (item TResult, ok bool) {
				var tmp T
				if tmp, ok = next(); !ok {
					return
				}

				if len(list) <= idx {
					ok = false
					return
				}

				snd := list[idx]
				item = resultSelector(tmp, snd)
				idx++
				return
			}
		},
	}
}

// ZipQ 将指定函数应用于两个序列的对应元素，以生成结果序列
func ZipQ[T, TSecond, TResult any](
	q *Query[T],
	qs *Query[TSecond],
	resultSelector func(T, TSecond) TResult,
) *Query[TResult] {
	return &Query[TResult]{
		Next: func() Iterator[TResult] {
			firstNext := q.Next()
			secondNext := qs.Next()
			return func() (item TResult, ok bool) {
				item1, ok1 := firstNext()
				item2, ok2 := secondNext()
				if ok1 && ok2 {
					return resultSelector(item1, item2), true
				}

				ok = false
				return
			}
		},
	}
}
