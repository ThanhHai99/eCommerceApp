package repository

import (
	"eCommerce/model"
	"github.com/google/uuid"
)

func GetAllRole() (categories []model.Role, err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Model(&model.Role{}).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return
}

func GetOneRole(id uuid.UUID) (res model.Role, err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Model(&model.Role{}).Where("id = ?", id).First(&res).Error
	if err != nil {
		return model.Role{}, err
	}
	return
}

func UpdateRole(roleId uuid.UUID, input map[string]interface{}) (res model.Role, err error) {
	res.ID = roleId
	db, err := model.Connect()
	defer db.Close()
	if err != nil {
		return model.Role{}, err
	}
	err2 := db.Model(&res).Updates(input).Error
	if err2 == nil {
		return GetOneRole(res.ID)
	}
	return
}

func CreateRole(role *model.Role) (pre *model.Role, err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Model(&model.Role{}).Create(role).Error
	return role, err
}

func DeleteRole(id uuid.UUID) (err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Delete(&model.Role{}, "id = ?", id).Error
	return
}
