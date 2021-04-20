package dao

import (
	"gorm.io/gorm"
)

type Account struct {
	CommonTimeModel
	Id            uint64 `json:"id" gorm"primaryKey"`
	Username      string `json:"username"`
	Password      string `json:"password`
	Salt          string `json:"salt`
	Level         uint64 `json:"level"`
	QQ            string `json:"qq"`
	Wechat        string `json:"wechat"`
	Cellphone     string `json:"cellphone"`
	Email         string `json:"email"`
	State         uint64 `gorm:"default:0" json:"state"`
	LastLoginTime string `json:"last_login_time"`
	IsDeleted     uint64 `gorm:"default:0" json:"is_deleted"`
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
