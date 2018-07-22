package main

import "fmt"

func main() {
	defer func() {
		switch p := recover(); p {
		case nil: // no panic has occurred
		case "It is the end of the world!":
			fmt.Printf("Recovering from our panic! %v", p)
		default:
			fmt.Printf("Recovering from unknown error! %v", p)
		}
	}()
	panic("It is the end of the world!")
}
