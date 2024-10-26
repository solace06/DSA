package main

import "fmt"

func binarySearch(arr []int, size int) {
	
	var element int
	fmt.Println("Enter the element to be searched for")
	fmt.Scan(&element)

	low:=0
	high:=size-1

	for low<=high{
		mid:=(low+high)/2
		if arr[mid]==element{
			fmt.Printf("Found the element at %d",mid)
			return
		} else if arr[mid]>element{
			high=mid-1
		} else{
			low=mid+1
		}
	}

	fmt.Println("The element is not in the list")
	return
}
