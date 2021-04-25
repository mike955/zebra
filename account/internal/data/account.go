package data

import (
	"context"
	"errors"
	"sync"

	"github.com/mike955/zebra/pkg/ecrypto"

	"github.com/mike955/zebra/account/internal/dao"
	"github.com/mike955/zebra/account/internal/rpc"

	age_pb "github.com/mike955/zebra/api/age"

	bank_pb "github.com/mike955/zebra/api/bank"
	cellphone_pb "github.com/mike955/zebra/api/cellphone"
	email_pb "github.com/mike955/zebra/api/email"
	flake_pb "github.com/mike955/zebra/api/flake"
	"github.com/sirupsen/logrus"
)

type AccountData struct {
	logger *logrus.Logger
	dao    *dao.AccountDao
	rpc    *rpc.Rpc
}

// data handle logic
func NewAccountData(logger *logrus.Logger) *AccountData {
	return &AccountData{
		logger: logger,
		dao:    dao.NewAccountDao(),
		rpc:    rpc.NewRpc(),
	}
}

func (s *AccountData) Get(ctx context.Context, username, password string) (accounts []dao.Account, err error) {
	// ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	// defer cancel()
	if username == "" {
		username = ecrypto.GenerateRandomString(16)
	}
	if password == "" {
		password = ecrypto.GenerateRandomString(16)
	}
	var fileds = make(map[string]interface{})
	fileds["username"] = username
	accounts, err = s.dao.FindByFields(fileds)
	if err != nil {
		s.logger.Errorf("app:account|data:account|func:get|info:get account error|params:%+v|error:%s", username, err.Error())
		err = errors.New("get account error")
		return
	}
	if len(accounts) != 0 {
		s.logger.Errorf("app:account|data:account|func:get|info:get account error|params:%+v|account:%s", username, accounts[0])
		err = errors.New("account exist error")
		return
	}

	var newAccount dao.Account
	var flake *flake_pb.NewResponse
	var age *age_pb.GetResponse
	var email *email_pb.GetResponse
	var bank *bank_pb.GetResponse
	var cellphone *cellphone_pb.GetResponse
	var wg sync.WaitGroup
	var errs []error

	wg.Add(5)
	// flake
	go func(ctx context.Context) {
		defer wg.Done()
		var err error
		flakeRes, err := s.rpc.Flake.New(ctx, &flake_pb.NewRequest{})
		if err != nil || flakeRes.Data == 0 {
			s.logger.Errorf("app:account|data:account|func:get|info:call falke.New error|params:%+v|error:%s", username, err.Error())
			errs = append(errs, errors.New("get id error"))
			return
		}
		if flakeRes.Code != 0 {
			s.logger.Errorf("app:account|data:account|func:get|info:call falke.New code error|params:%+v|error:%s|response:%+v", username, err.Error(), flakeRes)
			errs = append(errs, errors.New("get id code error"))
			return
		}
		flake = flakeRes
	}(ctx)

	// age
	go func(ctx context.Context) {
		defer wg.Done()
		ageRes, err := s.rpc.Age.Get(ctx, &age_pb.GetRequest{})
		if err != nil {
			s.logger.Errorf("app:account|data:account|func:get|info:call age.Get error|params:%+v|error:%s", nil, err.Error())
			errs = append(errs, errors.New("get age id error"))
			return
		}
		if ageRes.Code != 0 {
			s.logger.Errorf("app:account|data:account|func:get|info:call age.Get code error|params:%+v|error:%s|response:%+v", username, err.Error(), ageRes)
			errs = append(errs, errors.New("get age code error"))
			return
		}
		age = ageRes
	}(ctx)

	// // email
	go func(ctx context.Context) {
		defer wg.Done()
		emailRes, err := s.rpc.Email.Get(ctx, &email_pb.GetRequest{})
		if err != nil {
			s.logger.Errorf("app:account|data:account|func:get|info:call email.Get error|params:%+v|error:%s", username, err.Error())
			errs = append(errs, errors.New("get email error"))
			return
		}
		if emailRes.Code != 0 {
			s.logger.Errorf("app:account|data:account|func:get|info:call email.Get code error|params:%+v|error:%s|response:%+v", username, err.Error(), emailRes)
			errs = append(errs, errors.New("get email code error"))
			return
		}
		email = emailRes
	}(ctx)

	// bank
	go func(ctx context.Context) {
		defer wg.Done()
		bankRes, err := s.rpc.Bank.Get(ctx, &bank_pb.GetRequest{})
		if err != nil {
			s.logger.Errorf("app:account|data:account|func:get|info:call bank.Get error|params:%+v|error:%s", username, err.Error())
			errs = append(errs, errors.New("get bank error"))
			return
		}
		if bankRes.Code != 0 {
			s.logger.Errorf("app:account|data:account|func:get|info:call bank.Get code error|params:%+v|error:%s|response:%+v", username, err.Error(), bankRes)
			errs = append(errs, errors.New("get email code error"))
			return
		}
		bank = bankRes
	}(ctx)

	// cellphone
	go func(ctx context.Context) {
		defer wg.Done()
		cellphoneRes, err := s.rpc.Cellphone.Get(ctx, &cellphone_pb.GetRequest{})
		if err != nil {
			s.logger.Errorf("app:account|data:account|func:get|info:call cellphone.Get error|params:%+v|error:%s", username, err.Error())
			errs = append(errs, errors.New("get bank error"))
			return
		}
		if cellphoneRes.Code != 0 {
			s.logger.Errorf("app:account|data:account|func:get|info:call cellphone.Get code error|params:%+v|error:%s|response:%+v", username, err.Error(), cellphone)
			errs = append(errs, errors.New("get email code error"))
			return
		}
		cellphone = cellphoneRes
	}(ctx)

	wg.Wait()
	if len(errs) != 0 {
		s.logger.Errorf("app:account|data:account|func:get|info:get account error|params:%+v|error:%s", username, errs)
		errs = append(errs, errors.New("get account error"))
		return
	}

	newAccount.Id = (*flake).Data
	newAccount.Age = (*age).Data.Age
	newAccount.AgeId = (*age).Data.Id
	newAccount.Email = (*email).Data.Email
	newAccount.EmailId = (*email).Data.Id
	newAccount.Cellphone = (*cellphone).Data.Cellphone
	newAccount.CellphoneId = (*cellphone).Data.Id
	newAccount.Bank = (*bank).Data.Bank
	newAccount.BankId = (*bank).Data.Id
	newAccount.Username = username
	newAccount.Salt = ecrypto.GenerateRandomHex(64)
	newAccount.Password = ecrypto.GeneratePassword(password, newAccount.Salt)
	err = s.dao.Create(newAccount)
	return []dao.Account{newAccount}, nil
}

func (s *AccountData) Auth(ctx context.Context, params *AuthRequest) (account dao.Account, err error) {
	var fields = make(map[string]interface{})
	fields["username"] = params.Username
	accounts, err := s.dao.FindByFields(fields)
	if err != nil {
		return
	}
	if len(accounts) != 1 {
		err = errors.New("can not found account")
		return
	}
	account = accounts[0]
	if account.Password != ecrypto.GeneratePassword(params.Password, account.Salt) {
		err = errors.New("password error")
	}
	return
}
