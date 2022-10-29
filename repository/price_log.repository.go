package repository

import (
	"eCommerce/model"
	"github.com/google/uuid"
)

func GetAllPriceLog() (categories []model.PriceLog, err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Model(&model.PriceLog{}).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return
}

func GetOnePriceLog(id uuid.UUID) (res model.PriceLog, err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Model(&model.PriceLog{}).Where("id = ?", id).First(&res).Error
	if err != nil {
		return model.PriceLog{}, err
	}
	return
}

func UpdatePriceLog(itemId uuid.UUID, input map[string]interface{}) (res model.PriceLog, err error) {
	res.ID = itemId
	db, err := model.Connect()
	defer db.Close()
	if err != nil {
		return model.PriceLog{}, err
	}
	err2 := db.Model(&res).Updates(input).Error
	if err2 == nil {
		return GetOnePriceLog(res.ID)
	}
	return
}

func CreatePriceLog(item *model.PriceLog) (err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Model(&model.PriceLog{}).Create(item).Error
	return
}

func DeletePriceLog(id uuid.UUID) (err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Delete(&model.PriceLog{}, "id = ?", id).Error
	return
}
