package main

func main() {
	a := 1
	a = RichInt(a).Increment().toInt()
	println(a)
}

type RichInt int

func (n RichInt) Increment() RichInt {
	return n + 1
}

func (n RichInt) toInt() int {
	return int(n)
}
