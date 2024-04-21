package foo_tet

import (
	"fmt"
	"testing"
)

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)-1] // 选择最后一个元素作为基准值
	i := 0

	for j := 0; j < len(arr)-1; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i] // 将小于基准值的元素交换到左边
			i++
		}
	}
	arr[i], arr[len(arr)-1] = arr[len(arr)-1], arr[i] // 将基准值放到正确的位置

	left := quickSort(arr[:i])    // 对左侧子数组进行递归排序
	right := quickSort(arr[i+1:]) // 对右侧子数组进行递归排序

	return append(append(left, arr[i]), right...) // 合并排序结果
}

func TestFoo(t *testing.T) {
	arr := []int{3, 7, 1, 5, 2, 6, 8, 4}
	fmt.Println("Before sorting:", arr)
	fmt.Println("After sorting:", quickSort(arr))
}
