package main

import "fmt"

func deleteFront(head *Node) *Node{
	fmt.Println("Deleted Node: ", head.data)
	head=head.next
	return head
}
