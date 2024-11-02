package main

import "reflect"

func pivot(arr []int, low int, high int) int {
	pivotElement := arr[low]
	idx := low
	jdx := high
	swap := reflect.Swapper(arr)
	for idx < jdx {
		for idx < high && arr[idx] <= pivotElement {
			idx++
		}
		for jdx > low && arr[jdx] > pivotElement {
			jdx--
		}
		if idx < jdx {
			swap(idx,jdx)
		}
	}
	swap(low,jdx)
	return jdx
}

func quickSort(arr []int, low int, high int) {
	if low < high {
		pivotIndex := pivot(arr, low, high)
		quickSort(arr, low, pivotIndex-1)
		quickSort(arr, pivotIndex+1, high)
	}
}