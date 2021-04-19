// Go语言初始化结构体的成员变量

package main

import "fmt"

type People struct {
	name  string
	child *People
}

type Address struct {
	Province    string
	City        string
	ZipCode     int
	PhoneNumber string
}

// 打印消息类型, 传入匿名结构体
func printMsgType(msg *struct {
	id   int
	data string
}) {
	// 使用动词%T打印msg的类型
	fmt.Printf("%T\n", msg)
}

func main() {
	//relation := &People{
	//	name: "爷爷",
	//	child: &People{
	//		name: "爸爸",
	//		child: &People{
	//			name: "我",
	//		},
	//	},
	//}
	//fmt.Println(relation)

	//addr := Address{
	//	"四川",
	//	"成都",
	//	610000,
	//	"0",
	//}
	//fmt.Println(addr)

	msg := &struct { // 定义部分
		id   int
		data string
	}{ // 值初始化部分
		1024,
		"hello",
	}
	fmt.Println(msg)
	printMsgType(msg)
}
