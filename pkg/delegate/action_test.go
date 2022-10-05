package delegate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAction2Partial(t *testing.T) {
	hasCalled := map[int]bool{
		10: false,
		20: false,
	}
	tmpAction2 := func(arg1, arg2 int) {
		hasCalled[arg1+arg2] = true
	}

	func() {
		var f Action2[int, int] = tmpAction2
		delayCall := f.Partial(5)
		delayCall(5)
		delayCall(15)
		assert.True(t, hasCalled[10])
		assert.True(t, hasCalled[20])
	}()
}

func TestAction3Partial(t *testing.T) {
	hasCalled := map[int]bool{
		10: false,
		20: false,
		30: false,
	}
	tmpAction3 := func(arg1, arg2, arg3 int) {
		hasCalled[arg1+arg2+arg3] = true
	}

	func() {
		var f Action3[int, int, int] = tmpAction3
		delayCall := f.Partial(5)
		delayCall(5, 0)
		delayCall(6, 9)
		delayCall(11, 14)
		assert.True(t, hasCalled[10])
		assert.True(t, hasCalled[20])
		assert.True(t, hasCalled[30])
	}()
}

func TestAction3Partial2(t *testing.T) {
	hasCalled := map[int]bool{
		10: false,
		20: false,
		30: false,
	}
	tmpAction3 := func(arg1, arg2, arg3 int) {
		hasCalled[arg1+arg2+arg3] = true
	}

	func() {
		var f Action3[int, int, int] = tmpAction3
		delayCall := f.Partial2(5, 9)
		delayCall(-4)
		delayCall(6)
		delayCall(16)
		assert.True(t, hasCalled[10])
		assert.True(t, hasCalled[20])
		assert.True(t, hasCalled[30])
	}()
}
