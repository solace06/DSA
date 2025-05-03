package main

import "fmt"

func insertAtBack(head *Node, newNode *Node) {

	var currNode *Node = head
	for currNode.next != nil {
		currNode = currNode.next
	}

	currNode.next = newNode
}

func insertAtFront(head *Node, newNode *Node) *Node {
	newNode.next = head
	head = newNode
	return head
}

func insertAfter(head *Node, newNode *Node, value int) {

	var currNode *Node = head
	for currNode != nil {
		if currNode.data == value {
			var temp *Node = currNode.next
			currNode.next = newNode
			newNode.next = temp
			return
		}
		currNode = currNode.next
	}
	fmt.Println("The value does not exist")
}

func insertBefore(head *Node, newNode *Node, value int) *Node {
	var prev *Node = nil

	if head.data == value {
		newNode.next = head
		head = newNode
		return head
	}

	for curr := head; curr != nil; curr = curr.next {
		if curr.data == value {
			prev.next = newNode
			newNode.next = curr
			return head
		}
		prev = curr
	}
	fmt.Println("The value does not exist")
	return head
}
