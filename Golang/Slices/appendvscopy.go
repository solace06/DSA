package slices

import "fmt"

func appendVsCopy() {
  // append vs copy when capacity of slice is exceeded
  
  //append behaviour
  org:=make([]int, 3, 5)
  org[0],org[1],org[2]=1,2,3
  
  //within the capacity
  ext:=append(org,4)
  fmt.Printf("After append within capacity: %v, length: %d, capacity %d", ext, len(ext), cap(ext))
  fmt.Println()
  
  //more than capacity
  ext=append(ext,5,6)
  fmt.Printf("After append within capacity: %v, length: %d, capacity %d", ext, len(ext), cap(ext))
  fmt.Println()
  
  //affect on original slice
  fmt.Printf("Affect on the original slice:  %v, length: %d, capacity %d", org, len(org), cap(org))
  fmt.Println()
  
  //copy behaviour
  src:=[]int{15,16,17}
  //cap smaller than the src
  dest:=make([]int, 2)
  count:=copy(dest,src)
  fmt.Printf("Source: %v", src)
  fmt.Println()
  fmt.Printf("Destination: %v No of elements copied: %d", dest, count)
  fmt.Println()
  
  //larger than the src
  dest=make([]int, 5)
  count=copy(dest,src)
  fmt.Printf("Destination: %v No of elements copied: %d", dest, count)
  fmt.Println()
}