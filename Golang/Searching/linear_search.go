package main

import "fmt"

func linearSearch(arr []int, size int) {

	var element int
	fmt.Println("Enter the element to be searched for")
	fmt.Scan(&element)

	for i:=0; i<size; i++{
		if arr[i]==element{
			fmt.Printf("Found it at position %d",i)
			return
		}
	}
	fmt.Println("The element is not present in the list")
	return
}
