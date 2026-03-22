package repository

import (
	"backend/internal/models"
	"gorm.io/gorm"
)

type BudgetRepository interface {
	CreateBudget(budget *models.Budget) error
	GetBudgetsByUser(userID string) ([]models.Budget, error)
	GetBudgetByID(id string, userID string) (*models.Budget, error)
	DeleteBudget(id string, userID string) error
}

type budgetRepository struct {
	db *gorm.DB
}

func NewBudgetRepository(db *gorm.DB) BudgetRepository {
	return &budgetRepository{db}
}

func (r *budgetRepository) CreateBudget(budget *models.Budget) error {
	return r.db.Create(budget).Error
}

func (r *budgetRepository) GetBudgetsByUser(userID string) ([]models.Budget, error) {
	var budgets []models.Budget
	if err := r.db.Where("user_id = ?", userID).Find(&budgets).Error; err != nil {
		return nil, err
	}
	return budgets, nil
}

func (r *budgetRepository) GetBudgetByID(id string, userID string) (*models.Budget, error) {
	var budget models.Budget
	if err := r.db.Where("id = ? AND user_id = ?", id, userID).First(&budget).Error; err != nil {
		return nil, err
	}
	return &budget, nil
}

func (r *budgetRepository) DeleteBudget(id string, userID string) error {
	return r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Budget{}).Error
}
