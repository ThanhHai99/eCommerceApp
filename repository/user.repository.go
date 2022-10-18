package repository

import (
	"eCommerce/model"
	"github.com/google/uuid"
)

func CreateNew(user *model.User) error {
	db, _ := model.Connect()
	defer db.Close()
	err := db.Model(&model.User{}).Create(user).Error
	return err
}

func GenerateToken(username string) error {
	db, _ := model.Connect()
	defer db.Close()
	err := db.Where(model.User{Username: username}).Update(&model.User{}, model.User{VerifyToken: uuid.New().String()}).Error
	return err
}

func CheckExists(username string) bool {
	db, _ := model.Connect()
	defer db.Close()
	notFound := db.Model(&model.User{}).First(&model.User{}, model.User{Username: username}).RecordNotFound()
	return !notFound
}

func IsHaveVerifyToken(token string) *model.User {
	db, _ := model.Connect()
	defer db.Close()
	record := model.User{}
	//record = db.Model(&model.User{}).First(&model.User{}, model.User{VerifyToken: token})
	_ = db.First(&record, model.User{VerifyToken: token}).Error
	return &record
}

func ActiveAccount(username string) error {
	db, _ := model.Connect()
	defer db.Close()
	err := db.Where(model.User{Username: username}).Update(model.User{IsActive: true}).Error
	return err
}

func GetOneByUsername(username string) (user *model.User) {
	db, _ := model.Connect()
	defer db.Close()
	_ = db.First(&user, model.User{Username: username}).Error
	return user
}
