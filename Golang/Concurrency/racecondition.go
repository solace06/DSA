package main

import (
	"fmt"
	"time"
)

//executes successfully but gives incorrect results

func race() {
    var counter int
	for i := 0; i < 1000; i++ {
		go func() {
			counter++
		}()
	}

	time.Sleep(1 * time.Second)
	fmt.Println("The value for counter is :", counter)
}
