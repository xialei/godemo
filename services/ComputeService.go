package services

import (
	"log"
	"context"
	// pb "services/compute"
)

// ComputeService is used to implement helloworld.GreeterServer.
type ComputeService struct {
	UnimplementedComputeServer
}

// RunTask implements RunTask
func (s *ComputeService) RunTask(ctx context.Context, req *TaskRequest) (*TaskResponse, error) {
	log.Printf("Received: %v", req.GetMsg())
	return &TaskResponse{Result:"Hello " + req.GetMsg()}, nil
}



