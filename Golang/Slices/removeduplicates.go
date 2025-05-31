package slices

import "fmt"

func removeduplicates() {
	//removing duplicates from a sorted array
	//without using extra space

	slice := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5}

	i := 0
	j := 0
	count := 1

	for j < len(slice) {
		if slice[i] != slice[j] {
			i++
			slice[i], slice[j] = slice[j], slice[i]
			j++
			count++
		} else {
			j++
		}
	}

	fmt.Println("Unique elements in the slice are:")
	for i := 0; i < count; i++ {
		fmt.Print(slice[i])
		fmt.Print(" ")
	}

}
