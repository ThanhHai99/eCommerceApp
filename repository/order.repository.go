package repository

import (
	"eCommerce/model"
	"github.com/google/uuid"
)

func GetAllOrder() (invoices []model.Order, err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Model(&model.Order{}).Find(&invoices).Error
	if err != nil {
		return nil, err
	}
	return
}

func GetOneOrder(id uuid.UUID) (res model.Order, err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Model(&model.Order{}).Where("id = ?", id).First(&res).Error
	if err != nil {
		return model.Order{}, err
	}
	return
}

func UpdateOrder(invoiceId uuid.UUID, input map[string]interface{}) (res model.Order, err error) {
	res.ID = invoiceId
	db, err := model.Connect()
	defer db.Close()
	if err != nil {
		return model.Order{}, err
	}
	err2 := db.Model(&res).Updates(input).Error
	if err2 == nil {
		return GetOneOrder(res.ID)
	}
	return
}

func CreateOrder(invoice *model.Order) (err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Model(&model.Order{}).Create(invoice).Error
	return
}

func DeleteOrder(id uuid.UUID) (err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Delete(&model.Order{}, "id = ?", id).Error
	return
}
