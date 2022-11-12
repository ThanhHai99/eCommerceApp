package repository

import (
	"eCommerce/model"
	"github.com/google/uuid"
)

func GetAllItemOrder() (categories []model.ItemOrder, err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Model(&model.ItemOrder{}).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return
}

func GetOneItemOrder(id uuid.UUID) (res model.ItemOrder, err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Model(&model.ItemOrder{}).Where("id = ?", id).First(&res).Error
	if err != nil {
		return model.ItemOrder{}, err
	}
	return
}

func UpdateItemOrder(itemOrderId uuid.UUID, input map[string]interface{}) (res model.ItemOrder, err error) {
	res.ID = itemOrderId
	db, err := model.Connect()
	defer db.Close()
	if err != nil {
		return model.ItemOrder{}, err
	}
	err2 := db.Model(&res).Updates(input).Error
	if err2 == nil {
		return GetOneItemOrder(res.ID)
	}
	return
}

func CreateItemOrder(itemOrder *model.ItemOrder) (pre *model.ItemOrder, err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Model(&model.ItemOrder{}).Create(itemOrder).Error
	return itemOrder, err
}

func DeleteItemOrder(id uuid.UUID) (err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Delete(&model.ItemOrder{}, "id = ?", id).Error
	return
}
