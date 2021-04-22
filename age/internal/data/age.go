package data

import (
	"context"
	"errors"

	"github.com/mike955/zebra/age/internal/dao"
	"github.com/mike955/zebra/age/internal/rpc"
	flake_pb "github.com/mike955/zebra/api/flake"
	"github.com/sirupsen/logrus"
)

type AgeData struct {
	logger *logrus.Logger
	dao    *dao.AgeDao
	rpc    *rpc.Rpc
}

func NewAgeData(logger *logrus.Logger) *AgeData {
	return &AgeData{
		logger: logger,
		dao:    dao.NewAgeDao(),
		rpc:    rpc.NewRpc(),
	}
}

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
	flakeRes, err := s.rpc.Flake.New(ctx, &flake_pb.NewRequest{})
	if err != nil || flakeRes.Data == 0 {
		s.logger.Errorf("app:age|service:age|layer:data|func:get|info:call falke.New error|params:%+d|error:%s", age, err.Error())
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
