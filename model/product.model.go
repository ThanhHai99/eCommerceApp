package model

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lib/pq"
)

type Product struct {
	BaseModel
	Name string `gorm:"not null;" json:"name"`
	//Key         string         `gorm:"not null;" json:"key" validate:"required"`
	//JsonValue   postgres.Jsonb `gorm:"not null" json:"json_value"`
	//Url         pq.StringArray `gorm:"type:varchar(255)[];column:url;null;" json:"url"`
	//CreatorID   uuid.UUID      `gorm:"not null;type:uuid;" json:"creator_id" validate:"required"`
	//UrlEn       pq.StringArray `gorm:"type:varchar(255)[];column:url_en;null;" json:"url_en"`
	//UpdaterID   uuid.UUID      `gorm:"not null;type:uuid;" json:"updater_id"`
}

func (Product) TableName() string {
	return "product"
}

type ProductUpdateRequest struct {
	Id          *string          `json:"id"`
	Name        *string          `json:"name" `
	Key         *string          `json:"key"`
	Value       *string          `json:"value"`
	JsonValue   *postgres.Jsonb  `json:"json_value"`
	Type        *string          `json:"type"`
	Url         *pq.StringArray  `json:"url"`
	IsActive    *bool            `json:"is_active"`
	PlatformKey *string          `json:"platform_key"`
	UpdaterID   *uuid.UUID       `json:"updater_id"`
	UrlEn       *pq.StringArray  `json:"url_en"`
	OrderBy     *int             `json:"order_by"`
	Advanced    *json.RawMessage `json:"advanced"`
}
