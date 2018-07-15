package main

import "fmt"

type person struct {
	name string
	age  int
}


func main() {
	a := person{
		name: "Some name",
		age:  10,
	}

	a1 := person{}
	fmt.Println(a1)

	a2 := person{"name 2", 2}
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
}
