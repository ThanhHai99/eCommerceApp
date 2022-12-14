package repository

import (
	"eCommerce/model"
	"github.com/google/uuid"
)

func GetAllWarehouse() (categories []model.Warehouse, err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Model(&model.Warehouse{}).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return
}

func GetOneWarehouse(id uuid.UUID) (res model.Warehouse, err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Model(&model.Warehouse{}).Where("id = ?", id).First(&res).Error
	if err != nil {
		return model.Warehouse{}, err
	}
	return
}

func UpdateWarehouse(warehouseId uuid.UUID, input map[string]interface{}) (res model.Warehouse, err error) {
	res.ID = warehouseId
	db, err := model.Connect()
	defer db.Close()
	if err != nil {
		return model.Warehouse{}, err
	}
	err2 := db.Model(&res).Updates(input).Error
	if err2 == nil {
		return GetOneWarehouse(res.ID)
	}
	return
}

func CreateWarehouse(warehouse *model.Warehouse) (pre *model.Warehouse, err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Model(&model.Warehouse{}).Create(warehouse).Error
	return warehouse, err
}

func DeleteWarehouse(id uuid.UUID) (err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Delete(&model.Warehouse{}, "id = ?", id).Error
	return
}
