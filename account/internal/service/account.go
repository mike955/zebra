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
		Password:  req.Password,
		Level:     req.Level,
		QQ:        req.Qq,
		Wechat:    req.Wechat,
		Cellphone: req.Cellphone,
		Email:     req.Email,
	})
	if err != nil {
		s.logger.Errorf("app:account|service:account|layer:service|func:create|info:create account error|params:%+v|error:%s", req, err.Error())
		response.Msg = "create account error"
	}
	return
}

func (s *AccountService) Delete(ctx context.Context, req *pb.DeleteRequest) (response *pb.DeleteResponse, err error) {
	response = new(pb.DeleteResponse)
	err = s.data.Deletes(ctx, &data.DeletesRequest{
		Ids: []uint64{req.Id},
	})
	if err != nil {
		s.logger.Errorf("app:account|service:account|layer:service|func:delete|info:delete account error|params:%+v|error:%s", req, err.Error())
		response.Msg = "delete account error"
	}
	return
}

func (s *AccountService) Deletes(ctx context.Context, req *pb.DeletesRequest) (response *pb.DeletesResponse, err error) {
	response = new(pb.DeletesResponse)
	err = s.data.Deletes(ctx, &data.DeletesRequest{
		Ids: req.Ids,
	})
	if err != nil {
		s.logger.Errorf("app:account|service:account|layer:service|func:deletes|info:delete accounts error|params:%+v|error:%s", req, err.Error())
		response.Msg = "delete account error"
	}
	return
}

func (s *AccountService) Update(ctx context.Context, req *pb.UpdateRequest) (response *pb.UpdateResponse, err error) {
	response = new(pb.UpdateResponse)
	return
}

// TODO(mike.cai): add offset and limit
func (s *AccountService) Get(ctx context.Context, req *pb.GetRequest) (response *pb.GetResponse, err error) {
	response = new(pb.GetResponse)
	params := &data.GetsRequest{}
	if req.Id != 0 {
		params.Ids = []uint64{req.Id}
	}
	if req.Username != "" {
		params.Username = req.Username
	}
	if req.Level != 0 {
		params.Level = req.Level
	}
	if req.Qq != "" {
		params.QQ = req.Qq
	}
	if req.Wechat != "" {
		params.Wechat = req.Wechat
	}
	if req.Cellphone != "" {
		params.Cellphone = req.Cellphone
	}
	if req.Email != "" {
		params.Email = req.Email
	}
	accounts, err := s.data.Gets(ctx, params)
	if err != nil {
		s.logger.Errorf("app:account|service:account|layer:service|func:gets|info:gets accounts error|params:%+v|error:%s", req, err.Error())
		response.Msg = "delete account error"
		return
	}
	response.Data = &pb.Account{
		Id:        accounts[0].Id,
		Username:  accounts[0].Username,
		Level:     accounts[0].Level,
		Qq:        accounts[0].QQ,
		Wechat:    accounts[0].Wechat,
		Cellphone: accounts[0].Cellphone,
		Email:     accounts[0].Email,
	}
	return
}

// TODO(mike.cai): add offset and limit
func (s *AccountService) Gets(ctx context.Context, req *pb.GetsRequest) (response *pb.GetsResponse, err error) {
	response = new(pb.GetsResponse)

	params := &data.GetsRequest{}
	if len(req.Ids) > 0 {
		params.Ids = req.Ids
	}
	if req.Username != "" {
		params.Username = req.Username
	}
	if req.Level != 0 {
		params.Level = req.Level
	}
	if req.Qq != "" {
		params.QQ = req.Qq
	}
	if req.Wechat != "" {
		params.Wechat = req.Wechat
	}
	if req.Cellphone != "" {
		params.Cellphone = req.Cellphone
	}
	if req.Email != "" {
		params.Email = req.Email
	}
	accounts, err := s.data.Gets(ctx, params)
	if err != nil {
		s.logger.Errorf("app:account|service:account|layer:service|func:gets|info:gets accounts error|params:%+v|error:%s", req, err.Error())
		response.Msg = "delete account error"
		return
	}
	var acc []*pb.Account
	for _, a := range accounts {
		account := &pb.Account{
			Id:        a.Id,
			Username:  a.Username,
			Level:     a.Level,
			Qq:        a.QQ,
			Wechat:    a.Wechat,
			Cellphone: a.Cellphone,
			Email:     a.Email,
		}
		acc = append(acc, account)
	}
	response.Data = acc
	return
}

func (s *AccountService) Auth(ctx context.Context, req *pb.AuthRequest) (response *pb.AuthResponse, err error) {
	response = new(pb.AuthResponse)
	account, err := s.data.Auth(ctx, &data.AuthRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		s.logger.Errorf("app:account|service:account|layer:service|func:auth|info:auth accounts error|params:%+v|error:%s", req, err.Error())
		response.Msg = "auth account error"
		return
	}
	response.Data = &pb.Account{
		Id:        account.Id,
		Username:  account.Username,
		Level:     account.Level,
		Qq:        account.QQ,
		Wechat:    account.Wechat,
		Cellphone: account.Cellphone,
		Email:     account.Email,
	}
	return
}
