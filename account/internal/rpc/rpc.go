package rpc

import (
	"context"
	"log"

	"github.com/mike955/zebra/account/configs"
	age_pb "github.com/mike955/zebra/api/age"
	flake_pb "github.com/mike955/zebra/api/flake"
	"google.golang.org/grpc"
)

var _gRPCClientMap = map[string]interface{}{}

type Rpc struct {
	Flake flake_pb.FlakeServiceClient
	Age   age_pb.AgeServiceClient
}

func NewRpc() *Rpc {
	return &Rpc{
		Flake: flakeRpc(),
		Age:   ageRpc(),
	}
}

func flakeRpc() flake_pb.FlakeServiceClient {
	if _gRPCClientMap["flake"] == nil {

		var ctx = context.Background()
		flakeAddr := configs.GlobalConfig.Rpc.FlakeAddr
		conn, err := grpc.DialContext(ctx, flakeAddr, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("user grpc client did not connect: %v", err)
		}
		client := flake_pb.NewFlakeServiceClient(conn)
		_gRPCClientMap["flake"] = client
	}
	return _gRPCClientMap["flake"].(flake_pb.FlakeServiceClient)
}

func ageRpc() age_pb.AgeServiceClient {
	if _gRPCClientMap["age"] == nil {

		var ctx = context.Background()
		ageAddr := configs.GlobalConfig.Rpc.AgeAddr
		conn, err := grpc.DialContext(ctx, ageAddr, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("user grpc client did not connect: %v", err)
		}
		client := age_pb.NewAgeServiceClient(conn)
		_gRPCClientMap["age"] = client
	}
	return _gRPCClientMap["age"].(age_pb.AgeServiceClient)
}
