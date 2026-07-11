package repository

import (
	"context"

	"github.com/melisa92/billing/internal/model"
)

type LoanRepository struct {
	SampleData []*model.Loan
}

func NewLoanRepository(data []*model.Loan) *LoanRepository {
	return &LoanRepository{
		SampleData: data,
	}
}

func (r *LoanRepository) GetLoanList(ctx context.Context) ([]*model.Loan, error) {
	var result []*model.Loan
	for i := range r.SampleData {
		result = append(result, r.SampleData[i])
	}
	return result, nil
}

func (r *LoanRepository) GetLoanByID(ctx context.Context, loanID int) (*model.Loan, error) {
	for k := range r.SampleData {
		if r.SampleData[k].ID == loanID {
			return r.SampleData[k], nil
		}
	}
	return nil, nil
}
