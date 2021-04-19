package service

import (
	"context"

	"github.com/mike955/zebra/account/internal/data"
	pb "github.com/mike955/zebra/api/account"
	"github.com/sirupsen/logrus"
)

type AccountService struct {
	pb.UnimplementedAccountServiceServer
	logger *logrus.Logger
	data   *data.AccountData
}

func NewAccountService(logger *logrus.Logger) *AccountService {
	return &AccountService{
		logger: logger,
		data:   data.NewAccountData(logger),
	}
}

func (s *AccountService) Create(ctx context.Context, request *pb.CreateRequest) (response *pb.CreateResponse, err error) {
	return
}

func (s *AccountService) Delete(ctx context.Context, request *pb.DeleteRequest) (response *pb.DeleteResponse, err error) {
	return
}

func (s *AccountService) Deletes(ctx context.Context, request *pb.DeletesRequest) (response *pb.DeletesResponse, err error) {
	return
}

func (s *AccountService) Update(ctx context.Context, request *pb.UpdateRequest) (response *pb.UpdateResponse, err error) {
	return
}

func (s *AccountService) Get(ctx context.Context, request *pb.GetRequest) (response *pb.GetResponse, err error) {
	return
}

func (s *AccountService) Gets(ctx context.Context, request *pb.GetsRequest) (response *pb.GetsResponse, err error) {
	return
}
