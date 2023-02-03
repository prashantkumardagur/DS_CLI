package main

import (
	"fmt"
)


// Node struct for linked list
type LLNode struct {
	value item
	next *LLNode
}
func newLLNode(value item) *LLNode {
	return &LLNode{value: value}
}

// move the pointer to the next node
func (n *LLNode) NextNode() *LLNode { return n.next }

// move the pointer to the previous node
func (n *LLNode) PrevNode() *LLNode {
	fmt.Println("cli_out> Previous node is not available in linked list")
	return nil 
}

// get the value of the current node
func (n *LLNode) GetValue() item { return n.value }

// set the value of the current node
func (n *LLNode) SetValue(value item) { n.value = value }

// function to check if pointer is the end of the linked list
func (n *LLNode) IsEnd() bool {
	if n.next == nil { return true }
	return false
}



// ===================================================================================================



// LinkedList struct
type LinkedList struct {
	head *LLNode
	tail *LLNode
	kind string
}
func NewLinkedList(varkind string) *LinkedList {
	return &LinkedList{kind: varkind}
}
func (l *LinkedList) Kind() string { return l.kind }


// function to add a node in back of the linked list
func (l *LinkedList) AddBack(value item) {
	node := newLLNode(value)
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.next = node
		l.tail = node
	}
}


// function to add a node in front of the linked list
func (l *LinkedList) AddFront(value item) {
	node := newLLNode(value)
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		node.next = l.head
		l.head = node
	}
}


// function to delete a node from the back of the linked list
func (l *LinkedList) DeleteBack() {
	if l.head == nil { return } 
	if l.head == l.tail { 
		l.head = nil
		l.tail = nil
		return
	}
	current := l.head
	for current.next != l.tail { current = current.next }
	current.next = nil
	l.tail = current
}


// function to delete a node from the front of the linked list
func (l *LinkedList) DeleteFront() {
	if l.head == nil { return }
	if l.head == l.tail { 
		l.head = nil
		l.tail = nil
		return
	}
	l.head = l.head.next
}


// function to return front node of the linked list
func (l *LinkedList) Front() *LLNode {
	return l.head
}


// function to return back node of the linked list
func (l *LinkedList) Back() *LLNode {
	return l.tail
}


// function to insert a node after the node with given value
func (l *LinkedList) InsertAfter(value item, newValue item) {
	if l.head == nil { return }
	current := l.head
	for current.value != value {
		if current.IsEnd() { return }
		current = current.next
	}
	node := newLLNode(newValue)
	node.next = current.next
	current.next = node
}


// function to search for a node with given value
func (l *LinkedList) Search(value item) bool {
	if l.head == nil { return false }
	current := l.head
	for current != nil {
		if current.value == value { return true }
		current = current.next
	}
	return false
}


// function to print the linked list
func (l *LinkedList) Print() {
	if l.head == nil { return }
	fmt.Print("cli_out> Head ")
	current := l.head
	for current != nil {
		fmt.Printf("-> %v ", current.value)
		current = current.next
	}
	fmt.Printf("\n")
}