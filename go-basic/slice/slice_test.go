package slice

import "testing"

/**
初始化
二维初始化
追加元素
索引元素
是否原地修改
*/

func TestFoo(t *testing.T) {
	var a []int
	a = append(a, 10)
	println(a[0])

	a = make([]int, 1)
	println(a[0])
	a[0] = 1

	for val, idx := range a {
		println(val)
		println(idx)
	}

	for i := 0; i < len(a); i++ {
		println(a[i])
	}

	c := make([][]int, 0)
	c = append(c, []int{1, 2})
}
