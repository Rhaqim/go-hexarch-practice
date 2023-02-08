package grpc

import (
	"context"
	"log"

	"github.com/Rhaqim/hexarch/internal/adapters/framework/left/grpc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a Adapter) GetAddition(ctx context.Context, req *pb.OperationParams) (*pb.Answer, error) {
	log.Printf("Received: %v", req)

	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Errorf(codes.InvalidArgument, "a and b must be greater than 0")
	}

	result, err := a.api.GetAddition(int(req.GetA()), int(req.GetB()))
	if err != nil {
		return ans, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.Answer{Value: int32(result)}, nil
}

func (a Adapter) GetSubtraction(ctx context.Context, req *pb.OperationParams) (*pb.Answer, error) {
	log.Printf("Received: %v", req)

	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Errorf(codes.InvalidArgument, "a and b must be greater than 0")
	}

	result, err := a.api.GetSubtraction(int(req.GetA()), int(req.GetB()))
	if err != nil {
		return ans, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.Answer{Value: int32(result)}, nil
}

func (a Adapter) GetMultiplication(ctx context.Context, req *pb.OperationParams) (*pb.Answer, error) {
	log.Printf("Received: %v", req)

	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Errorf(codes.InvalidArgument, "a and b must be greater than 0")
	}

	result, err := a.api.GetMultiplication(int(req.GetA()), int(req.GetB()))
	if err != nil {
		return ans, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.Answer{Value: int32(result)}, nil
}

func (a Adapter) GetDivision(ctx context.Context, req *pb.OperationParams) (*pb.Answer, error) {
	log.Printf("Received: %v", req)

	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Errorf(codes.InvalidArgument, "a and b must be greater than 0")
	}

	result, err := a.api.GetDivision(int(req.GetA()), int(req.GetB()))
	if err != nil {
		return ans, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.Answer{Value: int32(result)}, nil
}
