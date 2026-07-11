package interfaces

import (
	"context"

	"github.com/melisa92/billing/internal/model"
)

type LoanRepositoryInterface interface {
	GetLoanList(ctx context.Context) ([]*model.Loan, error)
	GetLoanByID(ctx context.Context, loanID int) (*model.Loan, error)
}

type LoanScheduleRepositoryInterface interface {
	GetLoanScheduleByLoanIDAndDatePoint(ctx context.Context, loanID int, datePoint string) ([]*model.LoanSchedule, error)
	GetOutstandingScheduleByLoanIDAndDatePoint(ctx context.Context, loanID int, datePoint string) ([]*model.LoanSchedule, error)
	UpdateLoanSchedulePaidTime(ctx context.Context, loanID int, weekNumber int) error
}
