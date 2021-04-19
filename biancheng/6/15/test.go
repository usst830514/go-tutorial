// Go语言链表操作

package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func ShowNode(p *Node) { //遍历
	for p != nil {
		fmt.Println(*p)
		p = p.next //移动指针
	}
}

func main() {
	/* 单向链表 */
	// 使用 Struct 定义单链表
	//var head = new(Node)
	//head.data = 1
	//var node1 = new(Node)
	//node1.data = 2
	//head.next = node1
	//var node2 = new(Node)
	//node2.data = 3
	//node1.next = node2
	//ShowNode(head)
	// 插入结点
	// 1) 头插法
	//var head = new(Node)
	//head.data = 0
	//var tail *Node
	//tail = head //tail用于记录头结点的地址，刚开始tail的的指针指向头结点
	//for i := 1; i < 10; i++ {
	//	var node = Node{data: i}
	//	node.next = tail //将新插入的node的next指向头结点
	//	tail = &node     //重新赋值头结点
	//}
	//ShowNode(tail) //遍历结果
	// 2) 尾插法
	var head = new(Node)
	head.data = 0
	var tail *Node
	tail = head
	for i := 0; i < 10; i++ {
		var node = Node{data: i}
		tail.next = &node
		tail = &node
	}
	ShowNode(head)
}
