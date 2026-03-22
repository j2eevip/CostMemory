package repository

import (
	"backend/internal/models"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	CreateCategory(category *models.Category) error
	GetCategoriesByUser(userID string) ([]models.Category, error)
	GetCategoryByID(id string) (*models.Category, error)
	DeleteCategory(id string, userID string) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) CreateCategory(category *models.Category) error {
	return r.db.Create(category).Error
}

func (r *categoryRepository) GetCategoriesByUser(userID string) ([]models.Category, error) {
	var categories []models.Category
	if err := r.db.Where("user_id = ? OR is_system = ?", userID, true).Order("type ASC, sort_order ASC").Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *categoryRepository) GetCategoryByID(id string) (*models.Category, error) {
	var category models.Category
	if err := r.db.First(&category, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) DeleteCategory(id string, userID string) error {
	return r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Category{}).Error
}
