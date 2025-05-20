package slices

import (
    "fmt"

    )



func remove() {
 //remove an element from a slice 
 slice:=[]int{1,2,3,4,5,6,7,8,9}
 fmt.Printf("Slice before removing the element at index 4: %v",slice)
 fmt.Println()
 
 //removing index 4
 slice=append(slice[:4], slice[5:]...)
 fmt.Printf("Slice after removing the element at index 4: %v",slice)
 fmt.Println()
}