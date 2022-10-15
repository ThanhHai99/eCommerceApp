package model

import "github.com/google/uuid"

type ItemOrder struct {
	BaseModel
	Name    string    `gorm:"not null;" json:"name"`
	ItemId  uuid.UUID `gorm:"not null;" json:"item_id"`
	OrderID uuid.UUID `gorm:"not null;" json:"order_id"`
	Amount  int16     `gorm:"not null;" json:"amount"`
}

func (ItemOrder) TableName() string {
	return "item_order"
}

type ItemOrderUpdateRequest struct {
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
