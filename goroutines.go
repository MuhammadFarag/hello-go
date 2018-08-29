package main

import (
	"fmt"
	"time"
)

func main() {
	go printCountUp("A")
	printCountUp("B")
}

func printCountUp(prefix string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%s-%d ", prefix, i)
		time.Sleep(100 * time.Millisecond)
	}
}
