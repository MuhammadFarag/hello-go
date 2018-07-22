package main

import "fmt"

func main() {
	a := account{}
	fmt.Println("Account before:", a)
	a = a.add(15)
	fmt.Println("Account after add:", a)

	account.add(a, 1)
}

type account struct {
	money int
}

func (a account) add(amount int) account {
	a.money += amount
	return a
}
