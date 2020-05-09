package test

import (
	"fmt"
	"sync"
	"time"
	"strconv"
)

//DemoGoroutine 协程, golang从语言层面支持并发
//  * 并行：parallelism, 每个任务在不同的物理处理器上运行
//  * 并发：currency, 时间片轮换
//  * 多线程：每个线程只处理一个请求，请求处理结束，接收下一个请求，高并发下性能开销大
//  * 基于回调的异步IO：程序运行过程中可能产生大量回调导致维护成本高
//  * 协程：不需要抢占式调用，可以有效提升线程的并发性
//	* 通过读写锁、互斥锁来防止多个协程对全局资源的资源竞争
//  go 表达式
func DemoGoroutine() {

	fmt.Println("demo goroutine")
	// testWaitGroup()
	// testMutex()
	// testRWMutex()
	// testChannel2()
	testSelect()
}

func testWaitGroup() {
	var wg1 sync.WaitGroup
	wg1.Add(10)
	for i:=0; i<5; i++ {
		tmp := i + 1
		go func(){
			fmt.Println(tmp) // 如果直接打印i，那么因为主协程的循环很快就跑完了，各个协程才拿到结果5，都打印了5
			time.Sleep(3e9)
			wg1.Done()
		}()
		// 更好的写法如下：
		go func(t2 int){
			fmt.Println(t2 * 10 + 10)
			time.Sleep(3e9)
			wg1.Done()
		}(i)
	}

	wg1.Wait()
	fmt.Println("the end of testWaitGroup")
}

var (
	tickets = 10
	wg sync.WaitGroup
	mutex sync.Mutex
)

//testMutex 当多个goroutine访问同一个函数，且函数在操作一个全局资源，为了保障共享资源的安全性，使用sync.Mutex对内容加锁，牺牲一些性能
//	使用 go run -race 查看资源竞争
func testMutex() {
	wg.Add(10)
	for i:=0; i<10; i++ {
		go func(){
			mutex.Lock()
			for j:=0; j<10; j++ {
				tickets = tickets - 1
			}
			mutex.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("the end of testMutex")
}

//testRWMutex 互斥锁表示同一代码只有一个goroutine能运行，而读写锁表示在锁范围内数据的读写操作
//     Go语言map不是线程安全的，多个goroutine同时操作会出现错误，可以用读写锁
func testRWMutex() {
	var rwm sync.RWMutex
	kv := make(map[string]string)
	var wg2 sync.WaitGroup
	wg2.Add(10)
	for i:=0; i<10; i++ {
		go func(j int){
			rwm.Lock()
			kv["k" + strconv.Itoa(j)] = "v" + strconv.Itoa(j)
			fmt.Println(kv)
			rwm.Unlock()
			wg2.Done()
		}(i)
	}
	wg2.Wait()
	fmt.Println("the end of testRWMutex")
}

//testChannel chan提供了goroutine之间的通信或同步, chan是安全的。
//    	var 名称 chan 类型
//		var 名称 chan <- 类型 // 只写，向通道添加值
//		var 名称 <- chan 类型 // 只读，从通道取值
//		名称:=make(chan int) // 无缓存channel
//		名称:=make(chan int, 0) // 无缓存channel
//		名称:=make(chan int, 100) // 有缓存channel，不超出缓存个数，就不会阻塞
func testChannel1() {
	// 实现同步，主协程和子协程的通信
	ch := make(chan int)
	go func(i int) {
		fmt.Println("enter goroutine")
		ch <- i // 放在有效代码最后面

		close(ch)
	}(10)
	c, d := <- ch // 阻塞直至取出
	fmt.Println(c, d) // 10 true
	fmt.Println("the end of testChannel1")
}

func testChannel2() {
	
	ch1 := make(chan string) // 子协程和子协程的通信
	ch2 := make(chan int) // 实现同步，等待子协程执行完，主协程结束

	go func() {
		fmt.Println("testChannel2.goroutine1")
		for i:=97; i<97+26; i++ {
			ch1 <- "hi. there? " + fmt.Sprintf("%c", i)
		}
		ch2 <- 1

	}()
	go func() {
		fmt.Println("testChannel2.goroutine2")
		for msg := range ch1{
			fmt.Println("get msg from 1 :", msg)
		}
	}()
	
	<-ch2
	fmt.Println("the end of testChannel2")
}

func testSelect() {
	ch := make(chan int)
	for i:=0; i<10; i++ {
		go func(j int) {
			ch <- j
		}(i)
	}
	for {// 一直监听，需要default，否则会出现死锁
		select {
		case a := <-ch:
			fmt.Println(a)
		default:
		}
	}
	
}

