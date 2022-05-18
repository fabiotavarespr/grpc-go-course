package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/fabiotavarespr/grpc-go-course/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) {
	log.Println("doMax was invoked")

	stream, err := c.Max(context.Background())

	if err != nil {
		log.Fatalf("Error while opening the stream: %v\n", err)
	}

	waitc := make(chan struct{})

	go func(){
		numbers :=[]int32{3,5,1,36,13,51,43,69}
		for _, number := range numbers {
			log.Printf("Sending number: %d\n", number)
			stream.Send(&pb.MaxRequest{
				Number: number,
			})
			time.Sleep(1*time.Second)
		}
		stream.CloseSend()
	}()

	go func(){
		for{
			res, err := stream.Recv()
			if err == io.EOF{
				break
			}
			if err != nil{
				log.Printf("Error while reading server stream: %v\n", err)
				break
			}

			log.Printf("Received a new maximum: %d\n", res.Result)
		}
		close(waitc)
	}()

	<-waitc

}
