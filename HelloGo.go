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

	fmt.Println(a)

	b := struct {
		description string
	}{
		description: "This is an anonymous struct!",
	}

	fmt.Println(b.description)

	c := 5
	fmt.Println("the address of c is:", &c)

}
