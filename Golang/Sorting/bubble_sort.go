package main

import "reflect"

func bubbleSort(arr []int, size int) {

	var swapped int
	swap := reflect.Swapper(arr)
	for i := 0; i < size-1; i++ {
		swapped = 0
		for j := 0; j < size-i-1; j++ {
			if arr[j] > arr[j+1] {
				swap(j,j+1)
				swapped=1
			}
		}
		if swapped==0{
			break;
		}
	}
}
