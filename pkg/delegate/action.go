package delegate

type (
	// Action 无入参,无出参
	Action func()

	// Action1 1个入参,无出参
	Action1[T1 any] func(T1)

	// Action2 2个入参,无出参
	Action2[T1, T2 any] func(T1, T2)

	// Action3 3个入参,无出参
	Action3[T1, T2, T3 any] func(T1, T2, T3)

	// Action4 4个入参,无出参
	Action4[T1, T2, T3, T4 any] func(T1, T2, T3, T4)

	// Action5 5个入参,无出参
	Action5[T1, T2, T3, T4, T5 any] func(T1, T2, T3, T4, T5)
)

// Partial 局部参数
func (a Action2[T1, T2]) Partial(arg1 T1) Action1[T2] {
	return func(arg2 T2) {
		a(arg1, arg2)
	}
}

// Partial 局部参数
func (a Action3[T1, T2, T3]) Partial(arg1 T1) Action2[T2, T3] {
	return func(arg2 T2, arg3 T3) {
		a(arg1, arg2, arg3)
	}
}

// Partial2 局部参数
func (a Action3[T1, T2, T3]) Partial2(arg1 T1, arg2 T2) Action1[T3] {
	return func(arg3 T3) {
		a(arg1, arg2, arg3)
	}
}
