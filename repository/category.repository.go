package repository

import (
	"eCommerce/model"
	"github.com/google/uuid"
)

func GetAllCategory() (categories []model.Category, err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Model(&model.Category{}).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return
}

func GetOneCategory(id uuid.UUID) (res model.Category, err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Model(&model.Category{}).Where("id = ?", id).First(&res).Error
	if err != nil {
		return model.Category{}, err
	}
	return
}

func UpdateCategory(itemId uuid.UUID, input map[string]interface{}) (res model.Category, err error) {
	res.ID = itemId
	db, err := model.Connect()
	defer db.Close()
	if err != nil {
		return model.Category{}, err
	}
	err2 := db.Model(&res).Updates(input).Error
	if err2 == nil {
		return GetOneCategory(res.ID)
	}
	return
}

func CreateCategory(item *model.Category) (err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Model(&model.Category{}).Create(item).Error
	return
}

func DeleteCategory(id uuid.UUID) (err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Delete(&model.Category{}, "id = ?", id).Error
	return
}
