// Go语言内嵌结构体成员名字冲突

package main

import "fmt"

type A struct {
	a int
}
type B struct {
	a int
}
type C struct {
	A
	B
}

func main() {
	c := &C{}
	//c := new(C)
	c.A.a = 1
	fmt.Println(c)
}
