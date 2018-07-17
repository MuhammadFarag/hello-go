package main

import (
	"errors"
	"fmt"
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
