package repository

import (
	"eCommerce/model"
	"github.com/google/uuid"
)

func GetAllSale() (invoices []model.Sale, err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Model(&model.Sale{}).Find(&invoices).Error
	if err != nil {
		return nil, err
	}
	return
}

func GetOneSale(id uuid.UUID) (res model.Sale, err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Model(&model.Sale{}).Where("id = ?", id).First(&res).Error
	if err != nil {
		return model.Sale{}, err
	}
	return
}

func UpdateSale(saleId uuid.UUID, input map[string]interface{}) (res model.Sale, err error) {
	res.ID = saleId
	db, err := model.Connect()
	defer db.Close()
	if err != nil {
		return model.Sale{}, err
	}
	err2 := db.Model(&res).Updates(input).Error
	if err2 == nil {
		return GetOneSale(res.ID)
	}
	return
}

func CreateSale(sale *model.Sale) (pre *model.Sale, err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Model(&model.Sale{}).Create(sale).Error
	return sale, err
}

func DeleteSale(id uuid.UUID) (err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Delete(&model.Sale{}, "id = ?", id).Error
	return
}
