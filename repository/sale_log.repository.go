package repository

import (
	"eCommerce/model"
	"github.com/google/uuid"
)

func GetAllSaleLog() (categories []model.SaleLog, err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Model(&model.SaleLog{}).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return
}

func GetOneSaleLog(id uuid.UUID) (res model.SaleLog, err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Model(&model.SaleLog{}).Where("id = ?", id).First(&res).Error
	if err != nil {
		return model.SaleLog{}, err
	}
	return
}

func UpdateSaleLog(itemId uuid.UUID, input map[string]interface{}) (res model.SaleLog, err error) {
	res.ID = itemId
	db, err := model.Connect()
	defer db.Close()
	if err != nil {
		return model.SaleLog{}, err
	}
	err2 := db.Model(&res).Updates(input).Error
	if err2 == nil {
		return GetOneSaleLog(res.ID)
	}
	return
}

func CreateSaleLog(item *model.SaleLog) (err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Model(&model.SaleLog{}).Create(item).Error
	return
}

func DeleteSaleLog(id uuid.UUID) (err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Delete(&model.SaleLog{}, "id = ?", id).Error
	return
}
