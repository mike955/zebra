package data

import (
	"context"
	"errors"
	"sync"

	"github.com/mike955/zebra/pkg/ecrypto"
	"github.com/mike955/zebra/pkg/transform/grpc"

	"github.com/mike955/zebra/account/configs"
	"github.com/mike955/zebra/account/internal/dao"

	age_pb "github.com/mike955/zebra/api/age"

	bank_pb "github.com/mike955/zebra/api/bank"
	cellphone_pb "github.com/mike955/zebra/api/cellphone"
	email_pb "github.com/mike955/zebra/api/email"
	flake_pb "github.com/mike955/zebra/api/flake"
	"github.com/sirupsen/logrus"
)

type AccountData struct {
	logger *logrus.Entry
	dao    *dao.AccountDao
}

// data handle logic
func NewAccountData(logger *logrus.Entry) *AccountData {
	return &AccountData{
		logger: logger,
		dao:    dao.NewAccountDao(),
	}
}

func (s *AccountData) SetLogger(logger *logrus.Entry) {
	s.logger = logger
}

func (s *AccountData) Get(ctx context.Context, username, password string) (accounts []dao.Account, err error) {
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
		flakeRpc, err := grpc.NewFlakeRpc(configs.GlobalConfig.Rpc.FlakeAddr)
		if err != nil {
			s.logger.Errorf("app:account|data:account|func:get|info:create flake client error|params:%+d|error:%s", configs.GlobalConfig.Rpc.FlakeAddr, err.Error())
			errs = append(errs, errors.New("flake rpc call error"))
			return
		}
		flakeRes, err := flakeRpc.New(ctx, &flake_pb.NewRequest{})
		if err != nil || flakeRes.Code == 0 {
			s.logger.Errorf("app:account|data:account|func:get|info:call falke.New error|params:%+v|error:%s", username, err.Error())
			errs = append(errs, errors.New("get id error"))
			return
		}
		flake = flakeRes
	}(ctx)

	// age
	go func(ctx context.Context) {
		defer wg.Done()
		ageRpc, err := grpc.NewAgeRpc(configs.GlobalConfig.Rpc.AgeAddr)
		if err != nil {
			s.logger.Errorf("app:account|data:account|func:get|info:create flake client error|params:%+d|error:%s", configs.GlobalConfig.Rpc.FlakeAddr, err.Error())
			errs = append(errs, errors.New("ageRpc rpc call error"))
			return
		}
		ageRes, err := ageRpc.Get(ctx, &age_pb.GetRequest{})
		if err != nil || ageRes.Code != 0 {
			s.logger.Errorf("app:account|data:account|func:get|info:call age.Get error|params:%+v|error:%s", nil, err.Error())
			errs = append(errs, errors.New("get age id error"))
			return
		}
		age = ageRes
	}(ctx)

	// // email
	go func(ctx context.Context) {
		defer wg.Done()
		emailRpc, err := grpc.NewEmailRpc(configs.GlobalConfig.Rpc.EmailAddr)
		if err != nil {
			s.logger.Errorf("app:account|data:account|func:get|info:create flake client error|params:%+d|error:%s", configs.GlobalConfig.Rpc.FlakeAddr, err.Error())
			errs = append(errs, errors.New("ageRpc rpc call error"))
			return
		}
		emailRes, err := emailRpc.Get(ctx, &email_pb.GetRequest{})
		if err != nil || emailRes.Code != 0 {
			s.logger.Errorf("app:account|data:account|func:get|info:call email.Get error|params:%+v|error:%s", username, err.Error())
			errs = append(errs, errors.New("get email error"))
			return
		}
		email = emailRes
	}(ctx)

	// bank
	go func(ctx context.Context) {
		defer wg.Done()
		bankRpc, err := grpc.NewBankRpc(configs.GlobalConfig.Rpc.BankAddr)
		if err != nil {
			s.logger.Errorf("app:account|data:account|func:get|info:create flake client error|params:%+d|error:%s", configs.GlobalConfig.Rpc.FlakeAddr, err.Error())
			errs = append(errs, errors.New("ageRpc rpc call error"))
			return
		}
		bankRes, err := bankRpc.Get(ctx, &bank_pb.GetRequest{})
		if err != nil || bankRes.Code != 0 {
			s.logger.Errorf("app:account|data:account|func:get|info:call bank.Get error|params:%+v|error:%s", username, err.Error())
			errs = append(errs, errors.New("get bank error"))
			return
		}
		bank = bankRes
	}(ctx)

	// cellphone
	go func(ctx context.Context) {
		defer wg.Done()
		cellphoneRpc, err := grpc.NewCellphoneRpc(configs.GlobalConfig.Rpc.CellphoneAddr)
		if err != nil {
			s.logger.Errorf("app:account|data:account|func:get|info:create flake client error|params:%+d|error:%s", configs.GlobalConfig.Rpc.FlakeAddr, err.Error())
			errs = append(errs, errors.New("ageRpc rpc call error"))
			return
		}
		cellphoneRes, err := cellphoneRpc.Get(ctx, &cellphone_pb.GetRequest{})
		if err != nil || cellphoneRes.Code != 0 {
			s.logger.Errorf("app:account|data:account|func:get|info:call cellphone.Get error|params:%+v|error:%s", username, err.Error())
			errs = append(errs, errors.New("get bank error"))
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
