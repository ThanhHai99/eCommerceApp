package repository

import (
	"eCommerce/model"
	"github.com/google/uuid"
)

func GetAllItemLog() (categories []model.ItemLog, err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Model(&model.ItemLog{}).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return
}

func GetOneItemLog(id uuid.UUID) (res model.ItemLog, err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Model(&model.ItemLog{}).Where("id = ?", id).First(&res).Error
	if err != nil {
		return model.ItemLog{}, err
	}
	return
}

func UpdateItemLog(itemId uuid.UUID, input map[string]interface{}) (res model.ItemLog, err error) {
	res.ID = itemId
	db, err := model.Connect()
	defer db.Close()
	if err != nil {
		return model.ItemLog{}, err
	}
	err2 := db.Model(&res).Updates(input).Error
	if err2 == nil {
		return GetOneItemLog(res.ID)
	}
	return
}

func CreateItemLog(item *model.ItemLog) (err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Model(&model.ItemLog{}).Create(item).Error
	return
}

func DeleteItemLog(id uuid.UUID) (err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Delete(&model.ItemLog{}, "id = ?", id).Error
	return
}
