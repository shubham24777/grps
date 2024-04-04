// package main

// import (
//     "context"
//     "log"

//     "google.golang.org/grpc"

//     pb "veltris.com/pkg/grpc" // Update this import path with your actual import path
//     cal "veltris.com/veltris.com/calculator"
// )

// func main() {
//     /*conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
//     if err != nil {
//         log.Fatalf("did not connect: %v", err)
//     }
//     defer conn.Close()
//     c := pb.NewGreeterClient(conn)
//     c1 := cal.NewCalculatorClient(conn)

//     name := "Veltris"
//     r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
//     if err != nil {
//         log.Fatalf("could not greet: %v", err)
//     }
//     log.Printf("Greeting: %s", r.Message)
//     resp, err := c1.Add(context.Background(), &cal.AddRequest{Num1: 1,Num2: 3})
//     if err != nil {
//         log.Fatalf("could not add: %v", err)
//     }
//     log.Printf("ADDED: %v", resp.Result)*/

// }
package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"google.golang.org/grpc"
	pb "veltris.com/pkg/grpc"
	// cal "veltris.com/veltris.com/calculator"
)

func main() {
	// Set up a connection to the server.
	/*conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a client for the Greeter service.
	greeterClient := pb.NewGreeterClient(conn)

	// Create a client for the Calculator service.
	calculatorClient := cal.NewCalculatorClient(conn)

	// Create a context with a timeout of 5 seconds.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Send requests to the server concurrently.
	go func() {
		// Call the SayHello method of the Greeter service.
		resp, err := greeterClient.SayHello(ctx, &pb.HelloRequest{Name: "Alice"})
		if err != nil {
			log.Fatalf("Error calling SayHello: %v", err)
		}
		log.Printf("Response from SayHello: %s", resp.GetMessage())
	}()

	go func() {
		// Call the Add method of the Calculator service.
		resp, err := calculatorClient.Add(ctx, &cal.AddRequest{Num1: 10, Num2: 20})
		if err != nil {
			log.Fatalf("Error calling Add: %v", err)
		}
		log.Printf("Response from Add: %d", resp.GetResult())
	}()

	// Sleep to allow time for requests to complete.
	time.Sleep(2 * time.Second)*/
    /*lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &greeterServer{})

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			conn, err := grpc.Dial(":50051", grpc.WithInsecure())
			if err != nil {
				log.Fatalf("did not connect: %v", err)
			}
			defer conn.Close()
			client := pb.NewGreeterClient(conn)
			name := fmt.Sprintf("User%d", i)
			resp, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: name})
			if err != nil {
				log.Printf("Error calling SayHello: %v", err)
				return
			}
			log.Printf("Response from server: %s", resp.GetMessage())
		}(i)
	}

	log.Println("Server started at :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	wg.Wait()*/
      // Create a gRPC client
      conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
      if err != nil {
          log.Fatalf("failed to dial server: %v", err)
      }
      defer conn.Close()

      // Create a Greeter client
      client := pb.NewGreeterClient(conn)

      // Create a context with a timeout of 30 seconds
      ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
      defer cancel()

      // Create a channel to receive responses
      ch := make(chan *pb.HelloReply)

      // Concurrently make requests with different names
      go func() {
          // Make a request with name "shubham"
          fmt.Println("20 sec bad milte hai")
          resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "shubham"})
          if err != nil {
              log.Printf("RPC failed for shubham: %v", err)
          } else {
              ch <- resp
          }
      }()
      startTime := time.Now()
      go func() {
		for {
			// Get the current time
			currentTime := time.Now()

			// Calculate the elapsed time
			elapsedTime := currentTime.Sub(startTime)

			// Print the elapsed time
			fmt.Printf("\rElapsed time: %s", elapsedTime)

			// Sleep for 1 second before updating again
			time.Sleep(1 * time.Second)
		}
	}()
// for i:=0; i<9;i++{
//     go func() {
//         // Make a request with a different name
//         str:= strconv.Itoa(i)
//         resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "other"+str})
//         if err != nil {
//             log.Printf("RPC failed for other: %v", err)
//         } else {
//             ch <- resp
//         }
//     }()
// }

for i := 0; i < 9; i++ {
    go func(i int) {
        // Make a request with a different name
        str := strconv.Itoa(i)
        resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "other" + str})
        if err != nil {
            log.Printf("RPC failed for other%s: %v", str, err)
        } else {
            ch <- resp
        }
    }(i)
}
      // Wait for responses from both requests
      for i := 0; i < 10; i++ {
          select {
          case resp := <-ch:
              log.Printf("Response from server: %s", resp.GetMessage())
          case <-ctx.Done():
              log.Println("Context timeout exceeded")
              return
          }
      }
}
