package ysq

type (
	// Iterable ...
	Iterable[T any] interface {
		Next() Iterator[T]
	}

	// Iterator 迭代器
	Iterator[T any] func() (item T, ok bool)

	// IterableCount 迭代数量
	IterableCount interface {
		Count() uint64
	}

	// Comparable 比较器
	Comparable interface {
		CompareTo(Comparable) int
	}
)
