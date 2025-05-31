package slices

import "fmt"

func twoSum() {
  //two sum
  
  fmt.Println("Enter the target sum")
  var target int
  fmt.Scan(&target)
  
  slice:=[]int{3,5,2,11,7,8,3,9}
  
  m:=make(map[int]bool)
  
  flag:=0
  for _, val:= range slice{
      complement:=target-val
      if m[complement]{
          fmt.Println("The values that have the target sum are as follows:",val,"and",complement)
          flag=1
          break
      }
      m[val]=true
  }


  if flag == 0{
      fmt.Println("No two numbers add up to the target")
  }
}