package repository

import (
	"eCommerce/model"
)

func GetAllItem() (items []model.Item, err1 error) {
	db, _ := model.Connect()
	defer db.Close()
	err1 = db.Model(&model.Item{}).Find(&items).Error
	if err1 != nil {
		return nil, err1
	}
	return items, nil
}

func CreateItem(item *model.Item) error {
	db, _ := model.Connect()
	defer db.Close()
	err1 := db.Model(&model.Item{}).Create(item).Error
	return err1
}
