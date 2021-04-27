package data

import (
	"context"
	"errors"

	flake_pb "github.com/mike955/zebra/api/flake"
	"github.com/mike955/zebra/cellphone/internal/dao"
	"github.com/mike955/zebra/cellphone/internal/rpc"
	"github.com/sirupsen/logrus"
)

type CellphoneData struct {
	logger *logrus.Entry
	dao    *dao.CellphoneDao
	rpc    *rpc.Rpc
}

func NewCellphoneData(logger *logrus.Entry) *CellphoneData {
	return &CellphoneData{
		logger: logger,
		dao:    dao.NewCellphoneDao(),
		rpc:    rpc.NewRpc(),
	}
}

func (s *CellphoneData) SetLogger(logger *logrus.Entry) {
	s.logger = logger
}

func (s *CellphoneData) Get(ctx context.Context, cellphone uint64) (cp dao.Cellphone, err error) {
	var fields = make(map[string]interface{})
	fields["cellphone"] = cellphone
	cps, err := s.dao.FindByFields(fields)
	if err != nil {
		s.logger.Errorf("app:cellphone|data:cellphone|func:get|info:check cellphone error|params:%+d|error:%s", cellphone, err.Error())
		err = errors.New("check age error")
		return
	}
	if len(cps) != 0 {
		return cps[0], nil
	}
	flakeRes, err := s.rpc.Flake.New(ctx, &flake_pb.NewRequest{})
	if err != nil || flakeRes.Data == 0 {
		s.logger.Errorf("app:cellphone|data:cellphone|func:get|info:call falke.New error|params:%+d|error:%s", cellphone, err.Error())
		err = errors.New("create id error")
		return
	}
	err = s.dao.Create(dao.Cellphone{
		Id:        flakeRes.Data,
		Cellphone: cellphone,
	})
	cp.Id = flakeRes.Data
	cp.Cellphone = cellphone
	return
}
