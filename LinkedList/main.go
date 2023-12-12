package main

import "LinkedList/list"

func main() {

	l := list.List{}
	l.Print() // head
	l.Append(5)
	l.Append(4)
	l.Print() // head 5 4
	l.Prepend(1)
	l.Print()
	println("Length: ", l.Length())
	println(l.Search(5))
	l.DeleteFirst()
	l.Print()
	l.DeleteLast()
	l.Print()
	l.Append(15)
	l.Append(24)
	l.Print()
	l.Delete(15)
	l.Print()
}
