package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID               string         `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID           string         `gorm:"type:uuid;not null;index" json:"user_id"`
	AccountID        string         `gorm:"type:uuid;not null;index" json:"account_id"`
	CategoryID       *string        `gorm:"type:uuid;index" json:"category_id"`
	Amount           float64        `gorm:"type:decimal(15,2);not null" json:"amount"`
	Type             string         `gorm:"type:varchar(10);not null;index" json:"type"` // income, expense
	Description      string         `gorm:"type:text" json:"description"`
	TransactionDate  time.Time      `gorm:"type:date;not null;index" json:"transaction_date"`
	TransactionTime  *time.Time     `gorm:"type:time" json:"transaction_time,omitempty"`
	Location         string         `gorm:"type:varchar(255)" json:"location,omitempty"`
	Latitude         *float64       `gorm:"type:decimal(10,8)" json:"latitude,omitempty"`
	Longitude        *float64       `gorm:"type:decimal(11,8)" json:"longitude,omitempty"`
	ReceiptImageURL  string         `gorm:"type:varchar(500)" json:"receipt_image_url,omitempty"`
	IsRecurring      bool           `gorm:"default:false;index" json:"is_recurring"`
	RecurringPattern string         `gorm:"type:jsonb" json:"recurring_pattern,omitempty"`
	Tags             string         `gorm:"type:text" json:"tags,omitempty"` // simplified from text array for sqlite/postgres compatibility
	Notes            string         `gorm:"type:text" json:"notes,omitempty"`
	AIClassification string         `gorm:"type:jsonb" json:"ai_classification,omitempty"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
}
