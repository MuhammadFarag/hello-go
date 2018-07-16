package main

import "fmt"

func main() {

	const helloWorld = "Hello world! :)"

	fmt.Println(helloWorld)

	c := 5
	fmt.Println("the address of c is:", &c)

	const (
		c0 = iota // c0 == 0
		c1        // c1 == 1
		c2        // c2 == 2
	)

	fmt.Println(c2)

	type timeInSeconds int
	type timeInMilliSeconds int

	var seconds timeInSeconds = 10
	var millis timeInMilliSeconds = timeInMilliSeconds(1000 * int(seconds))

	fmt.Println(millis)

}
