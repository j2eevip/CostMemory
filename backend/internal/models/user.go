package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID                   string         `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Username             string         `gorm:"type:varchar(50);not null;unique" json:"username"`
	Email                string         `gorm:"type:varchar(100);not null;unique;index" json:"email"`
	Phone                string         `gorm:"type:varchar(20);unique;index" json:"phone"`
	PasswordHash         string         `gorm:"type:varchar(255);not null" json:"-"`
	AvatarURL            string         `gorm:"type:varchar(500)" json:"avatar_url"`
	SubscriptionType     string         `gorm:"type:varchar(20);default:'free';index" json:"subscription_type"`
	SubscriptionExpiresAt *time.Time    `json:"subscription_expires_at,omitempty"`
	Settings             string         `gorm:"type:jsonb;default:'{}'" json:"settings"`
	CreatedAt            time.Time      `json:"created_at"`
	UpdatedAt            time.Time      `json:"updated_at"`
	DeletedAt            gorm.DeletedAt `gorm:"index" json:"-"`
}
