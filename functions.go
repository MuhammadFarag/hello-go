package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	a, b, c := threeTimes("hello", "Go")
	println(a, b, c)
	println(namedResult())

	_, ok := somethingMightGoWrong()
	if !ok {
		// an error has occurred
	}
	println(ok)
	_, err := somethingElseMightGoWrong()
	if err != nil {
		// an error has occurred
	}

	if _, err := somethingElseMightGoWrong(); err != nil {
		// do something
	}
	println(err.Error())

	if err := returnSpecificError(); err == SpecificError {
		println(err.Error())
	}

	consumeBehaviour(namedResult)

	consumeBehaviour(func() string { return "Hi" })

	variadic(1)

	letsDefer("hello")

	captureMutation()
}

func threeTimes(a, b string) (r1 string, r2 string, r3 string) {
	r := fmt.Sprintf("%s %s", a, b)
	return r, r, r
}

func namedResult() (r string) {
	r = "Hello Go"
	return
}

func somethingMightGoWrong() (string, bool) {
	return "", false
}

func somethingElseMightGoWrong() (string, error) {
	return "", fmt.Errorf("error: %s", "error description")
}

var SpecificError = errors.New("Specific error")

func returnSpecificError() error {
	return SpecificError
}

func consumeBehaviour(f func() string) {
	println(f())
}

func variadic(x ...int) {
	println(fmt.Sprintf("The type of this function is: %T", variadic))
	println(fmt.Sprintf("The type of argument is %T and its value is %v", x, x))
}

func letsDefer(s string) (r string) {
	defer func() { println("defer 1:", s, r) }()
	r = "r-" + s
	defer func() { println("defer 2:", s, r) }()
	return r
}

func captureMutation() {
	defer func() func() {
		before := time.Now()
		return func() { println("Elapsed time:", (time.Now().Sub(before)).String()) }
	}()()
}
