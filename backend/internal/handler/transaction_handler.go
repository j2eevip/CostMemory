package handler

import (
	"net/http"

	"backend/internal/service"
	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	txService service.TransactionService
}

func NewTransactionHandler(txService service.TransactionService) *TransactionHandler {
	return &TransactionHandler{txService}
}

func (h *TransactionHandler) CreateTransaction(c *gin.Context) {
	userID := c.GetString("user_id")

	var req service.CreateTransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx, err := h.txService.CreateTransaction(userID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": tx})
}

func (h *TransactionHandler) GetTransactions(c *gin.Context) {
	userID := c.GetString("user_id")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	transactions, err := h.txService.GetTransactions(userID, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transactions"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": transactions})
}

func (h *TransactionHandler) DeleteTransaction(c *gin.Context) {
	userID := c.GetString("user_id")
	txID := c.Param("id")

	if err := h.txService.DeleteTransaction(txID, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete transaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction deleted successfully"})
}
