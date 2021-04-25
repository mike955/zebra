package rpc

import (
	"context"
	"fmt"
	"os"

	flake_pb "github.com/mike955/zebra/api/flake"
	"github.com/mike955/zebra/email/configs"
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
		flakeAddr := configs.GlobalConfig.Rpc.FlakeAddr
		if os.Getenv("FlAKE_ADDR") != "" {
			flakeAddr = os.Getenv("FlAKE_ADDR")
		}
		conn, err := grpc.DialContext(ctx, flakeAddr, grpc.WithInsecure())
		if err != nil {
			fmt.Printf("user grpc client did not connect: %v\n", err)
		}
		client := flake_pb.NewFlakeServiceClient(conn)
		_gRPCClientMap["flake"] = client
	}
	return _gRPCClientMap["flake"].(flake_pb.FlakeServiceClient)
}
