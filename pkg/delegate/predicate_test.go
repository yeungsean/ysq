package delegate

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartialFunc(t *testing.T) {
	assert.Panics(t, func() {
		PartialFunc[int, int](nil, 1, 2)
	})

	var testSumF = func(arg1, arg2, arg3 int) int {
		return arg1 + arg2 + arg3
	}

	var testSum1F = func(arg1, arg2, arg3, arg4 int) int {
		return arg1 + arg2 + arg3 + arg4
	}

	delayCall := PartialFunc[int, int](testSumF, 1, 2)
	for i := 0; i < 5; i++ {
		res := delayCall(i)
		assert.Equal(t, 3+i, res)
	}

	delayCall1 := PartialFunc[int, int](testSum1F, 1, 2)
	for i := 0; i < 5; i++ {
		res := delayCall1(1, i)
		assert.Equal(t, 3+i+1, res)
	}
}

func TestFuncInt3Partial(t *testing.T) {
	var testCheckSumF = func(arg1, arg2, arg3 int) bool {
		return arg1+arg2 == arg3
	}

	var f FuncInt3Bool = testCheckSumF
	partial := f.Partial(1)
	for i := 0; i < 3; i++ {
		assert.True(t, partial(i, i+1))
	}
}

func TestFuncString3Partial(t *testing.T) {
	var testCheckSumF = func(arg1, arg2, arg3 string) bool {
		return arg1+arg2 == arg3
	}

	var f FuncString3Bool = testCheckSumF
	partial := f.Partial("1")
	for i := 0; i < 3; i++ {
		assert.True(t, partial(strconv.Itoa(i), fmt.Sprintf("1%d", i)))
	}
}
