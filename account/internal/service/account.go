package service

import (
	"context"

	"github.com/mike955/zebra/account/internal/data"
	pb "github.com/mike955/zebra/api/account"
	"github.com/sirupsen/logrus"
)

type AccountService struct {
	pb.UnimplementedAccountServiceServer
	logger *logrus.Entry
	data   *data.AccountData
}

// service check params
func NewAccountService(logger *logrus.Entry) *AccountService {
	return &AccountService{
		logger: logger,
		data:   data.NewAccountData(logger),
	}
}

func (s *AccountService) Get(ctx context.Context, req *pb.GetRequest) (response *pb.GetResponse, err error) {
	if logger := ctx.Value("logger"); logger != nil {
		s.logger = logger.(*logrus.Entry)
		s.data.SetLogger(logger.(*logrus.Entry))
	}
	response = new(pb.GetResponse)
	accounts, err := s.data.Get(ctx, req.Username, req.Password)
	if err != nil {
		s.logger.Errorf("app:account|service:account|func:gets|info:get account error|params:%+v|error:%s", req, err.Error())
		response.Msg = "get account error"
		return
	}
	response.Data = &pb.Account{
		Id:        accounts[0].Id,
		Username:  accounts[0].Username,
		Age:       accounts[0].Age,
		Email:     accounts[0].Email,
		Bank:      accounts[0].Bank,
		Cellphone: accounts[0].Cellphone,
	}
	return
}

func (s *AccountService) Auth(ctx context.Context, req *pb.AuthRequest) (response *pb.AuthResponse, err error) {
	if logger := ctx.Value("logger"); logger != nil {
		s.logger = logger.(*logrus.Entry)
		s.data.SetLogger(logger.(*logrus.Entry))
	}
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
		Age:       account.Age,
		Email:     account.Email,
		Bank:      account.Bank,
		Cellphone: account.Cellphone,
	}
	return
}
