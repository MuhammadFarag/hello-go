package main

func main() {
	var s Speaker
	s = Person{"Muhammad"}
	s.speak()
}

type Speaker interface {
	speak()
}

type Person struct {
	name string
}

func (p Person) speak() {
	println("My name is", p.name)
}
