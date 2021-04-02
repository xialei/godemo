package main

import (
	"fmt"
	"godemo/demo"
	"sync"
)

func startGraphQLServer() {

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		msg, _ := demo.StartGraphQL()
		if msg == "exit" {
			wg.Done()
		}
	}()

	wg.Wait()
}

func test() {
	var a []string
	a = append(a, "testa")
	fmt.Print(a)
}

func main() {
	// demo.DemoObjectOriented()
	// user.DemoUser()
	// demo.Flow()
	// demo.DemoFile()
	// demo.DemoGoroutine()
	// demo.DemoClickhouse()
	// startGraphQLServer()
	// demo.DemoObserver()
	// demo.DemoVisitor()
	// demo.DBReadPerformance()
	// startGraphQLServer()

	// demo pipeline
	// core.MD5Sum("/Users/roger/Downloads/") //os.Args[1]

	// demo.TestDgraph()
	demo.TestEncrypt()

	// test()
}
