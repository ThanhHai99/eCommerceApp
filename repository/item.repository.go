package repository

import (
	"eCommerce/model"
	"github.com/google/uuid"
)

func GetAllItem() (items []model.Item, err1 error) {
	db, _ := model.Connect()
	defer db.Close()
	err1 = db.Model(&model.Item{}).Find(&items).Error
	if err1 != nil {
		return nil, err1
	}
	return
}

func GetOneItem(id uuid.UUID) (res model.Item, err1 error) {
	db, _ := model.Connect()
	defer db.Close()
	err1 = db.Model(&model.Item{}).Where("id = ?", id).First(&res).Error
	if err1 != nil {
		return model.Item{}, err1
	}
	return
}

func CreateItem(item *model.Item) error {
	db, _ := model.Connect()
	defer db.Close()
	err1 := db.Model(&model.Item{}).Create(item).Error
	return err1
}
