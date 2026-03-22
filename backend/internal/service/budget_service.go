package service

import (
	"time"

	"backend/internal/models"
	"backend/internal/repository"
)

type BudgetService interface {
	CreateBudget(userID string, req CreateBudgetRequest) (*models.Budget, error)
	GetBudgets(userID string) ([]models.Budget, error)
	DeleteBudget(id string, userID string) error
}

type budgetService struct {
	repo repository.BudgetRepository
}

func NewBudgetService(repo repository.BudgetRepository) BudgetService {
	return &budgetService{repo}
}

type CreateBudgetRequest struct {
	Name        string  `json:"name" binding:"required"`
	CategoryID  *string `json:"category_id"`
	Amount      float64 `json:"amount" binding:"required"`
	PeriodType  string  `json:"period_type" binding:"required,oneof=daily weekly monthly yearly"`
	PeriodStart string  `json:"period_start" binding:"required"` // YYYY-MM-DD
	PeriodEnd   string  `json:"period_end" binding:"required"`   // YYYY-MM-DD
}

func (s *budgetService) CreateBudget(userID string, req CreateBudgetRequest) (*models.Budget, error) {
	start, err := time.Parse("2006-01-02", req.PeriodStart)
	if err != nil {
		return nil, err
	}
	end, err := time.Parse("2006-01-02", req.PeriodEnd)
	if err != nil {
		return nil, err
	}

	budget := &models.Budget{
		UserID:      userID,
		Name:        req.Name,
		CategoryID:  req.CategoryID,
		Amount:      req.Amount,
		PeriodType:  req.PeriodType,
		PeriodStart: start,
		PeriodEnd:   end,
	}

	if err := s.repo.CreateBudget(budget); err != nil {
		return nil, err
	}

	return budget, nil
}

func (s *budgetService) GetBudgets(userID string) ([]models.Budget, error) {
	return s.repo.GetBudgetsByUser(userID)
}

func (s *budgetService) DeleteBudget(id string, userID string) error {
	return s.repo.DeleteBudget(id, userID)
}
