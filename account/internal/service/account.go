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

// service check params
func NewAccountService(logger *logrus.Logger) *AccountService {
	return &AccountService{
		logger: logger,
		data:   data.NewAccountData(logger),
	}
}

func (s *AccountService) Create(ctx context.Context, req *pb.CreateRequest) (response *pb.CreateResponse, err error) {
	response = new(pb.CreateResponse)
	err = s.data.Create(ctx, &data.CreateRequest{
		Username:  req.Username,
		Level:     req.Level,
		QQ:        req.Qq,
		Wechat:    req.Wechat,
		Cellphone: req.Cellphone,
		Email:     req.Email,
	})
	if err != nil {

	}
	return
}

func (s *AccountService) Delete(ctx context.Context, req *pb.DeleteRequest) (response *pb.DeleteResponse, err error) {
	response = new(pb.DeleteResponse)
	return
}

func (s *AccountService) Deletes(ctx context.Context, req *pb.DeletesRequest) (response *pb.DeletesResponse, err error) {
	response = new(pb.DeletesResponse)
	return
}

func (s *AccountService) Update(ctx context.Context, req *pb.UpdateRequest) (response *pb.UpdateResponse, err error) {
	response = new(pb.UpdateResponse)
	return
}

func (s *AccountService) Get(ctx context.Context, req *pb.GetRequest) (response *pb.GetResponse, err error) {
	response = new(pb.GetResponse)
	return
}

func (s *AccountService) Gets(ctx context.Context, req *pb.GetsRequest) (response *pb.GetsResponse, err error) {
	response = new(pb.GetsResponse)
	return
}

func (s *AccountService) Auth(ctx context.Context, req *pb.AuthRequest) (response *pb.AuthResponse, err error) {
	response = new(pb.AuthResponse)
	return
}
