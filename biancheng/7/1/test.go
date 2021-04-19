// Go语言接口声明（定义）

package main

func main() {
	/* 接口声明的格式 */
	//type 接口类型名 interface{
	//	方法名1( 参数列表1 ) 返回值列表1
	//	方法名2( 参数列表2 ) 返回值列表2
	//…
	//}
	/* 开发中常见的接口及写法 */
	type Writer interface {
		Write(p []byte) (n int, err error)
	}
	type Stringer interface {
		String() string
	}
}
