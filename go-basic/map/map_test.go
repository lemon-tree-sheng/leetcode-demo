package slice

import (
	"fmt"
	"testing"
)

func TestFoo(t *testing.T) {
	a := make(map[int]struct{})
	a[1] = struct{}{}

	if val, ok := a[1]; ok {
		fmt.Printf("%v\n", val)
	}

	for k, v := range a {
		fmt.Printf("%d, %v\n", k, v)
	}
}
