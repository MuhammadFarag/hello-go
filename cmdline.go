package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {

	for i := 0; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
	}

	f := floatFlag{9}
	flag.CommandLine.Var(&f, "flag-name", "usage of this test flag")
	flag.Parse()
	fmt.Println(f.float64)

}

type floatFlag struct {
	float64
}

func (m *floatFlag) String() string {
	return fmt.Sprintf("%v", m.float64)
}

func (m *floatFlag) Set(s string) error {
	m.float64, _ = strconv.ParseFloat(s, 32)
	return nil
}
