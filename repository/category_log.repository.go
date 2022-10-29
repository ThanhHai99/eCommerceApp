package repository

import (
	"eCommerce/model"
	"github.com/google/uuid"
)

func GetAllCategoryLog() (categories []model.CategoryLog, err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Model(&model.CategoryLog{}).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return
}

func GetOneCategoryLog(id uuid.UUID) (res model.CategoryLog, err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Model(&model.CategoryLog{}).Where("id = ?", id).First(&res).Error
	if err != nil {
		return model.CategoryLog{}, err
	}
	return
}

func UpdateCategoryLog(itemId uuid.UUID, input map[string]interface{}) (res model.CategoryLog, err error) {
	res.ID = itemId
	db, err := model.Connect()
	defer db.Close()
	if err != nil {
		return model.CategoryLog{}, err
	}
	err2 := db.Model(&res).Updates(input).Error
	if err2 == nil {
		return GetOneCategoryLog(res.ID)
	}
	return
}

func CreateCategoryLog(item *model.CategoryLog) (err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Model(&model.CategoryLog{}).Create(item).Error
	return
}

func DeleteCategoryLog(id uuid.UUID) (err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Delete(&model.CategoryLog{}, "id = ?", id).Error
	return
}
