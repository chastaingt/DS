package main

import "fmt"

type queue struct {
	items []int
}

func (q *queue) enqueue(value int) {
	q.items = append(q.items, value)
}

func (q *queue) dequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

func main() {
	queue := queue{}
	queue.enqueue(10)
	fmt.Println(queue)
	queue.dequeue()
	fmt.Println(queue)
}
