package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        string         `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID    string         `gorm:"type:uuid;not null;index" json:"user_id"`
	Name      string         `gorm:"type:varchar(100);not null" json:"name"`
	Icon      string         `gorm:"type:varchar(50)" json:"icon"`
	Color     string         `gorm:"type:varchar(7)" json:"color"`
	Type      string         `gorm:"type:varchar(10);not null;index" json:"type"` // income, expense
	ParentID  *string        `gorm:"type:uuid;index" json:"parent_id,omitempty"`
	IsSystem  bool           `gorm:"default:false" json:"is_system"`
	SortOrder int            `gorm:"default:0" json:"sort_order"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
