package multiply

import (
	"fmt"
)

func main() {
	a := ""
	b := ""
	_, err := fmt.Scan(&a, &b)
	if err != nil {

	}

	for _, valA := range a {
		flag := 0
		intC := 0
		tmps := make([]byte, 0)
		for _, valB := range b {
			intA := int(valA)
			intB := int(valB)

			intC = (intA*intB + flag) % 10
			flag = (intA*intB + flag) / 10

			tmps = append(tmps, byte(intC))
		}
	}

}
