package ysq

import "testing"

func TestQueryIter(t *testing.T) {
	q := FromElement(1, 2, 3, 4, 5, 6)
	q.Iter(q.Next(), func(i int) IterContinue {
		return IterContinueYes
	})
}
