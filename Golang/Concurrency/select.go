package main

import (
    "fmt"
    "time"
    )

func selectChannel() {
    // select with channels
    
    chan1:=make(chan string)
    chan2:=make(chan string)
    
    go func(){
       time.Sleep(2 * time.Second)
       chan1<-"Message sent to channel 1"
    }()
    
    go func(){
        time.Sleep(4 * time.Second)
        chan2<-"Message sent to channel 2"
    }()
    
    for i:=0; i<2; i++{
        select {
            case val:=<-chan1:
                fmt.Println(val)
            case val:=<-chan2:
                fmt.Println(val)
        }
    }
    
}