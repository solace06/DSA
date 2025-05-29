package slices

import (
	"fmt"
	"sort"
)

func secondlargest() {
	var n int
	fmt.Println("Enter the number of elements")
	fmt.Scan(&n)

	fmt.Println("Enter the elements")
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&slice[i])
	}

	//brute force approach
	sort.Ints(slice)
	largest := slice[n-1]
	second_largest := -1
	for i := n - 2; i >= 0; i-- {
		if slice[i] != largest {
			second_largest = slice[i]
			break
		}
	}

	//better approach
	largest = slice[0]
	for i := 0; i < n; i++ {
		if slice[i] > largest {
			largest = slice[i]
		}
	}

	second_largest = -1
	for i := 0; i < n; i++ {
		if slice[i] > second_largest && largest > slice[i] {
			second_largest = slice[i]
		}
	}

	fmt.Println("Second largest element is:", second_largest)

}
