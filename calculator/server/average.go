package main

import (
	"io"
	"log"

	pb "github.com/fabiotavarespr/grpc-go-course/calculator/proto"
)

func (s *Server) Average(stream pb.CalculatorService_AverageServer) error {
	log.Printf("Average function was invoked")

	var sum int32 = 0
	count := 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AverageResponse{
				Result: float64(sum) / float64(count),
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream %v\n", err)
		}

		log.Printf("Receiving number: %d\n", req.Number)

		sum += req.Number
		count++
	}
}
