package dao

import (
	"fmt"

	"gorm.io/gorm"
)

type Bank struct {
	CommonTimeModel
	Id        uint64 `json:"id" gorm"primaryKey"`
	Bank      uint64 `json:"bank"`
	IsDeleted uint64 `json:"is_deleted" gorm:"default:0" `
}

type BankDao struct {
	DB *gorm.DB
}

func NewBankDao() *BankDao {
	return &BankDao{
		DB: DB,
	}
}

func (dao *BankDao) Create(data Bank) (err error) {
	err = DB.Create(&data).Error
	return
}

func (dao *BankDao) FindByFields(fields map[string]interface{}) ([]Bank, error) {
	var ages []Bank

	err := DB.Where(fields).Find(&ages).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return ages, nil
}

func (dao *BankDao) FindAgeByAge(age uint64) (Bank, error) {
	var sage Bank
	fmt.Println(age)
	result := DB.Where("bank = ?", 12).Find(&sage)
	fmt.Println(result)
	// DB.Raw("select * from ages where age = ?", age).Scan(&sage)
	// fmt.Println(sage)
	return sage, nil
}

func (dao *BankDao) DeleteByIds(ids []uint64) (err error) {
	if err := DB.Where(&map[string]interface{}{"id": ids}).Update("is_deleted", 1).Error; err != nil {
		return err
	}
	return nil
}

func (dao *BankDao) UpdateFields(where map[string]interface{}, updateFileds map[string]interface{}) (err error) {
	// err = DB.Where(&where).Update(updateFileds).Error
	return
}
