package main

// Question: Start two goroutines, one printing "Hello" and the other printing "World", 
// ensuring both execute before the main function exits.

import (
    "fmt"
    "sync"
    )

func GoroutineSync(){
    wg:=&sync.WaitGroup{}
    wg.Add(2)
    go func(){
        defer wg.Done()
        fmt.Println("Hello")
    }()
    
    go func(){
        defer wg.Done()
        fmt.Println("World")
    }()
    wg.Wait()
}