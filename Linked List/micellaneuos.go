package main

func calLength(head *Node) int{

	if head==nil{
		return 0
	}
	
	return 1+calLength(head.next)
}