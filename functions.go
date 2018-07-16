package main

import (
	"fmt"
)

func main() {
	a, b, c := threeTimes("hello", "Go")
	println(a, b, c)
	println(namedResult())
}

func threeTimes(a, b string) (r1 string, r2 string, r3 string) {
	r := fmt.Sprintf("%s %s", a, b)
	return r, r, r
}

func namedResult() (r string){
	r = "Hello Go"
	return
}