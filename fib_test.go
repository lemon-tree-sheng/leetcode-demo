package leetcode_demo

import "testing"

// LCR 126. 斐波那契数
func fib(n int) int {
	if n <= 1 {
		return n
	}

	p := 0
	q := 1
	r := 0

	for i := 2; i <= n; i++ {
		r = (p + q) % 1000000007
		p = q
		q = r
	}
	return r
}

func TestFib(t *testing.T) {
	res := fib(5)
	if res != 5 {
		t.Fail()
	}
}
