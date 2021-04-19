// Go语言函数变量——把函数作为值保存到变量中

package main

import (
	"fmt"
)

func fire() {
	fmt.Println("fire")
}
func main() {
	//var f func()
	//f = fire
	f := fire
	f()
}
