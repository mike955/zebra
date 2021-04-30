package data

import (
	"context"
	"errors"

	flake_pb "github.com/mike955/zebra/api/flake"
	"github.com/mike955/zebra/bank/configs"
	"github.com/mike955/zebra/bank/internal/dao"
	"github.com/mike955/zebra/pkg/transform/grpc"
	"github.com/sirupsen/logrus"
)

type BankData struct {
	logger *logrus.Entry
	dao    *dao.BankDao
}

func NewBankData(logger *logrus.Entry) *BankData {
	return &BankData{
		logger: logger,
		dao:    dao.NewBankDao(),
	}
}

func (s *BankData) SetLogger(logger *logrus.Entry) {
	s.logger = logger
}

func (s *BankData) Get(ctx context.Context, bank uint64) (email dao.Bank, err error) {
	var fields = make(map[string]interface{})
	fields["bank"] = bank
	banks, err := s.dao.FindByFields(fields)
	if err != nil {
		s.logger.Errorf("app:email|data:bank|func:get|info:check bank error|params:%+d|error:%s", bank, err.Error())
		err = errors.New("check bank error")
		return
	}
	if len(banks) != 0 {
		return banks[0], nil
	}
	flakeRpc, err := grpc.NewFlakeRpc(configs.GlobalConfig.Rpc.FlakeAddr)
	if err != nil {
		s.logger.Errorf("app:email|data:age|func:get|info:create flake client error|params:%+d|error:%s", configs.GlobalConfig.Rpc.FlakeAddr, err.Error())
		err = errors.New("flake rpc call error")
		return
	}
	flakeRes, err := flakeRpc.New(ctx, &flake_pb.NewRequest{})
	if err != nil || flakeRes.Data == 0 {
		s.logger.Errorf("app:email|data:bank|func:get|info:call falke.New error|params:%+d|error:%s", email, err.Error())
		err = errors.New("create id error")
		return
	}
	err = s.dao.Create(dao.Bank{
		Id:   flakeRes.Data,
		Bank: bank,
	})
	email.Id = flakeRes.Data
	email.Bank = bank
	return
}
