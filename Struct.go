package main

import "fmt"

type person struct {
	name string
	age  int
}

type player struct {
	person
	favouriteGame string
	int
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

	p := player{person{"jack", 23}, "game", 1}

	fmt.Println(p)
	fmt.Println(p.person.name)
	fmt.Println(p.name)
	fmt.Println(p.int)
}
