package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {

	const helloWorld = "Hello world! :)"

	fmt.Println(helloWorld)

	a := person{
		name: "Some name",
		age:  10,
	}

	a1 := person {}
	fmt.Println(a1)

	a2 := person {"name 2", 2}
	fmt.Println(a2)

	a3 := new(person)
	fmt.Println(a3)

	fmt.Println(a)

	b := struct {
		description string
	}{
		description: "This is an anonymous struct!",
	}

	fmt.Println(b.description)
	fmt.Println((&b).description)

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
