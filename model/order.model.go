package model

import "github.com/google/uuid"

type Order struct {
	BaseModel
	DeliveryAddress string    `gorm:"not null;" json:"delivery_address"`
	PaymentMethod   int8      `gorm:"not null;default:1"`
	Paid            bool      `gorm:"default:false" json:"paid"`
	Cost            int16     `gorm:"not null" json:"cost"`
	CreatedBy       uuid.UUID `json:"created_by"`
	ItemOrders      []ItemOrder
	Invoices        []Invoice
}

func (Order) TableName() string {
	return "order"
}

type OrderUpdateRequest struct {
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
