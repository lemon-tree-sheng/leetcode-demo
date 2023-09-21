package leetcode_demo

import "testing"

func foo(a, b int) int {
	return a + b
}

func TestFoo(t *testing.T) {
	res := foo(1, 2)
	if res != 3 {
		t.Fail()
	}
}
