package main

import (
	"sync"
	"test"
)

func startGraphQLServer() {

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		msg, _ := test.StartGraphQL()
		if msg == "exit" {
			wg.Done()
		}
	}()

	wg.Wait()
}

func main() {
	// test.DemoObjectOriented()

	// user.DemoUser()

	// test.Flow()

	// test.DemoFile()

	// test.DemoGoroutine()
	// test.DemoClickhouse()
	// startGraphQLServer()
	// test.DemoObserver()
	// test.DemoVisitor()
	test.DBReadPerformance()
}
