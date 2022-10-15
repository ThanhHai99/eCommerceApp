package model

import "github.com/google/uuid"

type PriceLog struct {
	BaseModel
	Name      string    `gorm:"not null;" json:"name"`
	ItemId    uuid.UUID `gorm:"not null;" json:"item_id"`
	Price     int16     `gorm:"not null;" json:"price"`
	CreatedBy uuid.UUID `gorm:"not null;" json:"created_by"`
}

func (PriceLog) TableName() string {
	return "price_log"
}

type PriceLogUpdateRequest struct {
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
