// Go语言函数声明（函数定义）

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(test(3, 4))
}

// 普通函数声明（定义）
//func 函数名(形式参数列表)(返回值列表){
//	函数体
//}

func test(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}
func add(x int, y int) int   { return x + y }
func sub(x, y int) (z int)   { z = x - y; return }
func first(x int, _ int) int { return x }
func zero(int, int) int      { return 0 }

// 函数的返回值
// 1) 同一种类型返回值
//func test(x, y float64) (float64, float64) {
//	return math.Sqrt(x*x + y*y), math.Pi
//}

// 2) 带有变量名的返回值
//func test(x, y float64) (a, b float64) {
//	a = math.Sqrt(x*x + y*y)
//	b = math.Pi
//	return
//}
//
//func test(x, y=1.0 float64) float64 {
//	return math.Sqrt(x*x + y*y)
//}

// 调用函数
// 返回值变量列表 = 函数名(参数列表)
