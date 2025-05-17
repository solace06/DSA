package main

import "fmt"

//Struct of type Node
type Node struct {
	data int
	next *Node
}

//Function to print the linked list
func print(head *Node) {

	var currNode *Node = head

	for currNode != nil {
		fmt.Print(currNode.data, " ")
		currNode = currNode.next
	}
	fmt.Println()
}

func main() {
	//Creating a linked list

	var head = new(Node)
	var second = new(Node)
	var third = new(Node)

	head.data = 1
	head.next = second

	second.data = 2
	second.next = third

	third.data = 3
	third.next = nil

	fmt.Println("Original Linked List")
	print(head)

	var choice int

	fmt.Println("--------Function Menu--------")
	fmt.Println("1 to print the linked list")
	fmt.Println("2 for Inserting a node at the front of a linked list")
	fmt.Println("3 for Inserting a node at the end of a linked list")
	fmt.Println("4 for Inserting a node after a node value in a linked list")
	fmt.Println("5 for Inserting a node before a node value in a linked list")
	fmt.Println("6 for Deleting the first node of a linked list")
	fmt.Println("7 for Deleting the last node of a linked list")
	fmt.Println("8 for Deleting the node of a linked list after a value")
	fmt.Println("10 to EXIT")

loop:
	for {
		fmt.Println("Enter your choice")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			print(head)
		case 2:
			var value int
			fmt.Println("Enter value for the node")
			fmt.Scan(&value)
			newNode := new(Node)
			newNode.data = value
			newNode.next = nil
			head = insertAtFront(head, newNode)
			fmt.Println("After Insertion")
			print(head)
		case 3:
			var value int
			fmt.Println("Enter value for the node")
			fmt.Scan(&value)
			newNode := new(Node)
			newNode.data = value
			newNode.next = nil
			insertAtBack(head, newNode)
			fmt.Println("After Insertion")
			print(head)
		case 4:
			var value int
			var nodeValue int
			fmt.Println("Enter value for the node")
			fmt.Scan(&value)
			fmt.Println("Enter the value after which node should be inserted")
			fmt.Scan(&nodeValue)
			newNode := new(Node)
			newNode.data = value
			newNode.next = nil
			insertAfter(head, newNode, nodeValue)
			fmt.Println("After Insertion")
			print(head)
		case 5:
			var value int
			var nodeValue int
			fmt.Println("Enter value for the node")
			fmt.Scan(&value)
			fmt.Println("Enter the value before which node should be inserted")
			fmt.Scan(&nodeValue)
			newNode := new(Node)
			newNode.data = value
			newNode.next = nil
			head = insertBefore(head, newNode, nodeValue)
			fmt.Println("After Insertion")
			print(head)
		case 6:
			head = deleteFront(head)
			if head == nil {
				fmt.Println("The Linked List is empty")
			} else {
				fmt.Println("Updated Linked List")
				print(head)
			}
		case 7:
			head, err := deleteBack(head)
			if err == nil {
				fmt.Println("Updated Linked List")
				print(head)
			} else {
				fmt.Println(err)
			}
		case 8:
			fmt.Println("Enter the value")
			var value int
			fmt.Scan(&value)
			err := deleteAfter(head, value)
			if err != nil {
				fmt.Printf("error deleting the node: %s", err.Error())
				fmt.Println()
			} else {
				fmt.Println("Updated Linked List")
				print(head)
			}
		case 10:
			fmt.Println("You chose to exit")
			break loop
		default:
			fmt.Println("Please enter a valid choice")
		}
	}
}
