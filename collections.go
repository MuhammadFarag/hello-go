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

	a := [...]string{0: "a", 2: "b"}

	for i, s := range a {
		fmt.Println("check: ", i, s)
	}
	// slice doubles in capacity to expand until it reaches 1000 elements and then it increases by ~25%
	var zeroSlice []int
	info("zeroSlice", zeroSlice)

	slice := make([]int, 5)
	slice[4] = 1
	info("slice", slice)
	slice = append(slice, 20)
	info("slice", slice)

	slice = append(slice, 10)
	sliceWithCapacity := make([]int, 5, 8)
	info("sliceWithCapacity", sliceWithCapacity)
	sliceWithCapacity = append(sliceWithCapacity, 10, 20, 30, 40)
	info("sliceWithCapacity", sliceWithCapacity)

	subSlice := slice[4:5]
	info("subSlice", subSlice)

	fmt.Println("Let's change the world:")
	subSlice[0] = -1
	subSlice = append(subSlice, -2)
	info("subSlice", subSlice)
	info("slice", slice)

	// Since those two dudes are sharing the backing arrays, a change in one, will change the other!

	fmt.Println("Let's avoid changing the world")
	println()

	subSlice2 := slice[4:5:5]
	// append changes the backing array to be able to expand the sub slice
	subSlice2 = append(subSlice2, 100)
	info("subSlice2", subSlice2)
	subSlice2[0] = 1982
	info("subSlice2", subSlice2)
	info("slice", slice)
	info("subSlice", subSlice)

	map1 := make(map[string]int)
	map1["Samsung"] = 3
	map1["Samsung"] = 4
	map1["Apple"] = 0
	map1["Microsoft"] = 3
	fmt.Println(map1)
	delete(map1, "Apple")
	fmt.Println(map1)

	v, found := map1["Toshiba"]
	fmt.Println("None existing value: '", v, found, "'")

	map2 := map[string]int{
		"Google":   4,
		"LinkedIn": 5,
	}

	for k, v := range map2 {
		fmt.Println(k, v)
	}

}

func info(name string, s []int) {
	fmt.Println(name, "content", s, "len:", len(s), "cap:", cap(s))
}
