package main

import "fmt"

func main() {
	fmt.Println("defer1 result", defer1(100))
	fmt.Println("defer2 result", defer2(100))
	fmt.Println("defer3 result", defer3(100))
}

func defer1(i int) int {
	x := i
	defer func() {
		x += 1
	}()

	return x
}

func defer2(i int) (x int) {
	x = i
	defer func() {
		x += 1
	}()

	return x
}

func defer3(i int) (x int) {
	x = i
	defer func() {
		x += 1
	}()

	return 2
}
