package main

import (
	"fmt"
	"unsafe"
)

func test() {
	//i := 5
	//iPointer := &i

	//f := (*float32)(iPointer)
	//fmt.Println(f)
}

func unsafe1() {
	i := 5
	iPointer := &i

	f := (*float32)(unsafe.Pointer(iPointer))
	fmt.Println(f)
}

func unsafe2() {
	type Num struct {
		i string
		j int64
	}
	n := Num{i: "fish", j: 1}
	nPointer := unsafe.Pointer(&n)

	niPointer := (*string)(unsafe.Pointer(nPointer))
	*niPointer = "煎鱼"

	njPointer := (*int64)(unsafe.Pointer(uintptr(nPointer) + unsafe.Offsetof(n.j)))
	*njPointer = 2

	fmt.Printf("n.i: %s, n.j: %d", n.i, n.j)
}

func main() {
	//unsafe1()
	unsafe2()
}
