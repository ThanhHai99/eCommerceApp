package model

import "github.com/google/uuid"

type SaleItem struct {
	BaseModel
	ItemID uuid.UUID `gorm:"not null;" json:"item_id"`
	SaleID uuid.UUID `gorm:"not null;" json:"sale_id"`
	Amount int16     `gorm:"not null;" json:"amount"`
}

func (SaleItem) TableName() string {
	return "sale_item"
}

type SaleItemUpdateRequest struct {
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
