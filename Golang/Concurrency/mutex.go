package main

import (
    "fmt"
    "sync"
    )
    
    
func mutex() {
    var counter = 0
    var wg sync.WaitGroup
    var mu sync.Mutex
    
    for i:=0; i<100; i++{
        wg.Add(1)
        go func(){
            defer wg.Done()
            mu.Lock()
            counter++
            mu.Unlock()
        }()
    }
        
    wg.Wait()
    fmt.Println("The value of the counter is:", counter)
}