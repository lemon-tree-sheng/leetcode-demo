package main

import (
	"fmt"
)

func main() {
	var a int
	for {
		_, err := fmt.Scan(&a)
		if err != nil {
			fmt.Sprintln("error input")
			return
		}

		if a == 0 {
			break
		} else {
			fmt.Printf("%d\n", a/2)
		}
	}
}
