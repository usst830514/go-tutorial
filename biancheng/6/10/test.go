// Go语言初始化内嵌结构体

package main

import "fmt"

// 车轮
type Wheel struct {
	Size int
}

//// 引擎
//type Engine struct {
//	Power int    // 功率
//	Type  string // 类型
//}
//
//// 车
//type Car struct {
//	Wheel
//	Engine
//}

/* 初始化内嵌匿名结构体 */
type Car struct {
	Wheel
	// 引擎
	Engine struct {
		Power int    // 功率
		Type  string // 类型
	}
}

func main() {
	//c := Car{
	//	// 初始化轮子
	//	Wheel: Wheel{
	//		Size: 18,
	//	},
	//
	//	// 初始化引擎
	//	Engine: Engine{
	//		Type:  "1.4T",
	//		Power: 143,
	//	},
	//}
	//c := Car{
	//	// 初始化轮子
	//	Wheel{
	//		Size: 18,
	//	},
	//
	//	// 初始化引擎
	//	Engine{
	//		Type:  "1.4T",
	//		Power: 143,
	//	},
	//}
	c := Car{
		// 初始化轮子
		Wheel: Wheel{
			Size: 18,
		},

		// 初始化引擎
		Engine: struct {
			Power int
			Type  string
		}{
			Type:  "1.4T",
			Power: 143,
		},
	}
	fmt.Println(c)
	fmt.Printf("%+v\n", c)
}
