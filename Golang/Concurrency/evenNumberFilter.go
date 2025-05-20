package main

import (
    "fmt"
    "sync"
    )

func generate(nums chan int, n int, wg *sync.WaitGroup){
    defer wg.Done()
    for i:=1; i<=n; i++{
        nums<-i
    }
    close(nums)
}

func filter(nums <-chan int, evens chan<- int, wg *sync.WaitGroup){
    defer wg.Done()
    for num :=range nums{
        if num%2 == 0{
        evens<-num
        }
    }
    close(evens)
}

func filterEven() {
  //Create one goroutine which generates n numbers
  //Create one goroutine which filters even numbers from n numbers
  //Main function prints the even numbers
  
  n:=20
  var wg sync.WaitGroup
  
  wg.Add(2)
  
  nums:=make(chan int)
  evens:=make(chan int)
  
  go generate(nums, n, &wg)
  go filter(nums, evens, &wg)
  
  for val := range evens{
      fmt.Println(val)
  }
  
  wg.Wait()
}