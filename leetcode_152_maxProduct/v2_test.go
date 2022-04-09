package leetcode_152_maxProduct

import (
	"sheng.com/go-demo/common"
	"testing"
)

var res int

func TestMax(t *testing.T) {
	nums := []int{2, 3, -2, 4}

    res = nums[0]

	maxProductV2(nums, len(nums) - 1)

	println(res)
}

func maxProductV2(nums []int, idx int) (int, int) {
    if idx == 0 {
        return nums[0], nums[0]
    }

    ma, mi := maxProductV2(nums, idx - 1)
    maNew := common.Max(nums[idx], ma * nums[idx], mi * nums[idx])
    miNew := common.Min(nums[idx], ma * nums[idx], mi * nums[idx])

    if res < maNew {
        res = maNew
    }

    return maNew, miNew
}