package interfaces

import "context"

type LoanUsecaseInterface interface {
	//	GetOutstanding : This returns the current outstanding on a loan, 0 if no outstanding(or closed),
	GetOutstanding(ctx context.Context, loanID int, datePoint string) (float64, error)

	// IsDelinquent : If there are more than 2 weeks of Non payment of the loan amount
	IsDelinquent(ctx context.Context, loanID int) (bool, error)

	// MakePayment: Make a payment of certain amount on the loan
	MakePayment(ctx context.Context, loanID int, amount float64) (isPaymentSuccess bool, err error)
}
