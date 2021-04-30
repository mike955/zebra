package data

import (
	"context"
	"errors"

	"github.com/mike955/zebra/age/configs"
	"github.com/mike955/zebra/age/internal/dao"
	flake_pb "github.com/mike955/zebra/api/flake"
	"github.com/mike955/zebra/pkg/transform/grpc"
	"github.com/sirupsen/logrus"
)

type AgeData struct {
	logger *logrus.Entry
	dao    *dao.AgeDao
}

func NewAgeData(logger *logrus.Entry) *AgeData {
	return &AgeData{
		logger: logger,
		dao:    dao.NewAgeDao(),
	}
}

func (s *AgeData) SetLogger(logger *logrus.Entry) {
	s.logger = logger
}

// TODO(mike.cao): 递归每次返回不一样的 age
func (s *AgeData) Get(ctx context.Context, age uint64) (sage dao.Age, err error) {
	var fields = make(map[string]interface{})
	fields["age"] = age
	sages, err := s.dao.FindByFields(fields)
	if err != nil {
		s.logger.Errorf("app:age|data:age|func:get|info:check age error|params:%+d|error:%s", age, err.Error())
		err = errors.New("check age error")
		return
	}
	if len(sages) != 0 {
		return sages[0], nil
	}
	flakeRpc, err := grpc.NewFlakeRpc(configs.GlobalConfig.Rpc.FlakeAddr)
	if err != nil {
		s.logger.Errorf("app:age|data:age|func:get|info:create flake client error|params:%+d|error:%s", configs.GlobalConfig.Rpc.FlakeAddr, err.Error())
		err = errors.New("flake rpc call error")
		return
	}
	flakeRes, err := flakeRpc.New(ctx, &flake_pb.NewRequest{})
	if err != nil || flakeRes.Data == 0 {
		s.logger.Errorf("app:age|data:age|func:get|info:call falke.New error|params:%+d|error:%s", age, err.Error())
		err = errors.New("create id error")
		return
	}
	err = s.dao.Create(dao.Age{
		Id:  flakeRes.Data,
		Age: age,
	})
	sage.Id = flakeRes.Data
	sage.Age = age
	return
}
