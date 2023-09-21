package leetcode_283

import (
	"testing"
)

func moveZeroes(nums []int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			for j := i + 1; j < len(nums); j++ {
				swap(i, j, nums)
			}
		}
	}
}

func swap(i int, j int, nums []int) {
	tmp := nums[j]
	nums[j] = nums[i]
	nums[i] = tmp
}

func TestMoveZeroes(t *testing.T) {

}
