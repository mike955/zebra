package service

import (
	"context"
	"errors"

	"github.com/mike955/zebra/flake/internal/data"
	"github.com/sirupsen/logrus"

	pb "github.com/mike955/zebra/api/flake"
)

type FlakeService struct {
	pb.UnimplementedFlakeServiceServer
	logger *logrus.Entry
	data   *data.FlakeData
}

func NewFlakeService(logger *logrus.Entry) *FlakeService {
	return &FlakeService{
		logger: logger,
		data:   data.NewFlakeData(logger),
	}
}

func (s *FlakeService) New(ctx context.Context, request *pb.NewRequest) (response *pb.NewResponse, err error) {
	response = new(pb.NewResponse)
	if logger := ctx.Value("logger"); logger != nil {
		s.logger = logger.(*logrus.Entry)
		s.data.SetLogger(logger.(*logrus.Entry))
	}
	id, err := s.data.New(ctx)
	if err != nil {
		s.logger.Errorf("app:flake|service:flake|func:new|request:%+v|error:%s", request, err.Error())
		err = errors.New("generate id error")
		return
	}
	response.Data = id
	return
}
