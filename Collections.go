package main

import "fmt"

func main() {

	var array [5]string

	array[0] = "a"
	array[1] = "b"
	array[2] = "c"
	array[3] = "d"
	array[3] = "e"

	for i := 0; i < len(array); i++ {
		fmt.Println(i, array[i])
	}

	for i, s := range array {
		fmt.Println(i, s)
	}

	strings2 := [5]string{"a", "b", "c", "d", "e"}

	for i, s := range strings2 {
		fmt.Println(i, s)
	}

	for i := 0; i < len(strings2); i++ {
		fmt.Println(i, strings2[i])
	}

}
