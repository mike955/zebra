package rpc

import (
	"context"
	"os"
	"sync"

	flake_pb "github.com/mike955/zebra/api/flake"
	zrpc "github.com/mike955/zrpc/transform/grpc"
	"google.golang.org/grpc"
)

var gRPCClientMap sync.Map

func NewFlakeRpc(flakeAddr string) (flake_pb.FlakeServiceClient, error) {
	if _, ok := gRPCClientMap.Load("flake"); !ok {
		if os.Getenv("FlAKE_ADDR") != "" {
			flakeAddr = os.Getenv("FlAKE_ADDR")
		}
		conn, err := grpc.DialContext(context.Background(), flakeAddr, grpc.WithUnaryInterceptor(zrpc.ClientUnaryInterceptor), grpc.WithStreamInterceptor(zrpc.ClientStreamInterceptor), grpc.WithInsecure())
		if err != nil {
			return nil, err
		}
		gRPCClientMap.Store("flake", flake_pb.NewFlakeServiceClient(conn))
	}
	client, _ := gRPCClientMap.Load("flake")
	return client.(flake_pb.FlakeServiceClient), nil
}
