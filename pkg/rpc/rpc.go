package rpc

import (
	account_pb "github.com/mike955/zebra/api/account"
	flake_pb "github.com/mike955/zebra/api/flake"
)

type Rpc struct {
	Account *account_pb.AccountServiceClient
	Flake   *flake_pb.FlakeServiceClient
}
