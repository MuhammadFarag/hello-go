package main

import (
	"bufio"
	"os"
	"io/ioutil"
	"strings"
)

func main() {
	fileName := os.Args[1]
	scanFile(fileName)
	readFile(fileName)
}

func scanFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		println("Failed to open file: ", fileName, ". Error: ", err)
	} else {
		scanner := bufio.NewScanner(file)
		for i := 0; scanner.Scan(); i++ {
			println(i, ": ", scanner.Text())
		}
		file.Close()
	}
}

func readFile(fileName string) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		println("Failed to open file: ", fileName, ". Error: ", err)
	}
	for i, line := range strings.Split(string(file), "\n") {
		println(i, ": ", line)
	}
}
