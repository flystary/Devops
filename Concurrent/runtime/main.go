package main

import (
	"fmt"
	"runtime"
)

/*
func main() {
	//go func() {
	//	for i := 0; i < 2; i++ {
	//		fmt.Println("go")
	//	}
	//}()

	//for i := 0; i < 2; i++ {
	//	runtime.Gosched()
	//	fmt.Println("hello")
	//}
}
*/
func test() {
	defer fmt.Println("cccccccccccccccccccccccc")
	return
	//runtime.Goexit()
	fmt.Println("bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb")
}

func main() {
	//go func() {
	//	fmt.Println("aaaaaaaaaaaaaaaaaaaa")
	//	test()
	//	fmt.Println("ddddddddddddddddddddddddd")
	//}()
	//
	//for {}

	n := runtime.GOMAXPROCS(1)
	fmt.Println("n = ", n)

	for {
		go fmt.Print(1)
		fmt.Print(0)
	}
}

/*
第一次执行（runtime.GOMAXPROCS(1)）时，最多同时只能有一个goroutine被执行。所以会打印很多一。
过了一段时间后，GO语言调度器将其置为休眠，并唤醒另一个goroutine,这时候开始打印很多个0了，在打印的时候，goroutine是被调度到操作系统线程上的。
第二次执行(runtime.GOMAXPROCS(2))时，我们使用了两个CPU，所以两个goroutine可以一起被执行，以同样的频率交替打印0和1。

*/
