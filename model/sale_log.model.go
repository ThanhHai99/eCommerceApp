package model

import (
	"github.com/google/uuid"
	"time"
)

type SaleLog struct {
	BaseModel
	Name      string    `gorm:"not null;" json:"name"`
	SaleItem  uuid.UUID `gorm:"not null;" json:"sale_item"`
	StartDate time.Time `gorm:"not null;" json:"start_date"`
	EndDate   time.Time `gorm:"not null;" json:"end_date"`
	Amount    int16     `gorm:"not null;" json:"amount"`
	Discount  int16     `gorm:"not null;" json:"discount"`
	IsApplied bool      `gorm:"not null;" json:"is_applied"`
	Code      string    `gorm:"not null;" json:"code"`
	CreatedBy uuid.UUID `gorm:"not null;" json:"created_by"`
	SaleID    uuid.UUID `gorm:"not null;" json:"sale_id"`
}

func (SaleLog) TableName() string {
	return "sale_log"
}

type SaleLogUpdateRequest struct {
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
