// Go语言type关键字（类型别名）
package main

import (
	"fmt"
)

// 区分类型别名与类型定义

// 将NewInt定义为int类型
type NewInt int

// 将int取一个别名叫IntAlias
type IntAlias = int

func main() {
	// 将a声明为NewInt类型
	var a NewInt
	// 查看a的类型名
	fmt.Printf("a type: %T\n", a)
	// 将a2声明为IntAlias类型
	var a2 IntAlias
	// 查看a2的类型名
	fmt.Printf("a2 type: %T\n", a2)
}
