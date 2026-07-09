package interfaces

import (
	"context"

	"github.com/melisa92/billing/internal/model"
)

type LoanUsecaseInterface interface {
	GetOutstanding(ctx context.Context) ([]*model.Loan, error)
	GetOutstandingByLoanID(ctx context.Context, loanID int) (*model.Loan, error)

	IsDeliquent(ctx context.Context) error
	MakePayment(ctx context.Context) error
}
