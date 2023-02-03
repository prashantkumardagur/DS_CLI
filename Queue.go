package main

import (
	"fmt"
)


// struct for queue
type Queue struct {
	arr []item
	size int
	kind string
}
func NewQueue(varkind string) *Queue {
	return &Queue{kind: varkind}
}
func (q *Queue) Kind() string { return q.kind }


// function to push a value into the queue
func (q *Queue) Push(value item) {
	q.arr = append(q.arr, value)
	q.size++
}


// function to pop a value from the queue
func (q *Queue) Pop() item {
	if q.size == 0 { return -1 }
	value := q.arr[0]
	q.arr = q.arr[1:]
	q.size--
	return value
}


// function to get the front value of the queue
func (q *Queue) Front() item {
	if q.size == 0 { return -1 }
	return q.arr[0]
}


// function to get the back value of the queue
func (q *Queue) Back() item {
	if q.size == 0 { return -1 }
	return q.arr[q.size-1]
}


// function to get the size of the queue
func (q *Queue) Size() int {
	return q.size
}


// function to check if the queue is empty
func (q *Queue) IsEmpty() bool {
	return q.size == 0
}


// function to print the queue
func (q *Queue) Print() {
	fmt.Printf("cli_out> Queue: Front ")
	for i := 0; i < q.size; i++ {
		fmt.Printf("<- %v ", q.arr[i])
	}
	fmt.Printf("<- Back\n")
}