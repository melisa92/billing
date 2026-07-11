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
	GetLoanScheduleByLoanID(ctx context.Context, loanID int) ([]*model.LoanSchedule, error)
	GetLoanScheduleByLoanIDAndDatePoint(ctx context.Context, loanID int, datePoint string) ([]*model.LoanSchedule, error)
	GetTotalPaidLoanByLoanID(ctx context.Context, loanID int) (float64, error)
	GetOutstandingScheduleByLoanID(ctx context.Context, loanID int) ([]*model.LoanSchedule, error)
	UpdateLoanSchedulePaidTime(ctx context.Context, loanID int, weekNumber int) error
}
