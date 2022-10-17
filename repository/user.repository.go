package repository

import (
	"eCommerce/model"
	"github.com/jinzhu/gorm"
)

var (
	db, _ = model.Connect()
)

func IsHaveVerifyToken(token string) (tx *gorm.DB) {
	defer db.Close()
	record := db.Model(&model.User{}).First(&model.User{}, model.User{VerifyToken: token})

	return record
}

func ActiveAccount(tx *gorm.DB) {
	defer db.Close()
	 _ = tx.Update(model.User{IsActive: true}).Error
}
