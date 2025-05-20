package slices

import "fmt"

func subslices() {
  //modifying sublices will have affect on the original slices
  //because slices are views on the same underlying array
  
  org:=[]int{1,2,3,4,5,6,7,8}
  fmt.Printf("Original Slice: %v",org)
  fmt.Println()
  subslice:=org[2:5]
  fmt.Printf("Subslice before modification: %v", subslice)
  fmt.Println()
  subslice[0]=33
  subslice[1]=44
  fmt.Printf("Subslice after modification: %v", subslice)
  fmt.Println()
  fmt.Printf("Original slice after the subslice is modified: %v",org)
  fmt.Println()
}