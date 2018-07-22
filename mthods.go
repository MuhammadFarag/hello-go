package main

import "fmt"

func main() {
	a := account{}
	fmt.Println("Account before:", a)
	a.add(15)
	fmt.Println("Account after add:", a)
}

type account struct {
	money int
}

func (a *account) add(amount int) {
	a.money += amount
}
