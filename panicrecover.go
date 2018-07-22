package main

import "fmt"

func main() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("Recovering from our panic! %v", p)
		}
	}()
	panic("It is the end of the world!")
}
