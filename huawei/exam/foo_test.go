package main

import (
	"fmt"
)

func foo() {
	var nameSeq []string
	var nameMap map[string][]int
	var gradeMap map[string]string
	var excellentFlag bool
	var goodFlag bool
	var badFlag bool

	for name, stepArr := range nameMap {
		lgt30000Cnt := 0
		lgt10000Cnt := 0
		for val, _ := range stepArr {
			if val > 30000 {
				lgt30000Cnt++
			}
			if lgt30000Cnt > 4 {
				gradeMap[name] = "excellent"
			}

			if val > 10000 {
				lgt10000Cnt++
			}
			if lgt10000Cnt > 15 {
				gradeMap[name] = "excellent"
			}
		}

	}

	if !excellentFlag {
		fmt.Printf("excellent is null\n")
	}
	if !goodFlag {
		fmt.Printf("good is null\n")
	}
	if !badFlag {
		fmt.Printf("bad is null\n")
	}
}
