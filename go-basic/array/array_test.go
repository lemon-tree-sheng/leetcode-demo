package slice

import (
	"fmt"
	"testing"
)

func TestFoo(t *testing.T) {
	var a [10]int
	println(a[0])

	var b [10][10]int
	println(b[0][9])

	for val, idx := range b {
		fmt.Printf("%d, %d\n", val, idx)
	}
}
