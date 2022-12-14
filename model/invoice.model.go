package model

import "github.com/google/uuid"

type Invoice struct {
	BaseModel
	Name      string    `gorm:"not null;" json:"name"`
	Phone     string    `gorm:"not null;" json:"phone"`
	Cost      int16     `gorm:"not null;" json:"cost"`
	OrderID   uuid.UUID `gorm:"not null;" json:"order_id"`
	CreatedBy uuid.UUID `gorm:"not null;" json:"created_by"`
}

func (Invoice) TableName() string {
	return "invoice"
}

type InvoiceUpdateRequest struct {
	//Id          *string          `json:"id"`
	//Name        *string          `json:"name" `
	//Key         *string          `json:"key"`
	//Value       *string          `json:"value"`
	//JsonValue   *postgres.Jsonb  `json:"json_value"`
	//Type        *string          `json:"type"`
	//Url         *pq.StringArray  `json:"url"`
	//IsActive    *bool            `json:"is_active"`
	//PlatformKey *string          `json:"platform_key"`
	//UpdaterID   *uuid.UUID       `json:"updater_id"`
	//UrlEn       *pq.StringArray  `json:"url_en"`
	//OrderBy     *int             `json:"order_by"`
	//Advanced    *json.RawMessage `json:"advanced"`
}
