package main

import (
	"errors"
	"fmt"
)

func deleteFront(head *Node) *Node{
	if head==nil{
		fmt.Println("Linked List is empty")
		return nil
	}
	
	fmt.Println("Deleted Node: ", head.data)
	head=head.next
	return head
}

func deleteBack(head *Node) (*Node, error){
	if head == nil{
		return nil, errors.New("the linked list is empty")
	}

	if head.next==nil{
		return nil, nil
	}

	curr:=head
	for curr.next.next!=nil{
		curr=curr.next
	}
	curr.next=nil
	fmt.Println("The last node has been deleted successfully")
	return head, nil
}
