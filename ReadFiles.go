package main

import (
	"bufio"
	"os"
)

func main() {
	fileName := os.Args[1]
	scanFile(fileName)
}

func scanFile(fileName string){
	file, err := os.Open(fileName)
	if err != nil {
		println("Failed to open file: ", fileName, ". Error: ", err)
	} else {
		scanner := bufio.NewScanner(file)
		for i := 0; scanner.Scan(); i++ {
			println(i, ": ", scanner.Text())
		}
	}
}