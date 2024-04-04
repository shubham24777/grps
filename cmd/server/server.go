package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	pb "veltris.com/pkg/grpc" // Import the generated code package
	cal "veltris.com/veltris.com/calculator"
)

// greeterServer is the implementation of the Greeter service.
type greeterServer struct {
    pb.UnimplementedGreeterServer
	// cal.UnimplementedCalculatorServer
}
type CalculatorServer struct {
	cal.UnimplementedCalculatorServer
}

// SayHello implements the SayHello RPC method of the Greeter service.
func (s *greeterServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	// log.Printf("Received: %v", in.GetName())
	// return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
	if in.GetName() == "shubham" {
        // Create a new context with a timeout of 20 seconds
        timeoutCtx, cancel := context.WithTimeout(ctx, 20*time.Second)
        defer cancel()

        // Create a channel to signal when the response is ready
        ch := make(chan *pb.HelloReply)

        // Run the server logic in a goroutine
        go func() {
            // Simulate a delay of 20 seconds for requests with the name "shubham"
            time.Sleep(20 * time.Second)

            // Send the response back on the channel
            ch <- &pb.HelloReply{Message: "Hello " + in.GetName()}
        }()

        // Wait for either the response or the context timeout
        select {
        case <-timeoutCtx.Done():
            return nil, errors.New("request timed out")
        case reply := <-ch:
            return reply, nil
        }
    }

    // For requests with other names, proceed normally
    log.Printf("Received: %v", in.GetName())
    return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}
func (s *CalculatorServer) Add(ctx context.Context, req *cal.AddRequest) (*cal.AddResponse, error) {
	fmt.Println(req.Num1,req.Num2)
	result := req.Num1 + req.Num2
	return &cal.AddResponse{Result: result}, nil
}
func main() {
	/*lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &greeterServer{}) // Register the server implementation
	cal.RegisterCalculatorServer(s, &CalculatorServer{}) // Register the server implementation
	log.Println("Server started at :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}*/
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &greeterServer{})
	cal.RegisterCalculatorServer(s, &CalculatorServer{})

	go func() {
		log.Println("Server started at :50051")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Block indefinitely to keep the server running
	select {}
}
