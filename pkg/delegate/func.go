package delegate

type (
	// Func 无入参,有出参
	Func[TResult any] func() TResult

	// Func1 1个入参,有出参
	Func1[T1, TResult any] func(T1) TResult

	// Func2 2个入参,有出参
	Func2[T1, T2, TResult any] func(T1, T2) TResult

	// Func3 3个入参,有出参
	Func3[T1, T2, T3, TResult any] func(T1, T2, T3) TResult

	// Func4 4个入参,有出参
	Func4[T1, T2, T3, T4, TResult any] func(T1, T2, T3, T4) TResult

	// Func5 5个入参,有出参
	Func5[T1, T2, T3, T4, T5, TResult any] func(T1, T2, T3, T4, T5) TResult
)

// Partial 局部参数
func (f Func2[T1, T2, TResult]) Partial(arg1 T1) Func1[T2, TResult] {
	return func(arg2 T2) TResult {
		return f(arg1, arg2)
	}
}

// Partial 局部参数
func (f Func3[T1, T2, T3, TResult]) Partial(arg1 T1) Func2[T2, T3, TResult] {
	return func(arg2 T2, arg3 T3) TResult {
		return f(arg1, arg2, arg3)
	}
}

// Partial2 局部参数
func (f Func3[T1, T2, T3, TResult]) Partial2(arg1 T1, arg2 T2) Func1[T3, TResult] {
	return func(arg3 T3) TResult {
		return f(arg1, arg2, arg3)
	}
}
