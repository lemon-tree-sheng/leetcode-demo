package multiply

import (
	"testing"
)

// Multiply function to multiply two large integers represented as strings
func Multiply(num1 string, num2 string) string {
	len1 := len(num1)
	len2 := len(num2)
	result := make([]byte, len1+len2)

	for i := len1 - 1; i >= 0; i-- {
		carry := 0
		n1 := int(num1[i] - '0')

		for j := len2 - 1; j >= 0; j-- {
			n2 := int(num2[j] - '0')
			sum := n1*n2 + carry + int(result[i+j+1])
			carry = sum / 10
			result[i+j+1] = byte(sum%10 + '0')
		}

		if carry > 0 {
			result[i] = byte(carry + '0')
		}
	}

	// remove leading zeros
	start := 0
	for start < len(result) && result[start] == '0' {
		start++
	}

	return string(result[start:])
}

func TestFoo(t *testing.T) {
	m := make(map[string]int)
	delete(m, "a")
	intMax := 1<<31 - 1
	intMin := ^intMax
	println(intMin)
}
