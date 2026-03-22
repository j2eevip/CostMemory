package service

import (
	"errors"

	"backend/internal/config"
	"backend/internal/models"
	"backend/internal/repository"
	"backend/pkg/utils"
)

type AuthService interface {
	Register(req RegisterRequest) error
	Login(req LoginRequest) (string, error)
}

type authService struct {
	repo repository.UserRepository
	cfg  *config.Config
}

func NewAuthService(repo repository.UserRepository, cfg *config.Config) AuthService {
	return &authService{repo, cfg}
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Phone    string `json:"phone"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (s *authService) Register(req RegisterRequest) error {
	// check if user exists
	if _, err := s.repo.GetUserByEmail(req.Email); err == nil {
		return errors.New("email already registered")
	}

	hashed, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}

	user := &models.User{
		Username:     req.Username,
		Email:        req.Email,
		Phone:        req.Phone,
		PasswordHash: hashed,
	}

	return s.repo.CreateUser(user)
}

func (s *authService) Login(req LoginRequest) (string, error) {
	user, err := s.repo.GetUserByEmail(req.Email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	if !utils.CheckPasswordHash(req.Password, user.PasswordHash) {
		return "", errors.New("invalid email or password")
	}

	token, err := utils.GenerateToken(user.ID, s.cfg.JWT.Secret, s.cfg.JWT.ExpirationHours)
	if err != nil {
		return "", err
	}

	return token, nil
}
