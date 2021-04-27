package service

import (
	"context"
	"errors"

	"github.com/mike955/zebra/bank/internal/data"
	"github.com/mike955/zebra/pkg/ecrypto"
	"github.com/sirupsen/logrus"

	pb "github.com/mike955/zebra/api/bank"
)

type BankService struct {
	pb.UnimplementedBankServiceServer
	logger *logrus.Entry
	data   *data.BankData
}

func NewBankService(logger *logrus.Entry) *BankService {
	return &BankService{
		logger: logger,
		data:   data.NewBankData(logger),
	}
}

func (s *BankService) Get(ctx context.Context, request *pb.GetRequest) (response *pb.GetResponse, err error) {
	response = new(pb.GetResponse)
	if logger := ctx.Value("logger"); logger != nil {
		s.logger = logger.(*logrus.Entry)
		s.data.SetLogger(logger.(*logrus.Entry))
	}
	if request.Bank == 0 {
		request.Bank = ecrypto.GenerateRandomUint64()
	}
	bank, err := s.data.Get(ctx, request.Bank)
	if err != nil {
		s.logger.Errorf("app:email|service:email|func:new|request:%+v|error:%s", request, err.Error())
		err = errors.New("get email error")
		return
	}
	response.Data = &pb.Bank{
		Id:   bank.Id,
		Bank: bank.Bank,
	}
	return
}
