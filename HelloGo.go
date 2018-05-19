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

}
