package lang

import (
	"testing"
)

func TestListFiles(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []int{1, 2, 3, 4, 0}

	AssertValue("a,a", EqualArrays(a, a), true)
	AssertValue("a,x", EqualArrays(a, a[0:2]), false)
	AssertValue("a,b", EqualArrays(a, b), false)
}
