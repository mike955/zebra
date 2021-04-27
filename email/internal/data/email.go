package data

import (
	"context"
	"errors"

	flake_pb "github.com/mike955/zebra/api/flake"
	"github.com/mike955/zebra/email/internal/dao"
	"github.com/mike955/zebra/email/internal/rpc"
	"github.com/sirupsen/logrus"
)

type EmailData struct {
	logger *logrus.Entry
	dao    *dao.EmailDao
	rpc    *rpc.Rpc
}

func NewEmailData(logger *logrus.Entry) *EmailData {
	return &EmailData{
		logger: logger,
		dao:    dao.NewEmailDao(),
		rpc:    rpc.NewRpc(),
	}
}

func (s *EmailData) SetLogger(logger *logrus.Entry) {
	s.logger = logger
}

func (s *EmailData) Get(ctx context.Context, emailName string) (email dao.Email, err error) {
	var fields = make(map[string]interface{})
	fields["email"] = emailName
	emails, err := s.dao.FindByFields(fields)
	if err != nil {
		s.logger.Errorf("app:email|data:age|func:get|info:check email error|params:%+d|error:%s", emailName, err.Error())
		err = errors.New("check age error")
		return
	}
	if len(emails) != 0 {
		return emails[0], nil
	}
	flakeRes, err := s.rpc.Flake.New(ctx, &flake_pb.NewRequest{})
	if err != nil || flakeRes.Data == 0 {
		s.logger.Errorf("app:email|service:email|layer:data|func:get|info:call falke.New error|params:%+d|error:%s", email, err.Error())
		err = errors.New("create id error")
		return
	}
	err = s.dao.Create(dao.Email{
		Id:    flakeRes.Data,
		Email: emailName,
	})
	email.Id = flakeRes.Data
	email.Email = emailName
	return
}
