package dao

import "gorm.io/gorm"

type Account struct {
	CommonTimeModel
	ID        uint64 `json:"id" gorm"index"`
	Username  string `json:"username"`
	level     uint64 `json:"level"`
	qq        string `json:"qq"`
	wechat    string `json:"wechat"`
	cellphone string `json:"cellphone"`
	IsDelete  uint64 `gorm:"default:0" json:"is_delete"`
}

type AccountDao struct {
}

func NewAccountDao() *AccountDao {
	return &AccountDao{}
}

func (dao AccountDao) Add(data map[string]interface{}) (err error) {
	return nil
}

func (dao AccountDao) FindByFields(fields map[string]interface{}) ([]*Account, error) {
	var accounts []*Account
	err := DB.Where(fields).Find(&accounts).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return accounts, nil
}

func (dao AccountDao) DeleteByIds(ids []uint64) (err error) {
	if err := DB.Where(&map[string]interface{}{"id": ids}).Update("is_delete", 1).Error; err != nil {
		return err
	}
	return nil
}

func (dao AccountDao) UpdateFields(where map[string]interface{}, updateFileds map[string]interface{}) (err error) {
	// err = DB.Where(&where).Update(updateFileds).Error
	return
}
