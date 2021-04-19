// Go语言闭包（Closure）——引用了外部变量的匿名函数

package main

import "fmt"

func main() {
	/* 示例：闭包的记忆效应 */
	// 创建一个累加器, 初始值为1
	accumulator := Accumulate(1)
	fmt.Printf("%T\n", accumulator)
	// 累加1并打印
	fmt.Println(accumulator())
	fmt.Println(accumulator())
	// 打印累加器的函数地址
	fmt.Printf("%p\n", &accumulator)
	// 创建一个累加器, 初始值为10
	accumulator2 := Accumulate(10)
	// 累加1并打印
	fmt.Println(accumulator2())
	// 打印累加器的函数地址
	fmt.Printf("%p\n", &accumulator2)

	/* 示例：闭包实现生成器 */
	// 创建一个玩家生成器
	generator := playerGen("trump")
	// 返回玩家的名字和血量
	name, hp := generator()
	// 打印值
	fmt.Println(name, hp)
}

// 函数 + 引用环境 = 闭包

// 提供一个值, 每次调用函数会指定对值进行累加
func Accumulate(value int) func() int {
	// 返回一个闭包
	return func() int {
		// 累加
		value++
		// 返回一个累加值
		return value
	}
}

func playerGen(name string) func() (string, int) {
	// 血量一直为150
	hp := 150
	// 返回创建的闭包
	return func() (string, int) {
		// 将变量引用到闭包中
		return name, hp
	}
}
