package main

import (
	"context"
	"fmt"
	"log"
	"testing"

	pb "veltris.com/pkg/grpc"
	cal "veltris.com/veltris.com/calculator"
)
type greeterServer struct {
    pb.UnimplementedGreeterServer
	cal.UnimplementedCalculatorServer
}
type CalculatorServer struct {
	cal.UnimplementedCalculatorServer
}
func TestSayHello(t *testing.T) {
	// Create a new instance of the greeterServer
	server := &greeterServer{}

	// Prepare a mock context
	ctx := context.Background()

	// Prepare a mock HelloRequest
	req := &pb.HelloRequest{Name: "John"}

	// Call the SayHello method and check the response
	resp, err := server.SayHello(ctx, req)
	if err != nil {
		t.Errorf("SayHello failed: %v", err)
	}
	expectedMsg := "Hello John"
	if resp.Message != expectedMsg {
		t.Errorf("SayHello returned unexpected message: got %s, want %s", resp.Message, expectedMsg)
	}
}

func TestAdd(t *testing.T) {
	// Create a new instance of the greeterServer
	server := &CalculatorServer{}

	// Prepare a mock context
	ctx := context.Background()

	// Prepare a mock AddRequest
	req := &cal.AddRequest{Num1: 3, Num2: 5}

	// Call the Add method and check the response
	resp, err := server.Add(ctx, req)
	if err != nil {
		t.Errorf("Add failed: %v", err)
	}
	expectedResult := int32(8)
	if resp.Result != expectedResult {
		t.Errorf("Add returned unexpected result: got %d, want %d", resp.Result, expectedResult)
	}
}
func (s *CalculatorServer) Add(ctx context.Context, req *cal.AddRequest) (*cal.AddResponse, error) {
	fmt.Println(req.Num1,req.Num2)
	result := req.Num1 + req.Num2
	return &cal.AddResponse{Result: result}, nil
}
func (s *greeterServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}