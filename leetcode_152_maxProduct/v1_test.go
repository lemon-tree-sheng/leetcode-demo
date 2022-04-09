package leetcode_152_maxProduct

import (
    "sheng.com/go-demo/common"
    "testing"
)

func TestMaxProduct(t *testing.T) {
    // nums := []int{2, 3, -2, 4}
    nums := []int{2, 2, 2, 2}

    res := maxProduct(nums)

    println(res)
}

func maxProduct(nums []int) int {
    length := len(nums)
    ma := make([]int, length)
    mi := make([]int, length)

    ma[0] = nums[0]
    mi[0] = nums[0]
    res := ma[0]

    for i := 1; i < length; i++ {
        ma[i] = common.Max(nums[i], mi[i - 1] * nums[i], ma[i - 1] * nums[i])
        mi[i] = common.Min(nums[i], mi[i - 1] * nums[i], ma[i - 1] * nums[i])

        if res < ma[i] {
            res = ma[i]
        }
    }

    return res
}