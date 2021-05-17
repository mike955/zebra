package service

import (
	"context"
	"errors"

	"github.com/mike955/zebra/cellphone/internal/data"
	"github.com/mike955/zrpc/utils/ecrypto"
	"github.com/sirupsen/logrus"

	pb "github.com/mike955/zebra/api/cellphone"
)

type CellphoneService struct {
	pb.UnimplementedCellphoneServiceServer
	logger *logrus.Entry
	data   *data.CellphoneData
}

func NewCellphoneService(logger *logrus.Entry) *CellphoneService {
	return &CellphoneService{
		logger: logger,
		data:   data.NewCellphoneData(logger),
	}
}

func (s *CellphoneService) Get(ctx context.Context, request *pb.GetRequest) (response *pb.GetResponse, err error) {
	response = new(pb.GetResponse)
	if logger := ctx.Value("logger"); logger != nil {
		s.logger = logger.(*logrus.Entry)
		s.data.SetLogger(logger.(*logrus.Entry))
	} else {
		ctx = context.WithValue(ctx, "logger", s.logger)
	}
	if request.Cellphone == 0 {
		request.Cellphone = ecrypto.GenerateRandomUint64()
	}
	cellphone, err := s.data.Get(ctx, request.Cellphone)
	if err != nil {
		s.logger.Errorf("app:email|service:email|func:new|request:%+v|error:%s", request, err.Error())
		err = errors.New("get email error")
		return
	}
	response.Data = &pb.Cellphone{
		Id:        cellphone.Id,
		Cellphone: cellphone.Cellphone,
	}
	return
}
