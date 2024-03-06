package _3_

import "testing"

/*
*
给一个 res 用来存储最终的结果

设计一个递归函数

	dg(nums, start_idx, tmp_res)
		if start_idx >= len(nums) - 1
			return
		tmp_res.add(nums[start_idx])
		if is_valid(tmp_res)
			res.add(tmp_res)
			return
		for i:=start_idx + 1, i < len(nums)
			dg(nums, i, tmp_res)
*/
func permute(nums []int) [][]int {
	var res [][]int

	var track []int
	backtrack(nums, track, res)

	return res
}

func backtrack(nums []int, track []int, res [][]int) {

	if len(track) == len(nums) {
		res = append(res, track)
		return
	}

	for i := 0; i < len(nums); i++ {
		if contains(track, nums[i]) {
			continue
		}

		track = append(track, nums[i])

		backtrack(nums, track, res)

		track = track[:len(track)-1]
	}

}

func contains(slice []int, value int) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
		return false
	}
	return false
}

func TestName(t *testing.T) {
	permute(nil)
}
