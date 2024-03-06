package _3_

import "testing"

func TestName2(t *testing.T) {
	nums := []int{1, 2, 3}
	permute2(nums)
}

func permute2(nums []int) [][]int {
	var res [][]int
	var track []int

	backtrack2(nums, track, res)

	return res
}

func backtrack2(nums []int, track []int, res [][]int) {

	if len(track) == len(nums) {
		res = append(res, track)
		return
	}

	for i := 0; i < len(nums); i++ {
		if contains(track, nums[i]) {
			continue
		}
		track = append(track, nums[i])
		backtrack2(nums, track, res)
		track = track[:len(track)-1]
	}
}
