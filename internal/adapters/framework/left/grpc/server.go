package rpc

import (
	"log"
	"net"

	"github.com/Rhaqim/hexarch/internal/adapters/framework/left/grpc/pb"
	"github.com/Rhaqim/hexarch/internal/ports"

	"google.golang.org/grpc"
)

type Adapter struct {
	api ports.APIPort
}

func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{
		api: api,
	}
}

func (a Adapter) Run() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	arithmeticServiceServer := a

	s := grpc.NewServer()

	pb.RegisterArithmeticServer(s, arithmeticServiceServer)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
