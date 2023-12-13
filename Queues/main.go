package main

import "Queues/queues"

func main() {

	q := queues.Queue{}
	println(q.IsEmpty())
	q.Enqueue(55)
	q.Enqueue(45)
	q.Enqueue(35)
	q.Print()
	q.Dequeue()
	q.Dequeue()
	q.Print()
	print(q.IsEmpty())
}
