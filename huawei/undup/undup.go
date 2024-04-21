package main

import (
	"fmt"
	"sort"
)

func main() {
	foo()

	//res := []int{1, 6, 8, 2, 3, 1}
	//quickSort(res, 0, len(res)-1)
	//QuickSort(res, 0, len(res)-1)

	//sort.Slice(res, func(i, j int) bool {
	//	return res[i] <= res[j]
	//})

	//for i := 0; i < len(res); i++ {
	//	println(res[i])
	//}
}

func foo() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {

	}

	var res []int
	memo := make(map[int]bool)
	var a int
	for ; n > 0; n-- {
		_, err := fmt.Scan(&a)
		if err != nil {

		}

		if _, ok := memo[a]; !ok {
			memo[a] = true
			res = append(res, a)
		}
	}

	//quickSort(res, 0, len(res)-1)
	sort.Slice(res, func(i, k int) bool {
		return res[i] <= res[k]
	})

	for i := 0; i < len(res); i++ {
		fmt.Printf("%d\n", res[i])
	}
}

func quickSort(nums []int, left, right int) {
	tmp := nums[left]
	l := left
	r := right
	for l < r {
		for l < r {
			if tmp < nums[r] {
				r--
			} else {
				fuzhi(nums, l, r)
				l++
			}
		}

		for l < r {
			if tmp >= nums[l] {
				l++
			} else {
				fuzhi(nums, r, l)
				r--
			}
		}
	}
	nums[l] = tmp

	quickSort(nums, left, l)
	quickSort(nums, l+1, right)
}

func fuzhi(nums []int, targetIdx, curIdx int) {
	nums[targetIdx] = nums[curIdx]
}

// /////
func QuickSort(nums []int, left int, right int) {
	if left < right {
		pivot := partition(nums, left, right)
		QuickSort(nums, left, pivot-1)
		QuickSort(nums, pivot+1, right)
	}
}

func partition(arr []int, left, right int) int {
	pivot := arr[right]
	i := left - 1
	for j := left; j < right; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[right] = arr[right], arr[i+1]
	return i + 1
}
