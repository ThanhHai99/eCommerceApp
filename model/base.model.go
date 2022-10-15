package model

import (
	"github.com/google/uuid"
	"time"
)

type BaseModel struct {
	// https://gorm.io/docs/models.html
	ID        uuid.UUID  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type BaseResponseModel struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Total   int16       `json:"total"`
	Data    interface{} `json:"data"`
}
