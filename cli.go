package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)


// ======================================================================================================

type item interface{} // empty interface for generic types

type ds interface {
	Kind() string
	Print()
}

type variable struct {
	ds string
	value item
}

type CLI struct {
	vars map[string]variable
}

// function to create a new cli
func NewCLI() *CLI {   return &CLI{ vars: make(map[string]variable) }   }


// ======================================================================================================


// function to run the cli
func (c *CLI) Run() {
	in := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("\ncli_in> ")
		cmd, _ := in.ReadString('\n')   // read the command
		c.execute( strings.Fields(cmd) )  // execute the command
	}
}

// function to seed the cli with some commands
func (c *CLI) Seed( commands []string ) {
	fmt.Println("====================================================================")
	fmt.Printf("cli_out> Executing seed commands...\n\n")
	for _, cmd := range commands {  
		c.execute( strings.Fields(cmd) )
	}
	fmt.Printf("\ncli_out> Seed commands executed\n")
	fmt.Println("====================================================================")
}


// ======================================================================================================


// function to create a new data structure
func (c *CLI) createDS(ds, kind, name string) {
	switch ds {
		case "bst": c.vars[name] = variable{ds: "bst", value: NewBST()}
		case "stack": c.vars[name] = variable{ds: "stack", value: NewStack(kind)}
		case "queue": c.vars[name] = variable{ds: "queue", value: NewQueue(kind)}
		case "linkedlist": c.vars[name] = variable{ds: "linkedlist", value: NewLinkedList(kind)}
		case "doublelinkedlist": c.vars[name] = variable{ds: "doublelinkedlist", value: NewDoubleLinkedList(kind)}
		default: fmt.Println("cli_out> Invalid data structure")
	}
}

// function to create a new variable
func (c *CLI) createVar(kind, name string) {
	switch kind {
		case "int": c.vars[name] = variable{ds: "int", value: 0}
		case "float": c.vars[name] = variable{ds: "float", value: 0.0}
		case "string": c.vars[name] = variable{ds: "string", value: ""}
		case "bool": c.vars[name] = variable{ds: "bool", value: false}
		case "ptr": c.vars[name] = variable{ds: "ptr", value: nil}
		default: fmt.Println("cli_out> Invalid data type")
	}
}

// function to print the value of a variable
func (c *CLI) PrintVar(name string) {
	entry, ok := c.vars[name]
	if !ok { fmt.Println("cli_out> Variable not found"); return }
	fmt.Println("cli_out> ", entry.value)
}


// ======================================================================================================


