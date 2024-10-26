package main

import "fmt"

func main() {

	var size int
	fmt.Println("Enter the size of the slice")
	fmt.Scan(&size)

	arr:=make([]int, size)

	fmt.Println("Enter ", size, " elements")
	for i:=0; i<size; i++{
		fmt.Scan(&arr[i])
	}

	var search string
	
	fmt.Println("Choose the type of search you want to execute")
	fmt.Println("1.Select L for linear search")
	fmt.Println("1.Select B for binary search")
	fmt.Println("1.Select J for jump search")

	fmt.Scan(&search)

	switch search{
	case "L":
		fmt.Println("Starting Linear Search")
		linearSearch(arr, size)
	}
}
