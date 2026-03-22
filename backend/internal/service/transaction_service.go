package service

import (
	"time"
	"gorm.io/gorm"

	"backend/internal/models"
	"backend/internal/repository"
)

type TransactionService interface {
	CreateTransaction(userID string, req CreateTransactionRequest) (*models.Transaction, error)
	GetTransactions(userID string, startDate string, endDate string) ([]models.Transaction, error)
	DeleteTransaction(id string, userID string) error
}

type transactionService struct {
	txRepo      repository.TransactionRepository
	accountRepo repository.AccountRepository
}

func NewTransactionService(txRepo repository.TransactionRepository, accountRepo repository.AccountRepository) TransactionService {
	return &transactionService{txRepo, accountRepo}
}

type CreateTransactionRequest struct {
	AccountID       string  `json:"account_id" binding:"required"`
	CategoryID      *string `json:"category_id"`
	Amount          float64 `json:"amount" binding:"required"`
	Type            string  `json:"type" binding:"required,oneof=income expense"`
	Description     string  `json:"description"`
	TransactionDate string  `json:"transaction_date" binding:"required"` // YYYY-MM-DD
}

func (s *transactionService) CreateTransaction(userID string, req CreateTransactionRequest) (*models.Transaction, error) {
	date, err := time.Parse("2006-01-02", req.TransactionDate)
	if err != nil {
		return nil, err
	}

	tx := &models.Transaction{
		UserID:          userID,
		AccountID:       req.AccountID,
		CategoryID:      req.CategoryID,
		Amount:          req.Amount,
		Type:            req.Type,
		Description:     req.Description,
		TransactionDate: date,
	}

	// Calculate balance adjustment
	adjustment := req.Amount
	if req.Type == "expense" {
		adjustment = -req.Amount
	}

	// Use database transaction to guarantee atomicity
	db := s.txRepo.GetDB()
	err = db.Transaction(func(dbTx *gorm.DB) error {
		// 1. Update account balance
		if err := dbTx.Model(&models.Account{}).Where("id = ? AND user_id = ?", req.AccountID, userID).
			Update("balance", gorm.Expr("balance + ?", adjustment)).Error; err != nil {
			return err
		}

		// 2. Create transaction record
		if err := dbTx.Create(tx).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (s *transactionService) GetTransactions(userID string, startDate string, endDate string) ([]models.Transaction, error) {
	return s.txRepo.GetTransactionsByUser(userID, startDate, endDate)
}

func (s *transactionService) DeleteTransaction(id string, userID string) error {
	tx, err := s.txRepo.GetTransactionByID(id, userID)
	if err != nil {
		return err
	}

	adjustment := -tx.Amount
	if tx.Type == "expense" {
		adjustment = tx.Amount // restore balance
	}

	db := s.txRepo.GetDB()
	return db.Transaction(func(dbTx *gorm.DB) error {
		if err := dbTx.Model(&models.Account{}).Where("id = ? AND user_id = ?", tx.AccountID, userID).
			Update("balance", gorm.Expr("balance + ?", adjustment)).Error; err != nil {
			return err
		}

		if err := dbTx.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Transaction{}).Error; err != nil {
			return err
		}
		return nil
	})
}
