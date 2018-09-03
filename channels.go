package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	countChannel := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		printOutput(countChannel)
	}()

	go func() {
		defer close(countChannel)
		defer wg.Done()
		countUp(countChannel)
	}()

	wg.Wait()
}

func countUp(counts chan int) {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		counts <- i
	}
}

func printOutput(counts chan int) {
	for i := range counts {
		fmt.Printf("%d ", i)
	}
}
