package main

import (
	"log"
	"sync"
	"fmt"
	"context"
	"google.golang.org/grpc"
	"services"
	"os"
	"time"
)

const (
	address     = "localhost:19999"
)

func doJob() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := services.NewComputeClient(conn)

	// Contact the server and print out its response.
	msg := "hello world"
	if len(os.Args) > 1 {
		msg = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.RunTask(ctx, &services.TaskRequest{Msg: msg})
	if err != nil {
			log.Fatalf("error in task: %v", err)
	}
	log.Printf("result: %s", r.GetResult())

}

// Schedule 跨语言任务调度器，基于GRPC
func Schedule() {

	fmt.Println("start scheduling")
	// 1. 调用任务拆分方法


	// 2. 协程执行job，失败重试
	var wg sync.WaitGroup

	wg.Add(10)

	for i:=0; i<10; i++ {
		go func(j int) {

			for k:=0; k<5; k++ {
				doJob()
			}

			wg.Done()
		}(i)
	}
	wg.Wait()


	// 3. 整理结果报表

	fmt.Println("finish scheduling")
}