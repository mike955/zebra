package rpc

import (
	"context"
	"log"
	"os"

	"github.com/mike955/zebra/account/configs"
	age_pb "github.com/mike955/zebra/api/age"
	bank_pb "github.com/mike955/zebra/api/bank"
	cellphone_pb "github.com/mike955/zebra/api/cellphone"
	email_pb "github.com/mike955/zebra/api/email"
	flake_pb "github.com/mike955/zebra/api/flake"
	"google.golang.org/grpc"
)

var _gRPCClientMap = map[string]interface{}{}

type Rpc struct {
	Flake     flake_pb.FlakeServiceClient
	Age       age_pb.AgeServiceClient
	Email     email_pb.EmailServiceClient
	Bank      bank_pb.BankServiceClient
	Cellphone cellphone_pb.CellphoneServiceClient
}

func NewRpc() *Rpc {
	return &Rpc{
		Flake:     flakeRpc(),
		Age:       ageRpc(),
		Email:     emailRpc(),
		Bank:      bankRpc(),
		Cellphone: cellphoneRpc(),
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
		if os.Getenv("AGE_ADDR") != "" {
			ageAddr = os.Getenv("AGE_ADDR")
		}
		conn, err := grpc.DialContext(ctx, ageAddr, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("user grpc client did not connect: %v", err)
		}
		client := age_pb.NewAgeServiceClient(conn)
		_gRPCClientMap["age"] = client
	}
	return _gRPCClientMap["age"].(age_pb.AgeServiceClient)
}

func emailRpc() email_pb.EmailServiceClient {
	if _gRPCClientMap["email"] == nil {

		var ctx = context.Background()
		emailAddr := configs.GlobalConfig.Rpc.EmailAddr
		if os.Getenv("EMAIL_ADDR") != "" {
			emailAddr = os.Getenv("EMAIL_ADDR")
		}
		conn, err := grpc.DialContext(ctx, emailAddr, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("user grpc client did not connect: %v", err)
		}
		client := email_pb.NewEmailServiceClient(conn)
		_gRPCClientMap["email"] = client
	}
	return _gRPCClientMap["email"].(email_pb.EmailServiceClient)
}

func bankRpc() bank_pb.BankServiceClient {
	if _gRPCClientMap["bank"] == nil {

		var ctx = context.Background()
		bankAddr := configs.GlobalConfig.Rpc.BankAddr
		if os.Getenv("BANK_ADDR") != "" {
			bankAddr = os.Getenv("BANK_ADDR")
		}
		conn, err := grpc.DialContext(ctx, bankAddr, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("user grpc client did not connect: %v", err)
		}
		client := bank_pb.NewBankServiceClient(conn)
		_gRPCClientMap["bank"] = client
	}
	return _gRPCClientMap["bank"].(bank_pb.BankServiceClient)
}

func cellphoneRpc() cellphone_pb.CellphoneServiceClient {
	if _gRPCClientMap["cellphone"] == nil {

		var ctx = context.Background()
		cellphoneAddr := configs.GlobalConfig.Rpc.CellphoneAddr
		if os.Getenv("CELLPHONE_ADDR") != "" {
			cellphoneAddr = os.Getenv("CELLPHONE_ADDR")
		}
		conn, err := grpc.DialContext(ctx, cellphoneAddr, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("user grpc client did not connect: %v", err)
		}
		client := cellphone_pb.NewCellphoneServiceClient(conn)
		_gRPCClientMap["cellphone"] = client
	}
	return _gRPCClientMap["cellphone"].(cellphone_pb.CellphoneServiceClient)
}
