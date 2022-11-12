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

func UpdateCategory(categoryId uuid.UUID, input map[string]interface{}) (res model.Category, err error) {
	res.ID = categoryId
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

func CreateCategory(category *model.Category) (pre *model.Category, err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Model(&model.Category{}).Create(category).Error
	return category, err
}

func DeleteCategory(id uuid.UUID) (err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Delete(&model.Category{}, "id = ?", id).Error
	return
}
