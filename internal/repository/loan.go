package repository

import (
	"context"

	"github.com/melisa92/billing/internal/model"
)

type LoanRepository struct {
	SampleData []model.Loan
}

func NewLoanRepository(data []model.Loan) *LoanRepository {
	return &LoanRepository{
		SampleData: data,
	}
}

func (r *LoanRepository) GetLoanList(ctx context.Context) ([]*model.Loan, error) {
	var result []*model.Loan
	for _, v := range r.SampleData {
		result = append(result, &v)
	}
	return result, nil
}

func (r *LoanRepository) GetLoanByID(ctx context.Context, loanID int) (*model.Loan, error) {
	for _, v := range r.SampleData {
		if v.ID == loanID {
			return &v, nil
		}
	}
	return nil, nil
}
