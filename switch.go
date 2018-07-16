package main

import "strings"

func main() {
	switch "hi" {
	case "hello", "hi":
		println("hi")
		fallthrough
	case "ola":
		println("ola")
	default:
		println("The default")
	}

	x := "hi"
	switch {
	case strings.Contains("hi there", x):
		println("contains hi")
	default:
		println("default")
	}
}
