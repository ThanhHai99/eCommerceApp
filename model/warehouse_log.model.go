package model

import (
	"github.com/google/uuid"
	"time"
)

type WarehouseLog struct {
	BaseModel
	Status         string    `gorm:"not null;" json:"status"`
	Price          int16     `gorm:"not null;" json:"price"`
	ExpirationDate time.Time `gorm:"not null;" json:"expiration_date"`
	Amount         int16     `gorm:"not null;" json:"amount"`
	CreatedBy      uuid.UUID `gorm:"not null;" json:"created_by"`
	WarehouseID    uuid.UUID `gorm:"not null;" json:"warehouse"`
	ItemID         uuid.UUID `gorm:"not null;" json:"item"`
}

func (WarehouseLog) TableName() string {
	return "warehouse_log"
}

type WarehouseLogUpdateRequest struct {
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
