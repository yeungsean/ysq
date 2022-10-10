package ysq

type (
	// Iterable ...
	Iterable[T any] interface {
		Next() Iterator[T]
	}

	// Iterator 迭代器
	Iterator[T any] func() (item T, ok bool)

	// Comparable 比较器
	Comparable interface {
		CompareTo(Comparable) int
	}
)
