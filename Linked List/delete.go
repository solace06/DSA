package main

import (
	"errors"
	"fmt"
)

func deleteFront(head *Node) *Node {
	if head == nil {
		fmt.Println("Linked List is empty")
		return nil
	}

	fmt.Println("Deleted Node: ", head.data)
	head = head.next
	return head
}

func deleteBack(head *Node) (*Node, error) {
	if head == nil {
		return nil, errors.New("the linked list is empty")
	}

	if head.next == nil {
		return nil, nil
	}

	curr := head
	for curr.next.next != nil {
		curr = curr.next
	}
	curr.next = nil
	fmt.Println("The last node has been deleted successfully")
	return head, nil
}

func deleteAfter(head *Node, value int) error {
	curr := head

	for curr.data != value && curr != nil {
		curr = curr.next
	}

	if curr == nil {
		return fmt.Errorf("node with value %d does not exist", value)
	}

	if curr.next == nil {
		return fmt.Errorf("node with value %d is the last node", value)
	}

	var temp *Node = curr.next
	fmt.Printf("node after value %d and with data %d is successfully deleted", value, temp.data)

	curr.next = curr.next.next
	return nil
}

func deleteBefore(head *Node, value int) error {

	return nil
}
