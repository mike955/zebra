package service

import (
	"context"
	"errors"

	"github.com/mike955/zebra/email/internal/data"
	"github.com/mike955/zebra/pkg/ecrypto"
	"github.com/sirupsen/logrus"

	pb "github.com/mike955/zebra/api/email"
)

type EmailService struct {
	pb.UnimplementedEmailServiceServer
	logger *logrus.Logger
	data   *data.EmailData
}

func NewEmailService(logger *logrus.Logger) *EmailService {
	return &EmailService{
		logger: logger,
		data:   data.NewEmailData(logger),
	}
}

func (s *EmailService) Get(ctx context.Context, request *pb.GetRequest) (response *pb.GetResponse, err error) {
	response = new(pb.GetResponse)
	if request.Email == "" {
		request.Email = ecrypto.GenerateRandomString(50) + "@zebra.com"
	}
	email, err := s.data.Get(ctx, request.Email)
	if err != nil {
		s.logger.Errorf("app:email|service:email|func:new|request:%+v|error:%s", request, err.Error())
		err = errors.New("get email error")
		return
	}
	response.Data = &pb.Email{
		Id:    email.Id,
		Email: email.Email,
	}
	return
}
