package main

import "fmt"

func printArray(arr []int, size int){
	for i:=0; i<size; i++{
		fmt.Print(arr[i]," ")
	}
	fmt.Print("\n")
}

func main() {

	var size int
	fmt.Println("Enter the size of the array")
	fmt.Scan(&size)

	arr:=make([]int,size)

	fmt.Println("Enter ",size," elements")
	for i:=0; i<size; i++{
		fmt.Scan(&arr[i])
	}

	fmt.Println("Original Array")
	printArray(arr,size)

	var sort string
	fmt.Println("Select the type of sort you want to execute")
	fmt.Println("1.Select A for Bubble Sort")
	fmt.Println("2.Select B for Selection Sort")
	fmt.Println("3.Select C for Insertion Sort")
	fmt.Println("4.Select D for Quick Sort")
	fmt.Println("5.Select E for Merge Sort")
	fmt.Scan(&sort)
	switch sort{
	case "A":
		fmt.Println("Starting Bubble Sort")
		bubbleSort(arr, size)
		fmt.Println("Sorted Array")
		printArray(arr,size)
	case "B":
		fmt.Println("Starting Selection Sort")
		selectionSort(arr, size)
		fmt.Println("Sorted Array")
		printArray(arr,size)
	case "C":
		fmt.Println("Starting Insertion Sort")
		insertionSort(arr, size)
		fmt.Println("Sorted Array")
		printArray(arr, size)
	case "D":
		fmt.Println("Starting Quick Sort")
		quickSort(arr, 0, size-1)
		fmt.Println("Sorted Array")
		printArray(arr, size)
	case "E":
		fmt.Println("Starting Merge Sort")
		sortedArray:=mergeSort(arr)
		fmt.Println("Sorted Array")
		printArray(sortedArray, size)
	}

}
