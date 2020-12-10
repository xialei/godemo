package main

import (
	"demo"
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
	demo.MD5Sum("/Users/roger/Downloads/") //os.Args[1]
}
