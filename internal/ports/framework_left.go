package ports

import (
	"context"
)

type GRPCPort interface {
	Run()
	GetAddition(ctx context.Context, a, b int) (int32, error)
	GetSubtraction(ctx context.Context, a, b int) (int32, error)
	GetMultiplication(ctx context.Context, a, b int) (int32, error)
	GetDivision(ctx context.Context, a, b int) (int32, error)
}
