package repository

import (
	"eCommerce/model"
	"github.com/google/uuid"
)

func GetAllWarehouseLog() (categories []model.WarehouseLog, err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Model(&model.WarehouseLog{}).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return
}

func GetOneWarehouseLog(id uuid.UUID) (res model.WarehouseLog, err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Model(&model.WarehouseLog{}).Where("id = ?", id).First(&res).Error
	if err != nil {
		return model.WarehouseLog{}, err
	}
	return
}

func UpdateWarehouseLog(itemId uuid.UUID, input map[string]interface{}) (res model.WarehouseLog, err error) {
	res.ID = itemId
	db, err := model.Connect()
	defer db.Close()
	if err != nil {
		return model.WarehouseLog{}, err
	}
	err2 := db.Model(&res).Updates(input).Error
	if err2 == nil {
		return GetOneWarehouseLog(res.ID)
	}
	return
}

func CreateWarehouseLog(item *model.WarehouseLog) (err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Model(&model.WarehouseLog{}).Create(item).Error
	return
}

func DeleteWarehouseLog(id uuid.UUID) (err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Delete(&model.WarehouseLog{}, "id = ?", id).Error
	return
}
