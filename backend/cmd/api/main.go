package main

import (
	"fmt"
	"log"

	"backend/internal/config"
	"backend/internal/handler"
	"backend/internal/middleware"
	"backend/internal/repository"
	"backend/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	db, err := repository.InitDB(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	if cfg.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Dependencies
	userRepo := repository.NewUserRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	accountRepo := repository.NewAccountRepository(db)
	txRepo := repository.NewTransactionRepository(db)
	budgetRepo := repository.NewBudgetRepository(db)

	authService := service.NewAuthService(userRepo, cfg)
	categoryService := service.NewCategoryService(categoryRepo)
	accountService := service.NewAccountService(accountRepo)
	txService := service.NewTransactionService(txRepo, accountRepo)
	budgetService := service.NewBudgetService(budgetRepo)

	authHandler := handler.NewAuthHandler(authService)
	categoryHandler := handler.NewCategoryHandler(categoryService)
	accountHandler := handler.NewAccountHandler(accountService)
	txHandler := handler.NewTransactionHandler(txService)
	budgetHandler := handler.NewBudgetHandler(budgetService)

	r := gin.Default()

	// Routes
	v1 := r.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}

		// Protected routes
		protected := v1.Group("")
		protected.Use(middleware.AuthMiddleware(cfg))
		{
			categories := protected.Group("/categories")
			{
				categories.POST("", categoryHandler.CreateCategory)
				categories.GET("", categoryHandler.GetCategories)
				categories.DELETE("/:id", categoryHandler.DeleteCategory)
			}

			accounts := protected.Group("/accounts")
			{
				accounts.POST("", accountHandler.CreateAccount)
				accounts.GET("", accountHandler.GetAccounts)
				accounts.DELETE("/:id", accountHandler.DeleteAccount)
			}

			transactions := protected.Group("/transactions")
			{
				transactions.POST("", txHandler.CreateTransaction)
				transactions.GET("", txHandler.GetTransactions)
				transactions.DELETE("/:id", txHandler.DeleteTransaction)
			}

			budgets := protected.Group("/budgets")
			{
				budgets.POST("", budgetHandler.CreateBudget)
				budgets.GET("", budgetHandler.GetBudgets)
				budgets.DELETE("/:id", budgetHandler.DeleteBudget)
			}
		}
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	serverAddr := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Printf("Starting server on %s", serverAddr)
	if err := r.Run(serverAddr); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
