package _0_

import (
	"fmt"
	"testing"
)

func isValid(s string) bool {

	return false

}

func TestStack(t *testing.T) {
	s := new(Stack)
	fmt.Println(s.size())

	s.push("a")
	s.push("b")
	s.push("c")

	fmt.Println(s.size())

	fmt.Println(s.pop())
	fmt.Println(s.pop())
	fmt.Println(s.pop())

	fmt.Println(s.isEmpty())
}

type Stack struct {
	data []string
}

func (s *Stack) pop() string {
	if len(s.data) == 0 {
		return ""
	}
	val := s.data[len(s.data)-1]
	s.data = s.data[0 : len(s.data)-1]
	return val
}

func (s *Stack) push(newString string) {
	s.data = append(s.data, newString)
}

func (s *Stack) isEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) size() int {
	return len(s.data)
}
