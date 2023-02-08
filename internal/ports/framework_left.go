package ports

import (
	"context"

	"github.com/Rhaqim/hexarch/internal/adapters/framework/left/grpc/pb"
)

type GRPCPort interface {
	Run()
	GetAddition(ctx context.Context, req *pb.OperationParams) (*pb.Answer, error)
	GetSubtraction(ctx context.Context, req *pb.OperationParams) (*pb.Answer, error)
	GetMultiplication(ctx context.Context, req *pb.OperationParams) (*pb.Answer, error)
	GetDivision(ctx context.Context, req *pb.OperationParams) (*pb.Answer, error)
}
