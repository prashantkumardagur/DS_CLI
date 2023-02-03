package main

import (
	"fmt"
	"reflect"
)

// Node struct for binary search tree
type BSTNode struct {
	value int
	left *BSTNode
	right *BSTNode
}
func newBSTNode(value int) *BSTNode {
	return &BSTNode{value: value}
}


// BinarySearchTree struct
type BST struct {
	root *BSTNode
	kind string
}
func NewBST() *BST {
	return &BST{kind: "int"}
}
func (t *BST) Kind() string { return t.kind }


// function to get root value of the binary search tree
func (t *BST) Root() int {
	return t.root.value
}


// function to insert a node in the binary search tree
func (t *BST) Insert(value int) {
	node := newBSTNode(value)
	if t.root == nil {
		t.root = node
		return
	}
	current := t.root
	for {
		if value < current.value {
			if current.left == nil {
				current.left = node
				break
			} else {
				current = current.left
			}
		} else {
			if current.right == nil {
				current.right = node
				break
			} else {
				current = current.right
			}
		}
	}
}


// function to search a node in the binary search tree
func (t *BST) Search(value item) bool {
	if reflect.TypeOf(value).Kind() != reflect.Int { return false }
	return searchNode(t.root, value.(int))
}
func searchNode(node *BSTNode, value int) bool {
	if node == nil { return false }
	if node.value == value { return true }
	if value < node.value {
		return searchNode(node.left, value)
	} else {
		return searchNode(node.right, value)
	}
}


// function to delete a node from the binary search tree
func (t *BST) Delete(value int) {
	deleteNode(t.root, value)
}
func minValueNode(node *BSTNode) *BSTNode {
	current := node
	for current.left != nil { current = current.left }
	return current
}
func deleteNode(node *BSTNode, value int) *BSTNode {
	if node == nil { return nil }
	if value < node.value {
		node.left = deleteNode(node.left, value)
	} else if value > node.value {
		node.right = deleteNode(node.right, value)
	} else {
		if(node.left == nil && node.right == nil) { return nil }
		if(node.left == nil) { return node.right }
		if(node.right == nil) { return node.left }
		temp := minValueNode(node.right)
		node.value = temp.value
		node.right = deleteNode(node.right, temp.value)
	}
	return node
}



// function to print inorder traversal of the binary search tree
func (t *BST) Inorder() {
	fmt.Print("cli_out> Inorder: ")
	inorder(t.root)
	fmt.Println()
}
func inorder(node *BSTNode) {
	if node == nil { return }
	inorder(node.left)
	fmt.Printf("%d ", node.value)
	inorder(node.right)
}


// function to print preorder traversal of the binary search tree
func (t *BST) Preorder() {
	fmt.Print("cli_out> Preorder: ")
	preorder(t.root)
	fmt.Println()
}
func preorder(node *BSTNode) {
	if node == nil { return }
	fmt.Printf("%d ", node.value)
	preorder(node.left)
	preorder(node.right)
}


// function to print postorder traversal of the binary search tree
func (t *BST) Postorder() {
	fmt.Print("cli_out> Postorder: ")
	postorder(t.root)
	fmt.Println()
}
func postorder(node *BSTNode) {
	if node == nil { return }
	postorder(node.left)
	postorder(node.right)
	fmt.Printf("%d ", node.value)
}