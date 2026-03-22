package service

import (
	"backend/internal/models"
	"backend/internal/repository"
)

type AccountService interface {
	CreateAccount(userID string, req CreateAccountRequest) (*models.Account, error)
	GetAccounts(userID string) ([]models.Account, error)
	GetAccount(id string, userID string) (*models.Account, error)
	DeleteAccount(id string, userID string) error
}

type accountService struct {
	repo repository.AccountRepository
}

func NewAccountService(repo repository.AccountRepository) AccountService {
	return &accountService{repo}
}

type CreateAccountRequest struct {
	Name           string  `json:"name" binding:"required"`
	Type           string  `json:"type" binding:"required"`
	Currency       string  `json:"currency"`
	InitialBalance float64 `json:"initial_balance"`
	Color          string  `json:"color"`
	Icon           string  `json:"icon"`
}

func (s *accountService) CreateAccount(userID string, req CreateAccountRequest) (*models.Account, error) {
	if req.Currency == "" {
		req.Currency = "CNY"
	}

	account := &models.Account{
		UserID:         userID,
		Name:           req.Name,
		Type:           req.Type,
		Currency:       req.Currency,
		Balance:        req.InitialBalance, // set initial
		InitialBalance: req.InitialBalance,
		Color:          req.Color,
		Icon:           req.Icon,
		IsActive:       true,
	}

	if err := s.repo.CreateAccount(account); err != nil {
		return nil, err
	}
	return account, nil
}

func (s *accountService) GetAccounts(userID string) ([]models.Account, error) {
	return s.repo.GetAccountsByUser(userID)
}

func (s *accountService) GetAccount(id string, userID string) (*models.Account, error) {
	return s.repo.GetAccountByID(id, userID)
}

func (s *accountService) DeleteAccount(id string, userID string) error {
	return s.repo.DeleteAccount(id, userID)
}
