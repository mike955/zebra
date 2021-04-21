package dao

import (
	"fmt"

	"gorm.io/gorm"
)

type Email struct {
	CommonTimeModel
	Id        uint64 `json:"id" gorm"primaryKey"`
	Email     string `json:"email"`
	IsDeleted uint64 `json:"is_deleted" gorm:"default:0" `
}

type EmailDao struct {
	DB *gorm.DB
}

func NewEmailDao() *EmailDao {
	return &EmailDao{
		DB: DB,
	}
}

func (dao *EmailDao) Create(data Email) (err error) {
	err = DB.Create(&data).Error
	return
}

func (dao *EmailDao) FindByFields(fields map[string]interface{}) ([]Email, error) {
	var ages []Email

	err := DB.Where(fields).Find(&ages).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return ages, nil
}

func (dao *EmailDao) FindAgeByAge(age uint64) (Email, error) {
	var sage Email
	fmt.Println(age)
	result := DB.Where("email = ?", 12).Find(&sage)
	fmt.Println(result)
	// DB.Raw("select * from ages where age = ?", age).Scan(&sage)
	// fmt.Println(sage)
	return sage, nil
}

func (dao *EmailDao) DeleteByIds(ids []uint64) (err error) {
	if err := DB.Where(&map[string]interface{}{"id": ids}).Update("is_deleted", 1).Error; err != nil {
		return err
	}
	return nil
}

func (dao *EmailDao) UpdateFields(where map[string]interface{}, updateFileds map[string]interface{}) (err error) {
	// err = DB.Where(&where).Update(updateFileds).Error
	return
}
