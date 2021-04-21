package rpc

import (
	"context"
	"fmt"
	"log"

	"github.com/mike955/zebra/age/configs"
	flake_pb "github.com/mike955/zebra/api/flake"
	"google.golang.org/grpc"
)

var _gRPCClientMap = map[string]interface{}{}

type Rpc struct {
	Flake flake_pb.FlakeServiceClient
}

func NewRpc() *Rpc {
	return &Rpc{
		Flake: flakeRpc(),
	}
}

func flakeRpc() flake_pb.FlakeServiceClient {
	if _gRPCClientMap["flake"] == nil {

		var ctx = context.Background()
		fmt.Println("== %+V: ", configs.GlobalConfig.Rpc)
		flakeAddr := configs.GlobalConfig.Rpc.FlakeAddr
		fmt.Println("== flakeAddr: ", flakeAddr)
		conn, err := grpc.DialContext(ctx, flakeAddr, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("user grpc client did not connect: %v", err)
		}
		client := flake_pb.NewFlakeServiceClient(conn)
		_gRPCClientMap["flake"] = client
	}
	return _gRPCClientMap["flake"].(flake_pb.FlakeServiceClient)
}
