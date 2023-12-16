package leetcode_demo

import (
	"fmt"
	"strconv"
	"testing"
)

func addStrings(num1 string, num2 string) string {
	sli1 := string2Slice(num1)
	sli2 := string2Slice(num2)

	res := 0
	if len(sli1) > len(sli2) {
		tmp := 0
		for i := 0; i <= len(sli1); i++ {
			nn := numN(10, i) // 这里用了乘积的方式不是一个好的方式
			// 使用每一位计算好结果之后，得到个位和进位的值即可
			if i == 1 {
				res += tmp
			} else {
				res += (10 * i) * tmp
			}

			tmp = 0
			num1 := 0
			if len(sli1)-1-i >= 0 {
				num1 = sli1[len(sli1)-1-i]
			}
			num2 := 0
			if len(sli2)-1-i >= 0 {
				num2 = sli2[len(sli2)-1-i]
			}

			tmp = num1 + num2
		}
	}

	return strconv.Itoa(res)
}

func string2Slice(num1 string) []int {
	res := make([]int, 0, len(num1))
	for _, char := range num1 {
		res = append(res, int(char-'0'))
	}
	return res
}

func TestAddStrings(t *testing.T) {
	res := addStrings("1234", "789")
	if res != "2023" {
		t.Errorf(fmt.Sprintf("res is : %s", res))
	}
}
