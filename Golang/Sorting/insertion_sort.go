package main

func insertionSort(arr []int, arrSize int){
	var currentValue int
	var index int
	for i:=1; i<arrSize; i++{
		currentValue=arr[i]
		index=i-1
		for index>=0 && currentValue<arr[index]{
			arr[index+1]=arr[index]
			index=index-1
		}
		arr[index+1]=currentValue
	}
}
