package data

import (
	"github.com/mike955/zebra/account/internal/dao"
	"github.com/sirupsen/logrus"
)

type AccountData struct {
	logger *logrus.Logger
	dao    *dao.AccountDao
}

func NewAccountData(logger *logrus.Logger) *AccountData {
	return &AccountData{
		logger: logger,
		dao:    dao.NewAccountDao(),
	}
}
