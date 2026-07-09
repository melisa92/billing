package interfaces

import (
	"context"

	"github.com/melisa92/billing/internal/model"
)

type LoanRepositoryInterface interface {
	GetLoanByID(ctx context.Context, loanID int) ([]*model.Loan, error)
}

type LoanScheduleRepositoryInterface interface {
	GetLoanScheduleByLoanID(ctx context.Context, loanID int) ([]*model.LoanSchedule, error)
}
