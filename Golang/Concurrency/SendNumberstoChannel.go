package main

import (
	"fmt"
	"sync"
)

//Implement a function that launches multiple goroutines, each sending a number 
//(1 to N) to a channel, and collects all values in the main function.

func SendNumberstoChannel() {
	ch := make(chan int)
	wg := sync.WaitGroup{}

	go func() {
		for i := 1; i <= 100; i++ {
			wg.Add(1)
			go func(num int) {
				defer wg.Done()
				ch <- num
			}(i)
		}
		wg.Wait()
		close(ch)
	}()

	for i := range ch {
		fmt.Println(i)
	}
}
