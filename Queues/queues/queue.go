package queues

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(data int) {
	q.items = append(q.items, data)
}

func (q *Queue) Dequeue() {
	q.items = q.items[1:]
}

func (q *Queue) Print() {
	for _, item := range q.items {
		fmt.Print(item, " ")
	}
	println()
}

func (q *Queue) IsEmpty() bool {
	if q.items == nil {
		return true
	}
	return false
}
