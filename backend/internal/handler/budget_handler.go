package handler

import (
	"net/http"

	"backend/internal/service"
	"github.com/gin-gonic/gin"
)

type BudgetHandler struct {
	budgetService service.BudgetService
}

func NewBudgetHandler(budgetService service.BudgetService) *BudgetHandler {
	return &BudgetHandler{budgetService}
}

func (h *BudgetHandler) CreateBudget(c *gin.Context) {
	userID := c.GetString("user_id")

	var req service.CreateBudgetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	budget, err := h.budgetService.CreateBudget(userID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create budget"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": budget})
}

func (h *BudgetHandler) GetBudgets(c *gin.Context) {
	userID := c.GetString("user_id")

	budgets, err := h.budgetService.GetBudgets(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch budgets"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": budgets})
}

func (h *BudgetHandler) DeleteBudget(c *gin.Context) {
	userID := c.GetString("user_id")
	budgetID := c.Param("id")

	if err := h.budgetService.DeleteBudget(budgetID, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete budget"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Budget deleted successfully"})
}
