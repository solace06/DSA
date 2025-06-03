package main

import (
	"fmt"
	"sync"
)

func deadlock() {
	var mu1, mu2 sync.Mutex

	go func() {
		mu1.Lock()
		fmt.Println("Goroutine 1 acquired mu1")

		mu2.Lock()
		fmt.Println("Goroutine 1 acquired mu2")

		mu2.Unlock()
		mu1.Unlock()
	}()

	go func() {
		mu2.Lock()
		fmt.Println("Goroutine 2 acquired mu2")

		mu1.Lock()
		fmt.Println("Goroutine 2 acquired mu1")

		mu1.Unlock()
		mu2.Unlock()
	}()

	select {} 
}
