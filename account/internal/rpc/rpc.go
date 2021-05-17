package rpc

import (
	"context"
	"os"
	"sync"

	age_pb "github.com/mike955/zebra/api/age"
	bank_pb "github.com/mike955/zebra/api/bank"
	cellphone_pb "github.com/mike955/zebra/api/cellphone"
	email_pb "github.com/mike955/zebra/api/email"
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

func NewAgeRpc(ageAddr string) (age_pb.AgeServiceClient, error) {
	if _, ok := gRPCClientMap.Load("age"); !ok {
		if os.Getenv("AGE_ADDR") != "" {
			ageAddr = os.Getenv("AGE_ADDR")
		}
		conn, err := grpc.DialContext(context.Background(), ageAddr, grpc.WithUnaryInterceptor(zrpc.ClientUnaryInterceptor), grpc.WithStreamInterceptor(zrpc.ClientStreamInterceptor), grpc.WithInsecure())
		if err != nil {
			return nil, err
		}
		gRPCClientMap.Store("age", age_pb.NewAgeServiceClient(conn))
	}
	client, _ := gRPCClientMap.Load("age")
	return client.(age_pb.AgeServiceClient), nil
}

func NewEmailRpc(emailAddr string) (email_pb.EmailServiceClient, error) {
	if _, ok := gRPCClientMap.Load("email"); !ok {
		if os.Getenv("EMAIL_ADDR") != "" {
			emailAddr = os.Getenv("EMAIL_ADDR")
		}
		conn, err := grpc.DialContext(context.Background(), emailAddr, grpc.WithUnaryInterceptor(zrpc.ClientUnaryInterceptor), grpc.WithStreamInterceptor(zrpc.ClientStreamInterceptor), grpc.WithInsecure())
		if err != nil {
			return nil, err
		}
		gRPCClientMap.Store("email", email_pb.NewEmailServiceClient(conn))
	}
	client, _ := gRPCClientMap.Load("email")
	return client.(email_pb.EmailServiceClient), nil
}

func NewBankRpc(bankAddr string) (bank_pb.BankServiceClient, error) {
	if _, ok := gRPCClientMap.Load("bank"); !ok {
		if os.Getenv("BANK_ADDR") != "" {
			bankAddr = os.Getenv("BANK_ADDR")
		}
		conn, err := grpc.DialContext(context.Background(), bankAddr, grpc.WithUnaryInterceptor(zrpc.ClientUnaryInterceptor), grpc.WithStreamInterceptor(zrpc.ClientStreamInterceptor), grpc.WithInsecure())
		if err != nil {
			return nil, err
		}
		gRPCClientMap.Store("bank", bank_pb.NewBankServiceClient(conn))
	}
	client, _ := gRPCClientMap.Load("bank")
	return client.(bank_pb.BankServiceClient), nil
}

func NewCellphoneRpc(cellphoneAddr string) (cellphone_pb.CellphoneServiceClient, error) {
	if _, ok := gRPCClientMap.Load("cellphone"); !ok {
		if os.Getenv("CELLPHONE_ADDR") != "" {
			cellphoneAddr = os.Getenv("CELLPHONE_ADDR")
		}
		conn, err := grpc.DialContext(context.Background(), cellphoneAddr, grpc.WithUnaryInterceptor(zrpc.ClientUnaryInterceptor), grpc.WithStreamInterceptor(zrpc.ClientStreamInterceptor), grpc.WithInsecure())
		if err != nil {
			return nil, err
		}
		gRPCClientMap.Store("cellphone", cellphone_pb.NewCellphoneServiceClient(conn))
	}
	client, _ := gRPCClientMap.Load("cellphone")
	return client.(cellphone_pb.CellphoneServiceClient), nil
}
