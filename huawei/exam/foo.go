package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	n, m := 0, 0
	_, err := fmt.Scan(&n, &m)
	if err != nil {
		return
	}

	input := make([][]int, 0)
	for i := 0; i < n; i++ {
		reader := bufio.NewReader(os.Stdin)
		inputs, _ := reader.ReadString('\n')
		strArr := strings.Fields(inputs)

		var tmpInput []int
		for k := 0; k < len(strArr); k++ {
			num, err := strconv.Atoi(strArr[k])
			if err != nil {
				return
			}
			tmpInput = append(tmpInput, num)
		}
		input = append(input, tmpInput)
	}

	output := make([][]int, 0, 5)
	for i := 0; i < m; i++ {
		for k := 0; k < n; k++ {
			output[i][k*2] = input[k][i*2]
			output[i][k*2+1] = -input[k][i*2+1]
		}
	}

	for i := 0; i < m; i++ {
		for k := 0; k < len(output[i]); k++ {
			if k < len(output[i])-1 {
				fmt.Printf("%d ", output[i][k])
			} else {
				fmt.Printf("%d", output[i][k])
			}
		}
		fmt.Printf("\n")
	}
}
