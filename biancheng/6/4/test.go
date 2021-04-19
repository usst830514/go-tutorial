// Go语言构造函数

package main

import "fmt"

type Cat struct {
	Color string
	Name  string
}
type BlackCat struct {
	Sex string
	Cat // 嵌入Cat, 类似于派生
}

func NewCatByName(name string) *Cat {
	return &Cat{
		Name: name,
	}
}
func NewCatByColor(color string) *Cat {
	return &Cat{
		Color: color,
	}
}

// “构造基类”
func NewCat(name string) *Cat {
	return &Cat{
		Name: name,
	}
}

// “构造子类”
func NewBlackCat(color, name, sex string) *BlackCat {
	cat := &BlackCat{}
	cat.Color = color
	cat.Name = name
	cat.Sex = sex
	return cat
}
func main() {
	BlackCat := NewBlackCat("black", "tom", "male")
	fmt.Println(BlackCat)
}
