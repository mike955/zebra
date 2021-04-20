package data

import (
	"context"
	"errors"
	"time"

	"github.com/mike955/zebra/pkg/ecrypto"

	"github.com/mike955/zebra/account/internal/dao"
	"github.com/mike955/zebra/account/internal/rpc"
	flake_pb "github.com/mike955/zebra/api/flake"
	"github.com/sirupsen/logrus"
)

type AccountData struct {
	logger *logrus.Logger
	dao    *dao.AccountDao
}

// data handle logic
func NewAccountData(logger *logrus.Logger) *AccountData {
	return &AccountData{
		logger: logger,
		dao:    dao.NewAccountDao(),
	}
}

func (s *AccountData) Create(ctx context.Context, params *CreateRequest) (err error) {
	// 1.check username、cellphone、email
	var account dao.Account
	s.dao.DB.Where("username=?", params.Username).Or("cellphone=?", params.Cellphone).Or("email=?", params.Email).Find(&account)
	if account.Username == params.Username {
		return errors.New("account has been exist")
	}
	if account.Cellphone == params.Cellphone {
		return errors.New("cellphone has been exist")
	}
	if account.Email == params.Email {
		return errors.New("email has been exist")
	}

	// 2.get id from flake
	res, err := rpc.FlakeRpc().New(ctx, &flake_pb.NewRequest{})
	if err != nil || res.Data == 0 {
		s.logger.Errorf("app:account|service:account|layer:data|func:create|info:call falke.New error|params:%+v|error:%s", params, err.Error())
		return errors.New("create id error")
	}

	// 3.create user
	// account
	account.Id = res.Data
	account.Username = params.Username
	account.Level = params.Level
	account.QQ = params.QQ
	account.Wechat = params.Wechat
	account.Cellphone = params.Cellphone
	account.Email = params.Email
	account.LastLoginTime = time.Now().Format("2006-01-02 15:04:05")

	account.Salt = ecrypto.GenerateRandomHex(64)
	account.Password = ecrypto.GeneratePassword(params.Password, account.Salt)

	err = s.dao.Create(account)
	if err != nil {
		s.logger.Errorf("app:account|service:account|layer:data|func:create|info:create account error|params:%+v|error:%s", params, err.Error())
		return errors.New("create account error")
	}
	return
}

func (s *AccountData) Deletes(ctx context.Context, params *DeletesRequest) (err error) {
	// 1.check params
	var accounts []dao.Account
	var fields map[string]interface{}
	fields["id"] = params.Ids
	fields["is_deleted"] = 1
	accounts, err = s.dao.FindByFields(fields)
	if err != nil {
		return
	}
	if len(accounts) != len(params.Ids) {
		return errors.New("accounts can not found")
	}

	// 2. delete accounts
	err = s.dao.DeleteByIds(params.Ids)
	return
}

func (s *AccountData) Gets(ctx context.Context, params *GetsRequest) (accounts []dao.Account, err error) {
	var fields map[string]interface{}
	fields["id"] = params.Ids
	query := s.dao.DB.Where("name IN ?", params.Ids)
	if params.Level != 0 {
		query.Where("level = ?", params.Level)
	}
	if params.Username != "" {
		query.Where("username LIKE ?", params.Username)
	}
	if params.QQ != "" {
		query.Where("qq = ?", params.QQ)
	}
	if params.Wechat != "" {
		query.Where("wechat = ?", params.Wechat)
	}
	if params.Cellphone != "" {
		query.Where("cellphone = ?", params.Cellphone)
	}
	if params.Email != "" {
		query.Where("level = ?", params.Email)
	}
	err = query.Find(&accounts).Error
	return
}
