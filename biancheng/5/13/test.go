// Go语言宕机（panic）——程序终止运行

package main

import (
	"fmt"
	"net"
)

func main() {
	// 在运行依赖的必备资源缺失时主动触发宕机
	conn, error := net.Dial("tcp", "10.64.57.0:80")
	if error != nil {
		panic("无法连接10.64.57.0:80")
	}
	fmt.Printf("%T", conn)
}
