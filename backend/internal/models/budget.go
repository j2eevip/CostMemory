package models

import (
	"time"

	"gorm.io/gorm"
)

type Budget struct {
	ID             string         `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID         string         `gorm:"type:uuid;not null;index" json:"user_id"`
	Name           string         `gorm:"type:varchar(100);not null" json:"name"`
	CategoryID     *string        `gorm:"type:uuid;index" json:"category_id"`
	Amount         float64        `gorm:"type:decimal(15,2);not null" json:"amount"`
	PeriodType     string         `gorm:"type:varchar(10);not null" json:"period_type"` // daily, weekly, monthly, yearly
	PeriodStart    time.Time      `gorm:"type:date;not null;index" json:"period_start"`
	PeriodEnd      time.Time      `gorm:"type:date;not null;index" json:"period_end"`
	RolloverUnused bool           `gorm:"default:false" json:"rollover_unused"`
	Notifications  string         `gorm:"type:jsonb;default:'{\"warning_threshold\": 80, \"alert_enabled\": true}'" json:"notifications"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
}
