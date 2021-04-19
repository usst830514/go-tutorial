// Go语言append()为切片添加元素

package main

import "fmt"

func main() {
	//var a []int
	//a = append(a, 1) // 追加1个元素
	//fmt.Println(a)
	//a = append(a, 1, 2, 3) // 追加多个元素, 手写解包方式
	//fmt.Println(a)
	//a = append(a, []int{1,2,3}...) // 追加一个切片, 切片需要解包
	//fmt.Println(a)
	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("len: %d  cap: %d pointer: %p\n", len(numbers), cap(numbers), numbers)
	}
}
