package main

import (
	"log"
	"os"

	"github.com/Rhaqim/hexarch/internal/adapters/app/api"
	"github.com/Rhaqim/hexarch/internal/adapters/core/arithmetic"
	"github.com/Rhaqim/hexarch/internal/adapters/framework/right/db"
	"github.com/Rhaqim/hexarch/internal/ports"

	gRPC "github.com/Rhaqim/hexarch/internal/adapters/framework/left/grpc"
)

func main() {
	var err error

	var dbA ports.DBPort
	var apiA ports.APIPort
	var arithA ports.ArithmeticPort
	var grpcA ports.GRPCPort

	dbDriver := os.Getenv("DB_DRIVER")
	dbURL := os.Getenv("DB_URL")

	dbA, err = db.NewAdapter(dbDriver, dbURL)
	if err != nil {
		log.Fatal(err)
	}

	defer dbA.CloseDB()

	arithA = arithmetic.NewAdapter()

	apiA = api.NewAdapter(dbA, arithA)

	grpcA = gRPC.NewAdapter(apiA)

	grpcA.Run()

}
