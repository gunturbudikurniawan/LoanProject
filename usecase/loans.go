package usecase

import (
	"GunturProject/models"
	"GunturProject/repository"
)

type LoanUsecase interface {
	CreateLoan(loan *models.Loan) error
	GetLoanByID(id uint) (*models.Loan, error)
	UpdateLoan(loan *models.Loan) error
	DeleteLoan(id uint) error
}

type loanUsecase struct {
	repo repository.LoanRepository
}

func NewLoanUsecase(repo repository.LoanRepository) LoanUsecase {
	return &loanUsecase{repo}
}

func (u *loanUsecase) CreateLoan(loan *models.Loan) error {

	return u.repo.CreateLoan(loan)
}

func (u *loanUsecase) GetLoanByID(id uint) (*models.Loan, error) {
	return u.repo.GetLoanByID(id)
}

func (u *loanUsecase) UpdateLoan(loan *models.Loan) error {
	return u.repo.UpdateLoan(loan)
}

func (u *loanUsecase) DeleteLoan(id uint) error {
	return u.repo.DeleteLoan(id)
}
