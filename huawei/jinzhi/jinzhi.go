package main

import (
	"fmt"
)

func main() {
	a := 0
	_, err := fmt.Scan(&a)
	if err != nil {

	}
	fmt.Printf("%d\n", a)
}
