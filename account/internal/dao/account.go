package dao

import (
	"gorm.io/gorm"
)

type Account struct {
	CommonTimeModel
	Id          uint64 `json:"id" gorm"primaryKey"`
	Username    string `json:"username"`
	Password    string `json:"password`
	Salt        string `json:"salt`
	Age         uint64 `json:"age`
	AgeId       uint64 `json:"age_id`
	Email       string `json:"email"`
	EmailId     uint64 `json:"email_id`
	Bank        uint64 `json:"bank"`
	BankId      uint64 `json:"bank_id`
	CellphoneId uint64 `json:"cellphone_id`
	Cellphone   uint64 `json:"cellphone"`
	IsDeleted   uint64 `gorm:"default:0" json:"is_deleted"`
}

type AccountDao struct {
	DB *gorm.DB
}

func NewAccountDao() *AccountDao {
	return &AccountDao{
		DB: DB,
	}
}

func (dao *AccountDao) Create(data Account) (err error) {
	err = DB.Create(&data).Error
	return
}

func (dao *AccountDao) FindByFields(fields map[string]interface{}) ([]Account, error) {
	var accounts []Account
	err := DB.Where(fields).Find(&accounts).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return accounts, nil
}

func (dao *AccountDao) DeleteByIds(ids []uint64) (err error) {
	if err := DB.Where(&map[string]interface{}{"id": ids}).Update("is_deleted", 1).Error; err != nil {
		return err
	}
	return nil
}

func (dao *AccountDao) UpdateFields(where map[string]interface{}, updateFileds map[string]interface{}) (err error) {
	// err = DB.Where(&where).Update(updateFileds).Error
	return
}
