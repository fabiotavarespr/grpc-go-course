package main

import (
	"io"
	"log"

	pb "github.com/fabiotavarespr/grpc-go-course/greet/proto"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Printf("GreetEveryone function was invoked")
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error reading client stream: %v\n", err)
		}

		res := "Hello " + req.FirstName + "!"

		err = stream.Send(&pb.GreetResponse{
			Result: res,
		})

		if err != nil {
			log.Fatalf("Error while sending data to client: %v\n", err)
		}

	}
}
