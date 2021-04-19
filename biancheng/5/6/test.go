// Go语言匿名函数——没有函数名字的函数

package main

import "fmt"

// 定义一个匿名函数
//func(参数列表)(返回参数列表){
//	函数体
//}

func main() {
	// 1) 在定义时调用匿名函数
	func(data string) {
		fmt.Println("hello", data)
	}("world")

	// 2) 将匿名函数赋值给变量
	f := func(data string) {
		fmt.Println("hello", data)
	}
	f("world")

	visit([]int{1, 2, 3, 4}, func(v int) {
		fmt.Println(v)
	})
}

// 匿名函数用作回调函数
func visit(list []int, f func(int)) {
	for _, v := range list {
		f(v)
	}
}

// 使用匿名函数实现操作封装
