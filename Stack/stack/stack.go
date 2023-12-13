package stack

import "fmt"

type Stack struct {
	Items []int
}

func (s *Stack) Push(data int) {
	s.Items = append(s.Items, data)
}

func (s *Stack) Pop() {
	if s == nil {
		return
	}
	s.Items = s.Items[:len(s.Items)-1]
}

func (s Stack) Print() {
	fmt.Println(s.Items)
}

func (s *Stack) Top() {
	if s.Items == nil {
		fmt.Println("Stack is empty")
	} else {
		fmt.Print(s.Items[len(s.Items)-1])
	}
}

func (s *Stack) IsEmpty() bool {
	if s.Items == nil {
		fmt.Println("Emtpy Stack")
		return true
	}
	return false
}
