package models

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	ID             string         `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID         string         `gorm:"type:uuid;not null;index" json:"user_id"`
	Name           string         `gorm:"type:varchar(100);not null" json:"name"`
	Type           string         `gorm:"type:varchar(20);not null;index" json:"type"` // bank, wechat, alipay, cash, credit_card
	Currency       string         `gorm:"type:varchar(3);default:'CNY'" json:"currency"`
	Balance        float64        `gorm:"type:decimal(15,2);default:0.00" json:"balance"`
	InitialBalance float64        `gorm:"type:decimal(15,2);default:0.00" json:"initial_balance"`
	Color          string         `gorm:"type:varchar(7)" json:"color"`
	Icon           string         `gorm:"type:varchar(50)" json:"icon"`
	IsActive       bool           `gorm:"default:true" json:"is_active"`
	LastSyncAt     *time.Time     `json:"last_sync_at,omitempty"`
	Metadata       string         `gorm:"type:jsonb;default:'{}'" json:"metadata"` // e.g. bank account id
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
}
