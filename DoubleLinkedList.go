package main

import (
	"fmt"
)

// Node struct for double linked list
type DLLNode struct {
	value item
	next *DLLNode
	prev *DLLNode
}
func newDLLNode(value item) *DLLNode {
	return &DLLNode{value: value}
}

// move the pointer to the next node
func (n *DLLNode) NextNode() *DLLNode { return n.next }

// move the pointer to the previous node
func (n *DLLNode) PrevNode() *DLLNode { return n.prev }

// get the value of the current node
func (n *DLLNode) GetValue() item { return n.value }

// set the value of the current node
func (n *DLLNode) SetValue(value item) { n.value = value }

// function to check if pointer is the end of the double linked list
func (n *DLLNode) IsEnd() bool {
	if n.next == nil { return true }
	return false
}

// function to check if pointer is the head of the double linked list
func (n *DLLNode) IsHead() bool {
	if n.prev == nil { return true }
	return false
}


// ===================================================================================================


// DoubleLinkedList struct
type DoubleLinkedList struct {
	head *DLLNode
	tail *DLLNode
	kind string
}
func NewDoubleLinkedList(varkind string) *DoubleLinkedList {
	return &DoubleLinkedList{kind: varkind}
}
func (l *DoubleLinkedList) Kind() string { return l.kind }


// function to add a node in back of the double linked list
func (l *DoubleLinkedList) AddBack(value item) {
	node := newDLLNode(value)
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.next = node
		node.prev = l.tail
		l.tail = node
	}
}


// function to add a node in front of the double linked list
func (l *DoubleLinkedList) AddFront(value item) {
	node := newDLLNode(value)
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		node.next = l.head
		l.head.prev = node
		l.head = node
	}
}


// function to delete a node from the back of the double linked list
func (l *DoubleLinkedList) DeleteBack() {
	if l.head == nil { return } 
	if l.head == l.tail { 
		l.head = nil
		l.tail = nil
		return
	}
	l.tail = l.tail.prev
	l.tail.next = nil
}


// function to delete a node from the front of the double linked list
func (l *DoubleLinkedList) DeleteFront() {
	if l.head == nil { return } 
	if l.head == l.tail { 
		l.head = nil
		l.tail = nil
		return
	}
	l.head = l.head.next
	l.head.prev = nil
}


// function to insert a node after the node with given value
func (l *DoubleLinkedList) InsertAfter(value item, after item) {
	if l.head == nil { return }
	current := l.head
	for current != nil {
		if current.value == after {
			node := newDLLNode(value)
			node.next = current.next
			current.next.prev = node
			current.next = node
			node.prev = current
			return
		}
		current = current.next
	}
}


// function to return front node of the double linked list
func (l *DoubleLinkedList) Front() *DLLNode {
	return l.head
}


// function to return back node of the double linked list
func (l *DoubleLinkedList) Back() *DLLNode {
	return l.tail
}


// function to print the double linked list
func (l *DoubleLinkedList) Print() {
	if(l.head == nil) { return }
	fmt.Print("cli_out> Head ")
	current := l.head
	for current != nil {
		fmt.Printf("-> %v ", current.value)
		current = current.next
	}
	fmt.Printf("\n")
}


// functtion to search a node with given value
func (l *DoubleLinkedList) Search(value item) bool {
	current := l.head
	for current != nil {
		if current.value == value { return true }
		current = current.next
	}
	return false
}