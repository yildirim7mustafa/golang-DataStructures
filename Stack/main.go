package main

import (
	"Stack/stack"
)

func main() {
	s := stack.Stack{}
	s.IsEmpty()
	s.Push(55)
	s.Push(45)
	s.Push(35)
	s.Print()
	s.Pop()
	s.Print()
	s.Push(75)
	print(s.IsEmpty())
	s.Push(85)
	s.Print()
	s.Top()

}
