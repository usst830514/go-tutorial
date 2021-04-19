// Go语言类型与接口的关系

package main

import "io"

/* 一个类型可以实现多个接口 */
type Socket struct {
}

func (s *Socket) Write(p []byte) (n int, err error) {
	return 0, nil
}
func (s *Socket) Close() error {
	return nil
}

type Writer interface {
	Write(p []byte) (n int, err error)
}
type Closer interface {
	Close() error
}

// 使用io.Writer的代码, 并不知道Socket和io.Closer的存在
func usingWriter(writer io.Writer) {
	writer.Write(nil)
}

// 使用io.Closer, 并不知道Socket和io.Writer的存在
func usingCloser(closer io.Closer) {
	closer.Close()
}

/* 多个类型可以实现相同的接口 */
// 一个服务需要满足能够开启和写日志的功能
type Service interface {
	Start()     // 开启服务
	Log(string) // 日志输出
}

// 日志器
type Logger struct {
}

// 实现Service的Log()方法
func (g *Logger) Log(l string) {
}

// 游戏服务
type GameService struct {
	Logger // 嵌入日志器
}

// 实现Service的Start()方法
func (g *GameService) Start() {
}

func main() {
	/* 一个类型可以实现多个接口 */
	// 实例化Socket
	//s := new(Socket)
	//usingWriter(s)
	//usingCloser(s)
	/* 多个类型可以实现相同的接口 */
	var s Service = new(GameService)
	s.Start()
	s.Log("hello")
}
