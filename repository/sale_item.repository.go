package repository

import (
	"eCommerce/model"
	"github.com/google/uuid"
)

func GetAllSaleItem() (categories []model.SaleItem, err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Model(&model.SaleItem{}).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return
}

func GetOneSaleItem(id uuid.UUID) (res model.SaleItem, err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Model(&model.SaleItem{}).Where("id = ?", id).First(&res).Error
	if err != nil {
		return model.SaleItem{}, err
	}
	return
}

func UpdateSaleItem(itemId uuid.UUID, input map[string]interface{}) (res model.SaleItem, err error) {
	res.ID = itemId
	db, err := model.Connect()
	defer db.Close()
	if err != nil {
		return model.SaleItem{}, err
	}
	err2 := db.Model(&res).Updates(input).Error
	if err2 == nil {
		return GetOneSaleItem(res.ID)
	}
	return
}

func CreateSaleItem(item *model.SaleItem) (err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Model(&model.SaleItem{}).Create(item).Error
	return
}

func DeleteSaleItem(id uuid.UUID) (err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Delete(&model.SaleItem{}, "id = ?", id).Error
	return
}
