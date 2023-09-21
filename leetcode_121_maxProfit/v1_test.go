package leetcode_121_maxProfit

import (
	"sheng.com/leetcode-demo/common"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	prices := []int{2, 2, 4, -2, 4}

	res := maxProfit(prices)

	println(res)
}

func maxProfit(prices []int) int {
	length := len(prices)
	if length == 0 || length == 1 {
		return 0
	}

	dp := make([][][]int, length)
	dp[0][0][0] = 0
	dp[0][1][1] = -prices[0]

	for i := 1; i < length; i++ {
		dp[i][0][0] = dp[i-1][0][0]
		dp[i][1][1] = dp[i-1][0][0] - prices[i]

		if i > 1 {
			dp[i][0][1] = common.Max(dp[i-1][1][1], dp[i-1][2][1], dp[i-1][0][1])
			dp[i][2][1] = common.Max2(dp[i-1][0][1], dp[i-1][1][1]+prices[i])
		} else {
			dp[i][0][1] = common.Max2(dp[i-1][1][1], dp[i-1][0][1])
			dp[i][2][1] = dp[i-1][1][1] + prices[i]
		}
	}

	return common.Max(dp[length-1][0][1], dp[length-1][2][1], dp[length-1][0][0])
}
