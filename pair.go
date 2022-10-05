package ysq

type (
	// KeyValuePair 键值对
	KeyValuePair[K, V any] struct {
		Key K
		Value V
	}

	// KeyListPair 键列表对
	KeyListPair[K, V any] struct {
		Key K
		List []V
	}
)
