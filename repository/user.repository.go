package repository

import (
	"eCommerce/model"
)

func CreateNew(user *model.User) error {
	db, _ := model.Connect()
	defer db.Close()
	err := db.Model(&model.User{}).Create(user).Error
	return err
}

func CheckExists(username string) bool {
	db, _ := model.Connect()
	defer db.Close()
	exist := db.Model(&model.User{}).First(&model.User{}, model.User{Username: username})
	if exist.RowsAffected == 0 {
		return false
	}
	return true
}

func IsHaveVerifyToken(token string) *model.User {
	db, _ := model.Connect()
	defer db.Close()
	record := model.User{}
	//record = db.Model(&model.User{}).First(&model.User{}, model.User{VerifyToken: token})
	db.First(&record, model.User{VerifyToken: token})
	return &record
}

func ActiveAccount(username string) error {
	db, _ := model.Connect()
	defer db.Close()
	 err := db.Where(model.User{Username: username}).Update(model.User{IsActive: true}).Error
	 return err
}
