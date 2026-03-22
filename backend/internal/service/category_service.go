package service

import (
	"backend/internal/models"
	"backend/internal/repository"
)

type CategoryService interface {
	CreateCategory(userID string, req CreateCategoryRequest) (*models.Category, error)
	GetCategories(userID string) ([]models.Category, error)
	DeleteCategory(id string, userID string) error
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{repo}
}

type CreateCategoryRequest struct {
	Name     string `json:"name" binding:"required"`
	Icon     string `json:"icon" binding:"required"`
	Color    string `json:"color" binding:"required"`
	Type     string `json:"type" binding:"required,oneof=income expense"`
	ParentID *string `json:"parent_id"`
}

func (s *categoryService) CreateCategory(userID string, req CreateCategoryRequest) (*models.Category, error) {
	category := &models.Category{
		UserID:   userID,
		Name:     req.Name,
		Icon:     req.Icon,
		Color:    req.Color,
		Type:     req.Type,
		ParentID: req.ParentID,
		IsSystem: false,
	}

	if err := s.repo.CreateCategory(category); err != nil {
		return nil, err
	}

	return category, nil
}

func (s *categoryService) GetCategories(userID string) ([]models.Category, error) {
	return s.repo.GetCategoriesByUser(userID)
}

func (s *categoryService) DeleteCategory(id string, userID string) error {
	return s.repo.DeleteCategory(id, userID)
}
