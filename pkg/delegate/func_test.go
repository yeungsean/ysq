package delegate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testStart5AddOther(arg1, arg2 int) int {
	return arg1 + arg2
}

func TestFunc2Partial(t *testing.T) {
	var f2 Func2[int, int, int] = testStart5AddOther
	delayCall := f2.Partial(5)
	func() {
		res := delayCall(10)
		assert.Equal(t, 15, res)

		res = delayCall(-10)
		assert.Equal(t, -5, res)
	}()
}

func testStart5AddSub(arg1, arg2, arg3 int) int {
	return arg1 + arg2 - arg3
}

func TestFunc3Partial(t *testing.T) {
	var f3 Func3[int, int, int, int] = testStart5AddSub
	func() {
		delayCall := f3.Partial(5)
		res := delayCall(10, 1)
		assert.Equal(t, 14, res)

		res = delayCall(-10, 4)
		assert.Equal(t, -9, res)
	}()

	func() {
		delayCall := f3.Partial2(5, 10)
		res := delayCall(1)
		assert.Equal(t, 14, res)
	}()
}
