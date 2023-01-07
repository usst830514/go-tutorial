//package main
//
//import (
//	"fmt"
//	"github.com/go-git/go-git/v5"
//	"github.com/go-git/go-git/v5/plumbing"
//	"github.com/go-git/go-git/v5/plumbing/object"
//	"strings"
//)

//func main() {
//	repo, _ := git.PlainOpen("/Users/lihaowei/cmdb.git")
//	//objectIter, err := repo.Objects()
//	//if err != nil {
//	//	panic(err)
//	//}
//	//objectIter.ForEach(func(obj object.Object) error {
//	//	fmt.Println(obj.Type())
//	//	return nil
//	//})
//	refs, _ := repo.Branches()
//	if err := refs.ForEach(func(r *plumbing.Reference) error {
//		fmt.Printf("name: %v, type: %v, hash:%v, target: %v\n", r.Name(), r.Type(), r.Hash(), r.Target())
//		return nil
//	}); err != nil {
//
//	}
//	ref, _ := repo.Head()
//	fmt.Printf("name: %v, type: %v, hash:%v, target: %v\n", ref.Name(), ref.Type(), ref.Hash(), ref.Target())
//	commit, _ := repo.CommitObject(ref.Hash())
//	tree, _ := commit.Tree()
//	if err := tree.Files().ForEach(func(f *object.File) error {
//		content, _ := f.Contents()
//		if strings.Contains(content, "offline_upload_url") {
//			fmt.Printf("name: %v, type: %v, hash: %v, size: %v, mode: %v, id: %v\n", f.Name, f.Type(), f.Hash, f.Size, f.Mode, f.ID())
//		}
//		return nil
//	}); err != nil {
//		panic(err)
//	}
//}

//package main
//
//import (
//	"fmt"
//	"path"
//)
//
//func main() {
//	fmt.Println(path.Base("a/b"))
//	fmt.Println(path.Base("/"))
//	fmt.Println(path.Base(""))
//}

//package main
//
//import (
//	"fmt"
//	"path"
//)
//
//func main() {
//	fmt.Println(path.Join("a", "b", "*"))
//	fmt.Println(path.Join("a", "b/c"))
//	fmt.Println(path.Join("a/b", "c"))
//
//	fmt.Println(path.Join("a/b", "../../../xyz"))
//
//	fmt.Println(path.Join("", ""))
//	fmt.Println(path.Join("a", ""))
//	fmt.Println(path.Join("", "a"))
//
//}

//package main
//
//import (
//	"fmt"
//	"path/filepath"
//)
//
//func main() {
//	files, _ := filepath.Glob("/Users/lihaowei/repositories/*")
//	fmt.Println(files) // contains a list of all files in the current directory
//}

//package main
//
//import (
//	"fmt"
//	"io/ioutil"
//)
//
//func main() {
//	files, _ := ioutil.ReadDir("/Users/lihaowei/repositories/")
//	for _, f := range files {
//		fmt.Println(f.Name())
//	}
//}
//

//package main
//
//import (
//	"fmt"
//	"regexp"
//)
//
//func main() {
//	// Compile the expression once, usually at init time.
//	// Use raw strings to avoid having to quote the backslashes.
//	a:="test"
//	var validID = regexp.MustCompile(a)
//
//	fmt.Println(validID.MatchString("haha tet haha"))
//}

//package main
//
//import (
//	"fmt"
//	"sync"
//	"time"
//)
//
//var (
//	Mu = sync.Mutex{}
//)
//
//func main() {
//
//	ch := make(chan string)
//
//	go sendData(ch)
//	getData(ch)
//
//	time.Sleep(time.Hour)
//}
//
//func sendData(ch chan string) {
//	for _, i := range []string{"a", "b", "c", "d", "e"} {
//		ch <- i
//	}
//}
//
//func getData(ch chan string) {
//	// time.Sleep(2e9)
//	for {
//		go func() {
//			//Mu.Lock()
//			fmt.Printf("%s\n", <-ch)
//			//Mu.Unlock()
//			time.Sleep(time.Second)
//		}()
//	}
//}

//package main
//
//import (
//	"fmt"
//	"sync"
//	"sync/atomic"
//	"time"
//
//	"github.com/panjf2000/ants/v2"
//)
//
//var sum int32
//
//func myFunc(i interface{}) {
//	n := i.(int32)
//	atomic.AddInt32(&sum, n)
//	fmt.Printf("run with %d\n", n)
//}
//
//func demoFunc() {
//	time.Sleep(10 * time.Millisecond)
//	fmt.Println("Hello World!")
//}
//
//func main() {
//	defer ants.Release()
//
//	runTimes := 20
//
//	// Use the common pool.
//	var wg sync.WaitGroup
//	syncCalculateSum := func() {
//		demoFunc()
//		wg.Done()
//	}
//	for i := 0; i < runTimes; i++ {
//		wg.Add(1)
//		_ = ants.Submit(syncCalculateSum)
//	}
//	wg.Wait()
//	fmt.Printf("running goroutines: %d\n", ants.Running())
//	fmt.Printf("finish all tasks.\n")
//
//	// Use the pool with a function,
//	// set 10 to the capacity of goroutine pool and 1 second for expired duration.
//	p, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
//		time.Sleep(5 * time.Second)
//		myFunc(i)
//		wg.Done()
//	})
//	defer p.Release()
//	// Submit tasks one by one.
//	for i := 0; i < runTimes; i++ {
//		wg.Add(1)
//		_ = p.Invoke(int32(i))
//	}
//	wg.Wait()
//	fmt.Printf("running goroutines: %d\n", p.Running())
//	fmt.Printf("finish all tasks, result is %d\n", sum)
//}

package main

import (
	"log"
	"runtime"
	"time"
)

func readMemStats() {

	var ms runtime.MemStats

	runtime.ReadMemStats(&ms)

	log.Printf(" ===> Alloc:%d(bytes) HeapIdle:%d(bytes) HeapReleased:%d(bytes)", ms.Alloc, ms.HeapIdle, ms.HeapReleased)
}

func test() {
	//slice 会动态扩容，用slice来做堆内存申请
	container := make([]int, 8)

	log.Println(" ===> loop begin.")
	for i := 0; i < 32*1000*1000; i++ {
		container = append(container, i)
		if i == 16*1000*1000 {
			readMemStats()
		}
	}

	log.Println(" ===> loop end.")
}

func main() {
	log.Println(" ===> [Start].")

	readMemStats()
	test()
	readMemStats()

	log.Println(" ===> [force gc].")
	runtime.GC() //强制调用gc回收

	log.Println(" ===> [Done].")
	readMemStats()

	go func() {
		for {
			readMemStats()
			time.Sleep(10 * time.Second)
		}
	}()

	time.Sleep(36000 * time.Second) //睡眠，保持程序不退出
}
