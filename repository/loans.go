package repository

import (
	"GunturProject/models"

	"gorm.io/gorm"
)

type LoanRepository interface {
	CreateLoan(loan *models.Loan) error
	GetLoanByID(id uint) (*models.Loan, error)
	UpdateLoan(loan *models.Loan) error
	DeleteLoan(id uint) error
}

type loanRepository struct {
	db *gorm.DB
}

func NewLoanRepository(db *gorm.DB) LoanRepository {
	return &loanRepository{db}
}

func (r *loanRepository) CreateLoan(loan *models.Loan) error {
	return r.db.Create(loan).Error
}

func (r *loanRepository) GetLoanByID(id uint) (*models.Loan, error) {
	var loan models.Loan
	if err := r.db.First(&loan, id).Error; err != nil {
		return nil, err
	}
	return &loan, nil
}

func (r *loanRepository) UpdateLoan(loan *models.Loan) error {
	return r.db.Save(loan).Error
}

func (r *loanRepository) DeleteLoan(id uint) error {
	return r.db.Delete(&models.Loan{}, id).Error
}