// function to execute the command
func (c *CLI) execute(cmd []string) {

	// check if valid data structure is provided
	if ( 
	// Stack and Queue commands
	cmd[0] == "push" || cmd[0] == "pop" || cmd[0] == "top" || cmd[0] == "front" || cmd[0] == "back" || cmd[0] == "isEmpty" ||
	// Linked List commands
	cmd[0] == "addFront" || cmd[0] == "addBack" || cmd[0] == "deleteFront" || cmd[0] == "deleteBack" ||
	// Linked pointer commands
	cmd[0] == "ptrFront" || cmd[0] == "ptrBack" || cmd[0] == "ptrNext" || cmd[0] == "ptrPrev" || cmd[0] == "isEnd" || cmd[0] == "isHead" ||
	// Binary Search Tree commands
	cmd[0] == "insert" || cmd[0] == "delete" || cmd[0] == "root" || cmd[0] == "inorder" || cmd[0] == "postorder" || cmd[0] == "preorder" ||
	// General commands
	cmd[0] == "print" || cmd[0] == "search" ){
		if _, ok := c.vars[cmd[1]]; !ok { fmt.Println("cli_out> Data structure not found"); return }
	}


	// switch on the command and execute the appropriate function
	switch cmd[0] {

	// ===== General commands =======================================================

	case "exit": {fmt.Println(); os.Exit(0) } // exit the cli

	case "var": {  // create a new variable
		if len(cmd) == 4 { c.createDS(cmd[1], cmd[2], cmd[3]) 
		} else if len(cmd) == 3 { c.createVar(cmd[1], cmd[2]) }
	}

	case "print": {  // print the value of a variable or data structure
		switch c.vars[cmd[1]].ds {
			case "stack": fallthrough
			case "queue": fallthrough
			case "linkedlist": fallthrough
			case "doublelinkedlist": c.vars[cmd[1]].value.(ds).Print()
			case "bst": c.vars[cmd[1]].value.(*BST).Inorder()
			default: c.PrintVar(cmd[1])
		}
	}

	case "search": {  // search for a value in a data structure
		switch c.vars[cmd[1]].ds {
			case "bst": c.vars[cmd[1]].value.(*BST).Search(cmd[2])
			case "linkedlist": c.vars[cmd[1]].value.(*LinkedList).Search(cmd[2])
			case "doublelinkedlist": c.vars[cmd[1]].value.(*DoubleLinkedList).Search(cmd[2])
			default: fmt.Println("cli_out> Invalid data structure")
		}
	}



	// ===== Data structure manipulation commands ===================================


	// Stack and Queue commands-------------------------------------------------------

	case "push": {  // push a value into a stack or queue
		var temp = convert(cmd[2], c.vars[cmd[1]].value.(ds).Kind())  // convert the value to the correct type
		if(temp == nil) { fmt.Println("cli_out> Invalid data type"); return }  // if the value is invalid, return
		switch c.vars[cmd[1]].ds {
			case "stack": c.vars[cmd[1]].value.(*Stack).Push(temp)
			case "queue": c.vars[cmd[1]].value.(*Queue).Push(temp)
			default: fmt.Println("cli_out> Invalid data structure")
		}
	}

	case "pop": {  // pop a value from a stack or queue
		var temp item
		switch c.vars[cmd[1]].ds {  // switch to choose the correct data structure
			case "stack": temp = c.vars[cmd[1]].value.(*Stack).Pop()
			case "queue": temp = c.vars[cmd[1]].value.(*Queue).Pop()
			default: { fmt.Println("cli_out> Invalid data structure"); return }
		}
		entry, ok := c.vars[cmd[2]]  // check if the variable exists
		if !ok { fmt.Println("cli_out> Variable not found"); return }
		if entry.ds != c.vars[cmd[1]].value.(ds).Kind() { fmt.Println("cli_out> Data type mismatch"); return }  // check if the data type matches
		entry.value = temp
		c.vars[cmd[2]] = entry
	}

	case "isEmpty": {  // check if a stack or queue is empty
		var temp bool
		switch c.vars[cmd[1]].ds {  // switch to choose the correct data structure
			case "stack": temp = c.vars[cmd[1]].value.(*Stack).IsEmpty()
			case "queue": temp = c.vars[cmd[1]].value.(*Queue).IsEmpty()
			default: { fmt.Println("cli_out> Invalid data structure"); return }
		}
		entry, ok := c.vars[cmd[2]]  // check if the variable exists
		if !ok { fmt.Println("cli_out> Variable not found"); return }
		if entry.ds != "bool" { fmt.Println("cli_out> Data type mismatch"); return } // check if the data type matches
		entry.value = temp
		c.vars[cmd[2]] = entry
	}
 
	case "top": {  // get the top value of a stack and store it in a variable
		var temp item
		if c.vars[cmd[1]].ds == "stack" { temp = c.vars[cmd[1]].value.(*Stack).Top() // check if the data structure is a stack
		}	else { fmt.Println("cli_out> Invalid data structure"); return }
		entry, ok := c.vars[cmd[2]]		// check if the variable exists
		if !ok { fmt.Println("cli_out> Variable not found"); return }
		if entry.ds != c.vars[cmd[1]].value.(ds).Kind() { fmt.Println("cli_out> Data type mismatch"); return } // check if the data type matches
		entry.value = temp
		c.vars[cmd[2]] = entry
	}

	case "front": {  // get the front value of a queue and store it in a variable
		var temp item
		if c.vars[cmd[1]].ds == "queue" { temp = c.vars[cmd[1]].value.(*Queue).Front()  // check if the data structure is a queue
		} else { fmt.Println("cli_out> Invalid data structure"); return }
		entry, ok := c.vars[cmd[2]]  // check if the variable exists
		if !ok { fmt.Println("cli_out> Variable not found"); return }
		if entry.ds != c.vars[cmd[1]].value.(ds).Kind() { fmt.Println("cli_out> Data type mismatch"); return }  // check if the data type matches
		entry.value = temp
		c.vars[cmd[2]] = entry
	}

	case "back": {  // get the back value of a queue and store it in a variable
		var temp item
		if c.vars[cmd[1]].ds == "queue" { c.vars[cmd[1]].value.(*Queue).Back()  // check if the data structure is a queue
		} else { fmt.Println("cli_out> Invalid data structure"); return }
		entry, ok := c.vars[cmd[2]]  // check if the variable exists
		if !ok { fmt.Println("cli_out> Variable not found"); return }
		if entry.ds != c.vars[cmd[1]].value.(ds).Kind() { fmt.Println("cli_out> Data type mismatch"); return }  // check if the data type matches
		entry.value = temp
		c.vars[cmd[2]] = entry
	}


	// Linked List commands ----------------------------------------------------------

	// Data structure commands

	case "addBack": {  // add a value to the back of a linked list or double linked list
		var temp = convert(cmd[2], c.vars[cmd[1]].value.(ds).Kind())  // convert the value to the correct type
		if(temp == nil) { fmt.Println("cli_out> Invalid data type"); return }  // if the value is invalid, return
		switch c.vars[cmd[1]].ds {  // switch to choose the correct data structure
			case "linkedlist": c.vars[cmd[1]].value.(*LinkedList).AddBack(temp)
			case "doublelinkedlist": c.vars[cmd[1]].value.(*DoubleLinkedList).AddBack(temp)
			default: fmt.Println("cli_out> Invalid data structure")
		}
	}

	case "addFront": {  // add a value to the front of a linked list or double linked list
		var temp = convert(cmd[2], c.vars[cmd[1]].value.(ds).Kind())  // convert the value to the correct type
		if(temp == nil) { fmt.Println("cli_out> Invalid data type"); return }  // if the value is invalid, return
		switch c.vars[cmd[1]].ds {  // switch to choose the correct data structure
			case "linkedlist": c.vars[cmd[1]].value.(*LinkedList).AddFront(temp)
			case "doublelinkedlist": c.vars[cmd[1]].value.(*DoubleLinkedList).AddFront(temp)
			default: fmt.Println("cli_out> Invalid data structure")
		}
	}

	case "deleteBack": {  // delete the back value of a linked list or double linked list
		switch c.vars[cmd[1]].ds {  // switch to choose the correct data structure
			case "linkedlist": c.vars[cmd[1]].value.(*LinkedList).DeleteBack()
			case "doublelinkedlist": c.vars[cmd[1]].value.(*DoubleLinkedList).DeleteBack()
			default: fmt.Println("cli_out> Invalid data structure")
		}
	}

	case "deleteFront": {  // delete the front value of a linked list or double linked list
		switch c.vars[cmd[1]].ds {  // switch to choose the correct data structure
			case "linkedlist": c.vars[cmd[1]].value.(*LinkedList).DeleteFront()
			case "doublelinkedlist": c.vars[cmd[1]].value.(*DoubleLinkedList).DeleteFront()
			default: fmt.Println("cli_out> Invalid data structure")
		}
	}

	// Node pointer commands ----------------------------------------------------------

	case "getValue": {  // get the value of a node and store it in a variable
		var temp item
		switch c.vars[cmd[1]].ds {  // switch to choose the correct data structure
			case "linkedlist": temp = c.vars[cmd[1]].value.(*LLNode).GetValue()
			case "doublelinkedlist": temp = c.vars[cmd[1]].value.(*DLLNode).GetValue()
			default: { fmt.Println("cli_out> Invalid data structure"); return }
		}
		if entry, ok := c.vars[cmd[2]]; ok {  // check if the variable exists
			if(entry.ds != c.vars[cmd[1]].value.(ds).Kind()) { fmt.Println("cli_out> Data type mismatch"); return } // check if the data type matches
			entry.value = temp
			c.vars[cmd[2]] = entry
		} else { fmt.Println("cli_out> Invalid variable name") }
	}

	case "ptrFront": {  // get the front node of a linked list or double linked list and store it in a variable
		var temp item
		switch c.vars[cmd[1]].ds {  // switch to choose the correct data structure
			case "linkedlist": temp = c.vars[cmd[1]].value.(*LinkedList).Front()
			case "doublelinkedlist": temp = c.vars[cmd[1]].value.(*DoubleLinkedList).Front()
			default: { fmt.Println("cli_out> Invalid data structure"); return }
		}
		if entry, ok := c.vars[cmd[2]]; ok {  // check if the variable exists
			if(entry.ds != "ptr") { fmt.Println("cli_out> Data type mismatch"); return }  // check if the data type matches
			entry.value = temp
			c.vars[cmd[2]] = entry
		} else { fmt.Println("cli_out> Invalid variable name") }
	}

	case "ptrBack": {  // get the back node of a linked list or double linked list and store it in a variable
		var temp item
		switch c.vars[cmd[1]].ds {  // switch to choose the correct data structure
			case "linkedlist": temp = c.vars[cmd[1]].value.(*LinkedList).Back()
			case "doublelinkedlist": temp = c.vars[cmd[1]].value.(*DoubleLinkedList).Back()
			default: { fmt.Println("cli_out> Invalid data structure"); return }
		}
		if entry, ok := c.vars[cmd[2]]; ok {  // check if the variable exists
			if(entry.ds != "ptr") { fmt.Println("cli_out> Data type mismatch"); return }  // check if the data type matches
			entry.value = temp
			c.vars[cmd[2]] = entry
		} else { fmt.Println("cli_out> Invalid variable name") }
	}

	case "ptrNext": {  // get the next node of a linked list or double linked list and store it in a variable
		if entry, ok := c.vars[cmd[1]]; ok {  // check if the variable exists
			if(entry.value == nil) { fmt.Println("cli_out> Accessing nil pointer"); return }  // check if the pointer is nil
			entry.value = c.vars[cmd[1]].value.(*DLLNode).NextNode()  // get the next node
			c.vars[cmd[1]] = entry
		} else { 
			fmt.Println("cli_out> Invalid variable") 
		}
	}

	case "ptrPrev": {  // get the previous node of a linked list or double linked list and store it in a variable
		if entry, ok := c.vars[cmd[1]]; ok {  // check if the variable exists
			if(entry.value == nil) { fmt.Println("cli_out> Accessing nil pointer"); return }  // check if the pointer is nil
			entry.value = c.vars[cmd[1]].value.(*DLLNode).PrevNode()  // get the previous node
			c.vars[cmd[1]] = entry
		} else { 
			fmt.Println("cli_out> Invalid variable") 
		}
	}

	case "isHead": {  // check if a node is the head of a linked list or double linked list
		if entry, ok := c.vars[cmd[2]]; ok {  // check if the variable exists
			if( entry.ds != "bool" ){ fmt.Println("cli_out> Invalid variable type"); return }  // check if the data type matches
			if( c.vars[cmd[1]].value == nil ) { fmt.Println("cli_out> Accessing nil pointer"); return }  // check if the pointer is nil
			entry.value = c.vars[cmd[1]].value.(*DLLNode).IsHead()  // check if the node is the head
			c.vars[cmd[2]] = entry
		} else { 
			fmt.Println("cli_out> Invalid variable") 
		}		
	}

	case "isEnd": {  // check if a node is the end of a linked list or double linked list
		if entry, ok := c.vars[cmd[2]]; ok {  // check if the variable exists
			if( entry.ds != "bool" ){ fmt.Println("cli_out> Invalid variable type"); return }  // check if the data type matches
			if( c.vars[cmd[1]].value == nil ) { fmt.Println("cli_out> Accessing nil pointer"); return }  // check if the pointer is nil
			entry.value = c.vars[cmd[1]].value.(*DLLNode).IsEnd()  // check if the node is the end
			c.vars[cmd[2]] = entry
		} else {
			fmt.Println("cli_out> Invalid variable")
		}
	}



	// Binary Search Tree commands ---------------------------------------------------

	case "insert": {  // insert a number into a binary search tree
		num, err := strconv.Atoi(cmd[2])  // convert the string to an integer
		if err != nil { fmt.Println("cli_out> Invalid number"); return }  // check if conversion failed
		if c.vars[cmd[1]].ds == "bst" { c.vars[cmd[1]].value.(*BST).Insert(num) // insert the number
		} else { fmt.Println("cli_out> Invalid data structure") }
	}

	case "delete": {  // delete a number from a binary search tree
		num, err := strconv.Atoi(cmd[2])  // convert the string to an integer
		if err != nil { fmt.Println("cli_out> Invalid number"); return }  // check if conversion failed
		if c.vars[cmd[1]].ds == "bst" { c.vars[cmd[1]].value.(*BST).Delete(num)  // delete the number
		} else { fmt.Println("cli_out> Invalid data structure") }
	}

	case "root": {  // get the root of a binary search tree
		if c.vars[cmd[1]].ds == "bst" { fmt.Println("cli_out>", c.vars[cmd[1]].value.(*BST).Root() )  // print the root if ds is bst
		} else { fmt.Println("cli_out> Invalid data structure") }
	}

	case "inorder": { // print the binary search tree in inorder
		if c.vars[cmd[1]].ds == "bst" { c.vars[cmd[1]].value.(*BST).Inorder() // print the tree if ds is bst
		} else { fmt.Println("cli_out> Invalid data structure") }
	}

	case "preorder": { // print the binary search tree in preorder
		if c.vars[cmd[1]].ds == "bst" { c.vars[cmd[1]].value.(*BST).Preorder() // print the tree if ds is bst
		} else { fmt.Println("cli_out> Invalid data structure") }
	}

	case "postorder": { // print the binary search tree in postorder
		if c.vars[cmd[1]].ds == "bst" { c.vars[cmd[1]].value.(*BST).Postorder() // print the tree if ds is bst
		} else { fmt.Println("cli_out> Invalid data structure") }
	}


	// Arithmetic commands -----------------------------------------------------------

	case "+": { // add two numbers and store the result in a variable
		if val, ok := c.vars[cmd[2]]; ok { // check if the variable exists
			if val.ds == c.vars[cmd[1]].ds { // check if the data structures matches
				entry := c.vars[cmd[1]]
				switch val.ds { // add the numbers based on the data structure
					case "int": entry.value = entry.value.(int) + val.value.(int)
					case "float": entry.value = entry.value.(float64) + val.value.(float64)
					case "string": entry.value = entry.value.(string) + val.value.(string)
					default: { fmt.Println("cli_out> Invalid data structure"); return }
				}
				c.vars[cmd[1]] = entry
			} else { fmt.Println("cli_out> Mismatch data structure") }
		} else {  // if the variable does not exist, assume it is a number or item
			entry := c.vars[cmd[1]]
			temp := convert(cmd[2], entry.ds) // convert the string to the data structure
			if temp == nil { fmt.Println("cli_out> Invalid data type"); return } // check if conversion failed
			entry.value = entry.value.(int) + temp.(int)
			c.vars[cmd[1]] = entry
		}
	}

	case "-": { // subtract two numbers and store the result in a variable
		if val, ok := c.vars[cmd[2]]; ok { // check if the variable exists
			if val.ds == c.vars[cmd[1]].ds { // check if the data structures matches
				entry := c.vars[cmd[1]]
				switch val.ds { // subtract the numbers based on the data structure
					case "int": entry.value = entry.value.(int) - val.value.(int)
					case "float": entry.value = entry.value.(float64) - val.value.(float64)
					default: { fmt.Println("cli_out> Invalid data structure"); return }
				}
				c.vars[cmd[1]] = entry
			} else { fmt.Println("cli_out> Mismatch data structure") }
		} else { // if the variable does not exist, assume it is a number or item
			entry := c.vars[cmd[1]]
			switch entry.ds { // subtract the numbers based on the data structure
				case "int": {
					num, err := strconv.Atoi(cmd[2]) // convert the string to an integer
					if err != nil { fmt.Println("cli_out> Invalid number") } // check if conversion failed
					entry.value = entry.value.(int) - num
				}
				case "float": {
					num, err := strconv.ParseFloat(cmd[2], 64) // convert the string to a float
					if err != nil { fmt.Println("cli_out> Invalid number") } // check if conversion failed
					entry.value = entry.value.(float64) - num
				}
				default: fmt.Println("cli_out> Invalid data structure")
			}
			c.vars[cmd[1]] = entry
		}
	}

	case "*": { // multiply two numbers and store the result in a variable
		if val, ok := c.vars[cmd[2]]; ok { // check if the variable exists
			if val.ds == c.vars[cmd[1]].ds { // check if the data structures matches
				entry := c.vars[cmd[1]]
				switch val.ds { // multiply the numbers based on the data structure
					case "int": entry.value = entry.value.(int) * val.value.(int)
					case "float": entry.value = entry.value.(float64) * val.value.(float64)
					default: { fmt.Println("cli_out> Invalid data structure"); return }
				}
				c.vars[cmd[1]] = entry
			} else { fmt.Println("cli_out> Mismatch data structure") }
		} else { // if the variable does not exist, assume it is a number or item
			entry := c.vars[cmd[1]]
			switch entry.ds { // multiply the numbers based on the data structure
				case "int": {
					num, err := strconv.Atoi(cmd[2]) // convert the string to an integer
					if err != nil { fmt.Println("cli_out> Invalid number") } // check if conversion failed
					entry.value = entry.value.(int) * num
				}
				case "float": {
					num, err := strconv.ParseFloat(cmd[2], 64) // convert the string to a float
					if err != nil { fmt.Println("cli_out> Invalid number") } // check if conversion failed
					entry.value = entry.value.(float64) * num
				}
				default: fmt.Println("cli_out> Invalid data structure")
			}
			c.vars[cmd[1]] = entry
		}
	}

	case "/": { // divide two numbers and store the result in a variable
		if val, ok := c.vars[cmd[2]]; ok { // check if the variable exists
			if val.ds == c.vars[cmd[1]].ds { // check if the data structures matches
				entry := c.vars[cmd[1]]
				switch val.ds { // divide the numbers based on the data structure
					case "int": entry.value = entry.value.(int) / val.value.(int)
					case "float": entry.value = entry.value.(float64) / val.value.(float64)
					default: { fmt.Println("cli_out> Invalid data structure"); return }
				}
				c.vars[cmd[1]] = entry
			} else { fmt.Println("cli_out> Mismatch data structure") }
		} else { // if the variable does not exist, assume it is a number or item
			entry := c.vars[cmd[1]]
			switch entry.ds { // divide the numbers based on the data structure
				case "int": {
					num, err := strconv.Atoi(cmd[2]) // convert the string to an integer
					if err != nil { fmt.Println("cli_out> Invalid number") } // check if conversion failed
					entry.value = entry.value.(int) / num
				}
				case "float": {
					num, err := strconv.ParseFloat(cmd[2], 64) // convert the string to a float
					if err != nil { fmt.Println("cli_out> Invalid number") } // check if conversion failed
					entry.value = entry.value.(float64) / num
				}
				default: fmt.Println("cli_out> Invalid data structure")
			}
			c.vars[cmd[1]] = entry
		}
	}

	case "%": { // divide two numbers and store the modulus result in a variable
		if val, ok := c.vars[cmd[2]]; ok { // check if the variable exists
			if val.ds == c.vars[cmd[1]].ds { // check if the data structures matches
				entry := c.vars[cmd[1]]
				switch val.ds {
					case "int": entry.value = entry.value.(int) % val.value.(int)
					default: { fmt.Println("cli_out> Invalid data structure"); return }
				}
				c.vars[cmd[1]] = entry
			} else { fmt.Println("cli_out> Mismatch data structure") }
		} else { // if the variable does not exist, assume it is a number or item
			entry := c.vars[cmd[1]]
			switch entry.ds { // divide the numbers based on the data structure
				case "int": {
					num, err := strconv.Atoi(cmd[2]) // convert the string to an integer
					if err != nil { fmt.Println("cli_out> Invalid number") } // check if conversion failed
					entry.value = entry.value.(int) % num
				}
				default: fmt.Println("cli_out> Invalid data structure")
			}
			c.vars[cmd[1]] = entry
		}
	}

	case "=": { // assign a value to a variable
		if val, ok := c.vars[cmd[2]]; ok { // check if the variable exists
			if val.ds == c.vars[cmd[1]].ds { // check if the data structures matches
				entry := c.vars[cmd[1]]
				if ( val.ds == "int" || val.ds == "float" || val.ds == "string" || val.ds == "bool" ) { // check if the data structure is a primitive
					entry.value= val.value
					c.vars[cmd[1]] = entry
				} else { fmt.Println("cli_out> Invalid data structure") }
			} else { fmt.Println("cli_out> Mismatch data structure") }
		} else { // if the variable does not exist, assume it is a number or item
			entry := c.vars[cmd[1]]
			temp := convert(cmd[2], entry.ds) // convert the string to the data structure
			if(temp == nil) { fmt.Println("cli_out> Invalid data structure"); return } // check if conversion failed
			entry.value = temp
			c.vars[cmd[1]] = entry
		}
	}




case "if": { // if statement, if the variable is true, print the remaining words
		if val, ok := c.vars[cmd[1]]; ok { // check if the variable exists
			if val.ds == "bool" { // check if the data structure is a boolean
				if(val.value.(bool)) { 
					fmt.Print("cli_out> ")
					for i := 2; i < len(cmd); i++ { fmt.Print(cmd[i] + " ") }  // print all the remaining words
					fmt.Print("\n")
				}
			} else { fmt.Println("cli_out> Invalid data structure") }
		} else {
			fmt.Println("cli_out> Invalid variable")
		}
}



	// Default -----------------------------------------------------------------------

	default: fmt.Println("cli_out> Invalid command")
	}
}


// ======================================================================================================

