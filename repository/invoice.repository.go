package repository

import (
	"eCommerce/model"
	"github.com/google/uuid"
)

func GetAllInvoice() (invoices []model.Invoice, err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Model(&model.Invoice{}).Find(&invoices).Error
	if err != nil {
		return nil, err
	}
	return
}

func GetOneInvoice(id uuid.UUID) (res model.Invoice, err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Model(&model.Invoice{}).Where("id = ?", id).First(&res).Error
	if err != nil {
		return model.Invoice{}, err
	}
	return
}

func UpdateInvoice(invoiceId uuid.UUID, input map[string]interface{}) (res model.Invoice, err error) {
	res.ID = invoiceId
	db, err := model.Connect()
	defer db.Close()
	if err != nil {
		return model.Invoice{}, err
	}
	err2 := db.Model(&res).Updates(input).Error
	if err2 == nil {
		return GetOneInvoice(res.ID)
	}
	return
}

func CreateInvoice(invoice *model.Invoice) (err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Model(&model.Invoice{}).Create(invoice).Error
	return
}

func DeleteInvoice(id uuid.UUID) (err error) {
	db, _ := model.Connect()
	defer db.Close()
	err = db.Delete(&model.Invoice{}, "id = ?", id).Error
	return
}
