// Go语言实例化结构体——为结构体分配内存并初始化

package main

import "fmt"

type Player struct {
	Name        string
	HealthPoint int
	MagicPoint  int
}

func main() {
	tank := new(Player)
	tank.Name = "Canon"
	tank.HealthPoint = 300
	fmt.Printf("%T\n", *tank)
	fmt.Println(tank)
}
