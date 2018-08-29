package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		printCountUp("A")
	}()
	go func() {
		defer wg.Done()
		printCountUp("B")
	}()

	wg.Wait()
}

func printCountUp(prefix string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%s-%d ", prefix, i)
		time.Sleep(100 * time.Millisecond)
	}

}
