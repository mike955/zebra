package dao

import (
	"gorm.io/gorm"
)

type Cellphone struct {
	CommonTimeModel
	Id        uint64 `json:"id" gorm"primaryKey"`
	Cellphone uint64 `json:"cellphone"`
	IsDeleted uint64 `json:"is_deleted" gorm:"default:0" `
}

type CellphoneDao struct {
	DB *gorm.DB
}

func NewCellphoneDao() *CellphoneDao {
	return &CellphoneDao{
		DB: DB,
	}
}

func (dao *CellphoneDao) Create(data Cellphone) (err error) {
	err = DB.Create(&data).Error
	return
}

func (dao *CellphoneDao) FindByFields(fields map[string]interface{}) ([]Cellphone, error) {
	var ages []Cellphone

	err := DB.Where(fields).Find(&ages).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return ages, nil
}

func (dao *CellphoneDao) DeleteByIds(ids []uint64) (err error) {
	if err := DB.Where(&map[string]interface{}{"id": ids}).Update("is_deleted", 1).Error; err != nil {
		return err
	}
	return nil
}

func (dao *CellphoneDao) UpdateFields(where map[string]interface{}, updateFileds map[string]interface{}) (err error) {
	// err = DB.Where(&where).Update(updateFileds).Error
	return
}
