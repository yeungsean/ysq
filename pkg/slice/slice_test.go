package slice

import (
	"reflect"
	"testing"

	"github.com/yeungsean/ysq/pkg/delegate"
)

func TestIn(t *testing.T) {
	type args[T comparable] struct {
		slice []T
		want  T
	}
	tests := []struct {
		name string
		args args[int]
		want bool
	}{
		{
			name: "exists",
			args: args[int]{
				slice: []int{1, 2, 3, 4, 5, 6, 7},
				want:  3,
			},
			want: true,
		},
		{
			name: "not exists",
			args: args[int]{
				slice: []int{1, 2, 3, 4, 5, 6, 7},
				want:  100,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := In(tt.args.slice, tt.args.want); got != tt.want {
				t.Errorf("In() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInBy(t *testing.T) {
	type data struct {
		Name string
		Age  uint
	}
	dataList := []data{
		{Name: "Peter", Age: 10},
		{Name: "Andy", Age: 20},
		{Name: "Sally", Age: 15},
	}
	type args[T any] struct {
		slice []T
		fn    delegate.FuncTBool[T]
	}
	tests := []struct {
		name string
		args args[data]
		want bool
	}{
		{
			name: "exists",
			args: args[data]{
				slice: dataList,
				fn: func(d data) bool {
					return d.Name == "Peter"
				},
			},
			want: true,
		},
		{
			name: "not exists",
			args: args[data]{
				slice: dataList,
				fn: func(d data) bool {
					return d.Name == "Stupid"
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InBy(tt.args.slice, tt.args.fn); got != tt.want {
				t.Errorf("InBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAll(t *testing.T) {
	type args[T comparable] struct {
		slice []T
		want  T
	}
	tests := []struct {
		name string
		args args[int]
		want bool
	}{
		{
			name: "all equal",
			args: args[int]{
				slice: []int{2, 2, 2, 2, 2, 2, 2},
				want:  2,
			},
			want: true,
		},
		{
			name: "not equal",
			args: args[int]{
				slice: []int{1, 2, 3, 4, 5, 6, 7},
				want:  100,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := All(tt.args.slice, tt.args.want); got != tt.want {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAllBy(t *testing.T) {
	type data struct {
		Name string
		Age  uint
	}
	dataList := []data{
		{Name: "Peter", Age: 10},
		{Name: "Andy", Age: 20},
		{Name: "Sally", Age: 15},
	}
	type args[T any] struct {
		slice []T
		fn    delegate.FuncTBool[T]
	}
	tests := []struct {
		name string
		args args[data]
		want bool
	}{
		{
			name: "exists",
			args: args[data]{
				slice: dataList,
				fn: func(d data) bool {
					return d.Age < 30
				},
			},
			want: true,
		},
		{
			name: "not exists",
			args: args[data]{
				slice: dataList,
				fn: func(d data) bool {
					return d.Name == "Stupid"
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AllBy(tt.args.slice, tt.args.fn); got != tt.want {
				t.Errorf("AllBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInterface(t *testing.T) {
	type args[T comparable] struct {
		slice []T
		want  []interface{}
	}
	tests := []struct {
		name string
		args args[int]
	}{
		{
			name: "exists",
			args: args[int]{
				slice: []int{1, 2, 3, 4, 5, 6, 7},
				want:  []interface{}{1, 2, 3, 4, 5, 6, 7},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CastToInterface(tt.args.slice); !reflect.DeepEqual(got, tt.args.want) {
				t.Errorf("CastToInterface() = %v, want %v", got, tt.args.want)
			}
		})
	}
}
