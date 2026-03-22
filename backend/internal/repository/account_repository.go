package repository

import (
	"backend/internal/models"
	"gorm.io/gorm"
)

type AccountRepository interface {
	CreateAccount(account *models.Account) error
	GetAccountsByUser(userID string) ([]models.Account, error)
	GetAccountByID(id string, userID string) (*models.Account, error)
	UpdateBalance(id string, amount float64) error
	DeleteAccount(id string, userID string) error
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &accountRepository{db}
}

func (r *accountRepository) CreateAccount(account *models.Account) error {
	return r.db.Create(account).Error
}

func (r *accountRepository) GetAccountsByUser(userID string) ([]models.Account, error) {
	var accounts []models.Account
	if err := r.db.Where("user_id = ?", userID).Find(&accounts).Error; err != nil {
		return nil, err
	}
	return accounts, nil
}

func (r *accountRepository) GetAccountByID(id string, userID string) (*models.Account, error) {
	var account models.Account
	if err := r.db.Where("id = ? AND user_id = ?", id, userID).First(&account).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *accountRepository) UpdateBalance(id string, amount float64) error {
	// Increment/Decrement balance.
	// amount should be positive for income, negative for expense.
	return r.db.Model(&models.Account{}).Where("id = ?", id).
		Update("balance", gorm.Expr("balance + ?", amount)).Error
}

func (r *accountRepository) DeleteAccount(id string, userID string) error {
	return r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Account{}).Error
}
