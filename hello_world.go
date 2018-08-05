package main

import "fmt"

func main() {
	fmt.Println("hello world")

	const PI = 3.14

	var a int
	a = 10
	_ = a
	var b int = 20
	_ = b
	var c = 30
	_ = c
	d := 40
	_ = d

	f := add
	f(10, 20)

	// if
	i := 0
	if i == 0 {
	}

	// switch
	switch i {
	case 1:
		//fmt.Println(1)
	case 2:
		//fmt.Println(2)
	default:
		//fmt.Println("default")
	}

	// for
	for i := 0; i < 10; i++ {
	}

	// while
	for i < 100 {
	}

	// array
	var array []int
	_ = array
	var array1 [10]int
	_ = array1
	var array2 = []int{0, 1, 2, 3, 4}
	_ = array2
	var array3 = [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}
	_ = array3

	for i := range array {
		_ = i
	}

	for index, element := range array {
		_ = index
		_ = element
	}

	for _, element := range array {
		_ = element
	}

	// map
	map0 := map[string]int{}
	_ = map0
	map1 := make(map[string]int)
	_ = map1

	for key, value := range map0 {
		_ = key
		_ = value
	}
}

func add(a int, b int) {
	c := a + b
	_ = c
}

func add1(a int, b int) int {
	c := a + b
	return c
}
