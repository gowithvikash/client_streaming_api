package main

import (
	"fmt"
	"io"
	"log"
	"net"

	pb "github.com/gowithvikash/grpc_with_go/client_streaming_api/proto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.GreetServiceServer
}

var (
	network = "tcp"
	address = "0.0.0.0:50051"
)

func main() {
	lis, err := net.Listen(network, address)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})
	if err = s.Serve(lis); err != nil {
		log.Fatal(err)
	}

}
func (s *Server) Long_Greet(stream pb.GreetService_Long_GreetServer) error {
	var res = ""
	fmt.Println("___ Long_Greet Function Was Invoked At Client___")
	for {
		fmt.Println("-Receiving Requests From Client")
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{Result: res})
		}
		if err != nil {
			log.Fatal(err)
		}

		// if err = stream.SendAndClose(&pb.GreetResponse{Result: res}); err != nil {
		// 	log.Fatal(err)
		// }
		res += fmt.Sprintf("Receiving Response :  Hello , %s !\n", req.Name)
	}

}
