package handler

import (
	"net/http"

	"backend/internal/service"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	categoryService service.CategoryService
}

func NewCategoryHandler(categoryService service.CategoryService) *CategoryHandler {
	return &CategoryHandler{categoryService}
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	userID := c.GetString("user_id")
	
	var req service.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := h.categoryService.CreateCategory(userID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": category})
}

func (h *CategoryHandler) GetCategories(c *gin.Context) {
	userID := c.GetString("user_id")

	categories, err := h.categoryService.GetCategories(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch categories"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": categories})
}

func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	userID := c.GetString("user_id")
	categoryID := c.Param("id")

	if err := h.categoryService.DeleteCategory(categoryID, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}
