package main

import (
    "fmt"
    "time"
    )

func goroutineSignalling() {
  signal:=make(chan struct{})
  
  go func(){
    fmt.Println("Goroutine waiting for the signal")
    <-signal
    time.Sleep(1*time.Second)
    fmt.Println("Goroutine recieved the signal")
  }()
  
  time.Sleep(2*time.Second)
  fmt.Println("Sending the signal now")
  signal<-struct{}{}
  
  time.Sleep(2*time.Second)
}