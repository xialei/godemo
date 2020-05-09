package main

import (
	// "log"
	// "net"
	// "google.golang.org/grpc"
	// "services"
)

const (
	port = ":19999"
)

// func main() {

// 	lis, err := net.Listen("tcp", port)
// 	if err != nil {
// 			log.Fatalf("failed to listen: %v", err)
// 	}
// 	rpcServer := grpc.NewServer()
// 	services.RegisterComputeServer(rpcServer, new(services.ComputeService))
	
// 	if err := rpcServer.Serve(lis);err != nil {
// 		log.Fatalf("failed to serve: %v", err)
// 	}
	
// }