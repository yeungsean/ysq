package ysq

import (
	"fmt"
)

// Cast 类型转换
func Cast[T, TResult any](q *Query[T], caster func(T) TResult) *Query[TResult] {
	return &Query[TResult]{
		Next: func() Iterator[TResult] {
			next := q.Next()
			return func() (item TResult, ok bool) {
				var tmp T
				if tmp, ok = next(); ok {
					item = caster(tmp)
				}
				return
			}
		},
	}
}

// CastToString 转string的helper
func CastToString[T any]() func(T) string {
	return func(t T) string {
		return fmt.Sprint(t)
	}
}

// go的能力原因，暂不支持结构体函数上定义范型，看以后版本是否支持，目前写法讨巧诡异

// CastToStringBy 转string
func (q *Query[T]) CastToStringBy(caster func(T) string) *Query[string] {
	return Cast(q, caster)
}

// CastToInt64By 转int64
func (q *Query[T]) CastToInt64By(caster func(T) int64) *Query[int64] {
	return Cast(q, caster)
}

// CastToIntBy 转int
func (q *Query[T]) CastToIntBy(caster func(T) int) *Query[int] {
	return Cast(q, caster)
}

// CastToInt32By 转int32
func (q *Query[T]) CastToInt32By(caster func(T) int32) *Query[int32] {
	return Cast(q, caster)
}

// CastToInterfaceBy 转interface{}
func (q *Query[T]) CastToInterfaceBy() *Query[interface{}] {
	return Cast(q, func(v T) interface{} {
		var box interface{} = v
		return box
	})
}
