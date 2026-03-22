package repository

import (
	"backend/internal/models"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(tx *models.Transaction) error
	GetTransactionsByUser(userID string, startDate string, endDate string) ([]models.Transaction, error)
	GetTransactionByID(id string, userID string) (*models.Transaction, error)
	DeleteTransaction(id string, userID string) error
	GetDB() *gorm.DB // Expose DB for transaction wrapping in logic
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) GetDB() *gorm.DB {
	return r.db
}

func (r *transactionRepository) CreateTransaction(tx *models.Transaction) error {
	return r.db.Create(tx).Error
}

func (r *transactionRepository) GetTransactionsByUser(userID string, startDate string, endDate string) ([]models.Transaction, error) {
	var transactions []models.Transaction
	query := r.db.Where("user_id = ?", userID)
	
	if startDate != "" {
		query = query.Where("transaction_date >= ?", startDate)
	}
	if endDate != "" {
		query = query.Where("transaction_date <= ?", endDate)
	}

	if err := query.Order("transaction_date DESC, created_at DESC").Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *transactionRepository) GetTransactionByID(id string, userID string) (*models.Transaction, error) {
	var tx models.Transaction
	if err := r.db.Where("id = ? AND user_id = ?", id, userID).First(&tx).Error; err != nil {
		return nil, err
	}
	return &tx, nil
}

func (r *transactionRepository) DeleteTransaction(id string, userID string) error {
	return r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Transaction{}).Error
}
