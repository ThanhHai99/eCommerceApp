package model

import "github.com/google/uuid"

type User struct {
	BaseModel
	UserName    string    `gorm:"not null;" json:"user_name"`
	Password    string    `gorm:"not null;" json:"password"`
	Name        string    `gorm:"not null;" json:"name"`
	Phone       string    `gorm:"not null;" json:"phone"`
	Address     string    `gorm:"not null;" json:"address"`
	IsActive    bool      `gorm:"not null;" json:"is_active"`
	VerifyToken string    `gorm:"not null;" json:"verify_token"`
	RoleId      uuid.UUID `gorm:"not null;" json:"role_id"`
	IsLocked    bool      `gorm:"not null;" json:"is_locked"`
}

func (User) TableName() string {
	return "user"
}

type USerUpdateRequest struct {
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
