// Go语言defer（延迟执行语句）

package main

import (
	"fmt"
	"os"
	"sync"
)

var (
	// 一个演示用的映射
	valueByKey = make(map[string]int)
	// 保证使用映射时的并发安全的互斥锁
	valueByKeyGuard sync.Mutex
)

func main() {
	/* 多个延迟执行语句的处理顺序 */
	fmt.Println("defer begin")
	// 将defer放入延迟调用栈，用于打开和关闭文件、接收请求和回复请求、加锁和解锁等
	defer fmt.Println(1)
	panic("crash")
	defer fmt.Println(2)
	// 最后一个放入, 位于栈顶, 最先调用
	defer fmt.Println(3)
	fmt.Println("defer end")

	/* 使用延迟执行语句在函数退出时释放资源 */
	// 1) 使用延迟并发解锁
	//readValue("test")
	//// 2) 使用延迟释放文件句柄
	//fmt.Println("文件大小为", fileSize("/etc/resolv.conf"))
}

func readValue(key string) int {
	valueByKeyGuard.Lock()
	// defer后面的语句不会马上调用, 而是延迟到函数结束时调用
	defer valueByKeyGuard.Unlock()
	return valueByKey[key]
}

func fileSize(filename string) int64 {
	f, err := os.Open(filename)
	if err != nil {
		return 0
	}
	// 延迟调用Close, 此时Close不会被调用
	defer f.Close()
	info, err := f.Stat()
	if err != nil {
		// defer机制触发, 调用Close关闭文件
		return 0
	}
	size := info.Size()
	// defer机制触发, 调用Close关闭文件
	return size
}
