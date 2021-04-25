package rpc

import (
	"context"
	"log"

	account_pb "github.com/mike955/zebra/api/account"
	flake_pb "github.com/mike955/zebra/api/flake"
	"github.com/mike955/zebra/gateway/configs"
	"google.golang.org/grpc"
)

var _gRPCClientMap = map[string]interface{}{}

var Account account_pb.AccountServiceClient = accountRpc()
var Flake flake_pb.FlakeServiceClient = flakeRpc()

type Rpc struct {
	Account *account_pb.AccountServiceClient
	Flake   *flake_pb.FlakeServiceClient
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

func accountRpc() account_pb.AccountServiceClient {
	if _gRPCClientMap["account"] == nil {

		var ctx = context.Background()
		flakeAddr := configs.GlobalConfig.Rpc.AccountAddr
		conn, err := grpc.DialContext(ctx, flakeAddr, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("user grpc client did not connect: %v", err)
		}
		client := flake_pb.NewFlakeServiceClient(conn)
		_gRPCClientMap["account"] = client
	}
	return _gRPCClientMap["account"].(account_pb.AccountServiceClient)
}
