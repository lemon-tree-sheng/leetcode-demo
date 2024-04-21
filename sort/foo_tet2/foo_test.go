package foo_tet

import (
	"fmt"
	"testing"
)

func quickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	left := 0
	right := len(nums) - 1
	pivot := nums[right]

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < pivot {
			nums[left], nums[i] = nums[i], nums[left]
			left++
		}
	}

	nums[left], nums[right] = nums[right], nums[left]

	lefts := quickSort(nums[:left])
	rights := quickSort(nums[left+1:])

	return append(append(lefts, nums[left]), rights...)
}

func TestFoo(t *testing.T) {
	arr := []int{3, 7, 1, 5, 2, 6, 8, 4}
	fmt.Println("Before sorting:", arr)
	fmt.Println("After sorting:", quickSort(arr))
}

func TestBar(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums[0:2])

	left := 0
	right := 1
	nums[left], nums[right] = nums[right], nums[left]

	fmt.Println(left)
	fmt.Println(right)

	fmt.Println(nums)
}
