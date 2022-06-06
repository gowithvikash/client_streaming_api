package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/gowithvikash/grpc_with_go/client_streaming_api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := pb.NewGreetServiceClient(conn)
	do_Long_Greet(c)

}

func do_Long_Greet(c pb.GreetServiceClient) {
	fmt.Println("\n_______________ Action Number : 01 _______________")
	fmt.Println("_____  do_Long_Greet Function Was Invoked At Client  ____")

	var reqs = []*pb.GreetRequest{{Name: "Bijender Kumar"}, {Name: "Vikash Parashar"}, {Name: "Khushboo Panday"}, {Name: "Niyati Parashar"}, {Name: "Ritika Parashar"}, {Name: "Rampati Devi"}}
	stream, err := c.Long_Greet(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range reqs {
		fmt.Println("Client Is Sending Requests To Server : ", v)
		err = stream.Send(v)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("do_Long_Greet_Result: %v\n", res.Result)

}
