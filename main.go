package main

import (
	"fmt"
	"strconv"
)


func main() {

	fmt.Println("\nHello to ds_cli")

	cli := NewCLI()
	cli.Run()
	// cli.Seed(seed1)
	
}




// ===== INSTRUCTIONS =====================================================================================

// Valid data types are:                int, float, string, bool, ptr
// Valid data structures are:           stack, queue, bst, linkedlist, doublelinkedlist



// ===== GENERAL METHODS ============================================================================================

// var <Data_Type> <Variable_name>   			  				 // Creates a new variable of given data type
// var <DS_Type> <Data_Type> <Variable_name>         // Creates a new variable of given data structure
// print <variable_name>                             // Prints the contents of the variable
// search <Variable_name> <value>                    // Boolean - Checks if the given exists in given variable



// ===== STACK METHODS ==============================================================================================

// push <Variable_name> <value>                      // Push value in the given stack
// pop <Variable_name> <res>      				  				 // Pops the top value of the stack and stores it in res
// top <Variable_name> <res>                         // Stores the top value of the stack in res
// isEmpty <Variable_name> <res>					  				 // checks if the stack is empty and stores result in res



// ===== QUEUE METHODS ==============================================================================================

// push <variable_name> <value>                      // Push value to the queue
// pop <variable_name> <res>					               // Pops the front element of queue and stores it in res
// front <variable_name> <res>			                 // Stores the front element of queue in res
// back <variable_name> <res>					               // Stores the back element of queue in res
// isEmpty <variable_name> <res>					           // Checks if the queue is empty and stores result in res



// ===== BINARY SEARCH TREE METHODS =================================================================================

// insert <variable_name> <value>			               // Inserts given value into the BST
// delete <variable_name> <value>			               // Deletes the given value from BST
// root <variable_name> 					                   // Prints the value at root of BST
// inorder <variable_name> 				                   // Prints in-order traversal of BST
// preorder <variable_name>              		         // Prints pre-order traversal of BST
// postorder <variable_name> 				                 // Prints post-order traversal of BST



// ===== LINKED LIST METHODS ========================================================================================

// addBack <variable_name> <value>			             // Adds the value to the back of list
// delBack <variable_name>					                 // Deletes value from back of list
// addFront <variable_name> <value>			             // Adds the value to the front of list
// delFront <variable_name>				                   // Deletes value from front of list
// addAfter <variable_name> <node_value> <value>	   // Adds value after the given value in given list

// ptrFront <pointer_name> 				                   // Points given pointer to the front of list
// ptrBack <pointer_name>					                   // Points given pointer to the back of list
// getNode <pointer_name>					                   // Gets node value from pointer
// nextNode <pointer_name>					                 // Moves pointer to the next node
// prevNode <pointer_name>					                 // Moves pointer to the previous node
// isEnd <pointer_name>	<res>				                 // Checks if the pointer points to back of the list and stores result in res
// isHead <pointer_name> <res>				               // Checks if the pointer points to front of the list and stores result in res




// ================================================================================================================================


// Seed data for testing
var seed1 = []string{
	"var int a",
	"var int b",
	"= a 10",
	"= b 20",
	"print a",
	"print b",
	"+ a b",
	"print a",
}
// Expected output
// 10
// 20
// 30



// ================================================================================================================================


// Converts value to the specified type
func convert(value string, kind string) item {
	switch kind {
		case "int": {
			val, err := strconv.Atoi(value)
			if err != nil { return nil } else { return val }
		}
		case "float":{
			val, err := strconv.ParseFloat(value, 64)
			if err != nil { return nil } else { return val }
		}
		case "string": {
			return value
		}
		case "bool": {
			val, err := strconv.ParseBool(value)
			if err != nil { return nil } else { return val }
		}
	}
	return nil
}
