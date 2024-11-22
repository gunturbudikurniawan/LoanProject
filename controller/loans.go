package controller

import (
	"GunturProject/models"
	"GunturProject/usecase"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LoanController struct {
	loanUsecase usecase.LoanUsecase
}

func NewLoanController(u usecase.LoanUsecase) *LoanController {
	return &LoanController{loanUsecase: u}
}

func (c *LoanController) CreateLoan(ctx *gin.Context) {
	var loan models.Loan
	if err := ctx.ShouldBindJSON(&loan); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format: " + err.Error()})
		return
	}

	// Log the parsed loan struct
	fmt.Printf("Parsed Loan Struct: %+v\n", loan)

	// Validate the date fields
	if loan.StartDate.IsZero() || loan.EndDate.IsZero() {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "start_date and end_date are required and must be valid"})
		return
	}

	if err := c.loanUsecase.CreateLoan(&loan); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create loan"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Loan created successfully"})
}

func (c *LoanController) GetLoanByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	loan, err := c.loanUsecase.GetLoanByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Loan not found"})
		return
	}

	ctx.JSON(http.StatusOK, loan)
}

func (c *LoanController) UpdateLoan(ctx *gin.Context) {
	var loan models.Loan
	if err := ctx.ShouldBindJSON(&loan); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.loanUsecase.UpdateLoan(&loan); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update loan"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Loan updated successfully"})
}

func (c *LoanController) DeleteLoan(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := c.loanUsecase.DeleteLoan(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete loan"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Loan deleted successfully"})
}
