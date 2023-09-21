package leetcode_demo

import "testing"

// LCR 126. 斐波那契数
func fib(n int) int {
	if n <= 1 {
		return n
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % 1000000007
	}
	return dp[n]
}

func TestFib(t *testing.T) {
	res := fib(5)
	if res != 5 {
		t.Fail()
	}
}
