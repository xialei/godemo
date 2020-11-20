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
<<<<<<< HEAD
	// startGraphQLServer()
	// test.DemoObserver()
	// test.DemoVisitor()
	test.DBReadPerformance()
=======
	startGraphQLServer()
>>>>>>> fc011bf682ae2da86b6cf3baf21d6cbaf2509e2a
}
