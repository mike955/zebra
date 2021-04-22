package service

import (
	"context"
	"errors"

	"github.com/mike955/zebra/age/internal/data"
	"github.com/mike955/zebra/pkg/ecrypto"
	"github.com/sirupsen/logrus"

	pb "github.com/mike955/zebra/api/age"
)

type AgeService struct {
	pb.UnimplementedAgeServiceServer
	logger *logrus.Logger
	data   *data.AgeData
}

func NewAgeService(logger *logrus.Logger) *AgeService {
	return &AgeService{
		logger: logger,
		data:   data.NewAgeData(logger),
	}
}

func (s *AgeService) Get(ctx context.Context, request *pb.GetRequest) (response *pb.GetResponse, err error) {
	response = new(pb.GetResponse)
	if request.Age == 0 {
		request.Age = ecrypto.GenerateRandomUint64()
	}
	age, err := s.data.Get(ctx, request.Age)
	if err != nil {
		s.logger.Errorf("app:flake|service:flake|func:new|request:%+v|error:%s", request, err.Error())
		err = errors.New("get age error")
		return
	}
	response.Data = &pb.Age{
		Id:  age.Id,
		Age: age.Age,
	}
	return
}
