package main

func merge(left, right []int) []int{
	size:=len(left)+len(right)
	result:=make([]int,size)
	i:=0
	j:=0
	k:=0

	for i<len(left) && j<len(right){
		if left[i]<=right[j]{
			result[k]=left[i]
			i++
		} else{
			result[k]=right[j]
			j++
		}
		k++
	}
	for i<len(left){
		result[k]=left[i]
		i++
		k++
	}
	for j<len(right){
		result[k]=right[j]
		j++
		k++
	}
	return result
}

func mergeSort(arr []int) []int{
	if len(arr)<=1{
		return arr
	}
	mid:=len(arr)/2
	left:=mergeSort(arr[:mid])
	right:=mergeSort(arr[mid:])
	return merge(left,right)
}