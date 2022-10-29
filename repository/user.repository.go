package repository

import (
	"eCommerce/model"
	"github.com/google/uuid"
)

func CreateNew(user *model.User) error {
	db, _ := model.Connect()
	defer db.Close()
	err1 := db.Model(&model.User{}).Create(user).Error
	return err1
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

func GetOneByUsername(username string) *model.User {
	record := model.User{}
	db, _ := model.Connect()
	defer db.Close()
	_ = db.First(&record, model.User{Username: username}).Error
	return &record
}

func GetAllUser() (invoices []model.User, err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Model(&model.User{}).Find(&invoices).Error
	if err != nil {
		return nil, err
	}
	return
}

func GetOneUser(id uuid.UUID) (res model.User, err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Model(&model.User{}).Where("id = ?", id).First(&res).Error
	if err != nil {
		return model.User{}, err
	}
	return
}

func UpdateUser(invoiceId uuid.UUID, input map[string]interface{}) (res model.User, err error) {
	res.ID = invoiceId
	db, err := model.Connect()
	defer db.Close()
	if err != nil {
		return model.User{}, err
	}
	err2 := db.Model(&res).Updates(input).Error
	if err2 == nil {
		return GetOneUser(res.ID)
	}
	return
}

func CreateUser(invoice *model.User) (err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Model(&model.User{}).Create(invoice).Error
	return
}

func DeleteUser(id uuid.UUID) (err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Delete(&model.User{}, "id = ?", id).Error
	return
}

