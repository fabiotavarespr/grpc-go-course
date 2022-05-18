package main

import (
	"context"
	"log"

	pb "github.com/fabiotavarespr/grpc-go-course/calculator/proto"
)

func doAverage(c pb.CalculatorServiceClient) {
	log.Println("doAverage was invoked")

	stream, err := c.Average(context.Background())

	if err != nil {
		log.Fatalf("Error while opening the stream: %v\n", err)
	}

	numbers := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, number := range numbers {
		log.Printf("Sending number: %d\n", number)

		stream.Send(&pb.AverageRequest{
			Number: number,
		})
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response: %v\n", err)
	}

	log.Printf("Average: %f\n", res.Result)

}
