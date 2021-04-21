package dao

import (
	"fmt"

	"gorm.io/gorm"
)

type Age struct {
	CommonTimeModel
	Id        uint64 `json:"id" gorm"primaryKey"`
	Age       uint64 `json:"age"`
	IsDeleted uint64 `json:"is_deleted" gorm:"default:0" `
}

type AgeDao struct {
	DB *gorm.DB
}

func NewAgeDao() *AgeDao {
	return &AgeDao{
		DB: DB,
	}
}

func (dao *AgeDao) Create(data Age) (err error) {
	err = DB.Create(&data).Error
	return
}

func (dao *AgeDao) FindByFields(fields map[string]interface{}) ([]Age, error) {
	var ages []Age

	err := DB.Where(fields).Find(&ages).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return ages, nil
}

func (dao *AgeDao) FindAgeByAge(age uint64) (Age, error) {
	var sage Age
	fmt.Println(age)
	result := DB.Where("age = ?", 12).Find(&sage)
	fmt.Println(result)
	// DB.Raw("select * from ages where age = ?", age).Scan(&sage)
	// fmt.Println(sage)
	return sage, nil
}

func (dao *AgeDao) DeleteByIds(ids []uint64) (err error) {
	if err := DB.Where(&map[string]interface{}{"id": ids}).Update("is_deleted", 1).Error; err != nil {
		return err
	}
	return nil
}

func (dao *AgeDao) UpdateFields(where map[string]interface{}, updateFileds map[string]interface{}) (err error) {
	// err = DB.Where(&where).Update(updateFileds).Error
	return
}
