package main

import (
	"GunturProject/controller"
	"GunturProject/models"
	"GunturProject/repository"
	"GunturProject/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:Anisa16!@tcp(127.0.0.1:3306)/loan_project_local?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Auto migrate models
	db.AutoMigrate(&models.Loan{})

	// Initialize layers
	loanRepo := repository.NewLoanRepository(db)
	loanUsecase := usecase.NewLoanUsecase(loanRepo)
	loanController := controller.NewLoanController(loanUsecase)

	// Setup router
	r := gin.Default()

	r.POST("/loans", loanController.CreateLoan)
	r.GET("/loans/:id", loanController.GetLoanByID)
	r.PUT("/loans", loanController.UpdateLoan)
	r.DELETE("/loans/:id", loanController.DeleteLoan)

	r.Run(":8080")
}
