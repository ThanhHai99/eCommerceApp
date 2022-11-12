package repository

import (
	"eCommerce/model"
	"github.com/google/uuid"
)

func GetAllItem() (items []model.Item, err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Model(&model.Item{}).Find(&items).Error
	if err != nil {
		return nil, err
	}
	return
}

func GetOneItem(id uuid.UUID) (res model.Item, err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Model(&model.Item{}).Where("id = ?", id).First(&res).Error
	if err != nil {
		return model.Item{}, err
	}
	return
}

func UpdateItem(itemId uuid.UUID, input map[string]interface{}) (res model.Item, err error) {
	res.ID = itemId
	db, err := model.Connect()
	defer db.Close()
	if err != nil {
		return model.Item{}, err
	}
	err2 := db.Model(&res).Updates(input).Error
	if err2 == nil {
		return GetOneItem(res.ID)
	}
	return
}

func CreateItem(item *model.Item) (pre *model.Item, err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Model(model.Item{}).Create(item).Error
	return item, err
}

func DeleteItem(id uuid.UUID) (err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Delete(&model.Item{}, "id = ?", id).Error
	return
}
