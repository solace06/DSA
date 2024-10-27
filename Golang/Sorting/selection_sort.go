package main

import "reflect"

func selectionSort(arr []int, size int) {
	swap := reflect.Swapper(arr)
	for i := 0; i < size-1; i++ {
		minNum := arr[i]
		minIndex:=i
		for j := i + 1; j < size; j++ {
			if minNum > arr[j] {
				minNum = arr[j]
				minIndex=j
			}
		}
		swap(i,minIndex)
	}
}