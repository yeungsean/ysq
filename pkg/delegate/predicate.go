package delegate

import (
	"reflect"
)

type (
	// FuncTBool 一个入参,返回bool
	FuncTBool[T any] Func1[T, bool]

	// FuncTIntBool 1个范型入参,1个int入参,返回bool
	FuncTIntBool[T any] Func2[T, int, bool]

	// FuncInt3Bool 3个int入参,返回bool
	FuncInt3Bool Func3[int, int, int, bool]

	// FuncString3Bool 3个string入参,返回bool
	FuncString3Bool Func3[string, string, string, bool]
)

var fnHash = map[interface{}]reflect.Value{}

// PartialFunc 相同入参类型的局部函数
func PartialFunc[T, TResult any](fn interface{}, stashArgs ...T) func(...T) TResult {
	if fn == nil {
		panic("fn not nil")
	}

	rvf, ok := fnHash[&fn]
	if !ok {
		rvf = reflect.ValueOf(fn)
		fnHash[&fn] = rvf
	}

	n := rvf.Type().NumIn()
	args := make([]reflect.Value, 0, n)
	for i := 0; i < len(stashArgs); i++ {
		args = append(args, reflect.ValueOf(stashArgs[i]))
	}
	return func(lastArgs ...T) TResult {
		reflectArgs := args
		for i := 0; i < len(lastArgs); i++ {
			reflectArgs = append(reflectArgs, reflect.ValueOf(lastArgs[i]))
		}
		rets := rvf.Call(reflectArgs)
		return rets[0].Interface().(TResult)
	}
}

func partial3Func[T, TResult any](f func(T, T, T) TResult, stashArgs ...T) func(...T) TResult {
	args := make([]T, 0, 3)
	args = append(args, stashArgs...)
	return func(lastArgs ...T) TResult {
		if len(args) >= 3 {
			return f(args[0], args[1], args[2])
		}

		args = append(args, lastArgs...)
		return f(args[0], args[1], args[2])
	}
}

// Partial 局部参数
func (f FuncInt3Bool) Partial(stashArgs ...int) func(...int) bool {
	return partial3Func(f, stashArgs...)
}

// Partial 局部参数
func (f FuncString3Bool) Partial(stashArgs ...string) func(...string) bool {
	return partial3Func(f, stashArgs...)
}
