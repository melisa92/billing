package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	errUtil "github.com/melisa92/billing/internal/errors"
	iface "github.com/melisa92/billing/internal/interfaces"
)

type LoanUsecase struct {
	loanRepo         iface.LoanRepositoryInterface
	loanScheduleRepo iface.LoanScheduleRepositoryInterface
}

func NewLoanUsecase(loanRepo iface.LoanRepositoryInterface, loanScheduleRepo iface.LoanScheduleRepositoryInterface) *LoanUsecase {
	return &LoanUsecase{
		loanRepo:         loanRepo,
		loanScheduleRepo: loanScheduleRepo,
	}
}

func (u *LoanUsecase) GetOutstanding(ctx context.Context, loanID int) (float64, error) {
	loanSchedule, err := u.loanScheduleRepo.GetLoanScheduleByLoanID(ctx, loanID)
	if err != nil {
		return 0, err
	}

	var outstanding float64
	for _, dueDet := range loanSchedule {
		if !dueDet.IsPaid {
			outstanding += dueDet.Amount
		}
	}
	return outstanding, nil
}

func (u *LoanUsecase) IsDelinquent(ctx context.Context, loanID int) (bool, error) {
	// Implementation for checking if the loan is delinquent
	loanSchedule, err := u.loanScheduleRepo.GetLoanScheduleByLoanID(ctx, loanID)
	if err != nil {
		return false, err
	}

	isDeliquent := false
	var consecutiveLate int
	for _, dueDet := range loanSchedule {
		if isLate(dueDet.DueTime, dueDet.PaidTime) {
			consecutiveLate++
		} else {
			consecutiveLate = 0
		}

		if consecutiveLate > 1 {
			isDeliquent = true
			break
		}
	}
	return isDeliquent, nil
}

func (u *LoanUsecase) MakePayment(ctx context.Context, loanID int, amount float64) (isPaymentSuccess bool, err error) {
	// Implementation for making a payment on the loan
	loanSchedule, err := u.loanScheduleRepo.GetOutstandingScheduleByLoanID(ctx, loanID)
	if err != nil {
		return false, err
	}

	var dueAmount float64
	for _, dueDet := range loanSchedule {
		if dueDet.DueTime.After(time.Now().Add(24 * time.Hour)) {
			continue
		}
		dueAmount += dueDet.Amount
	}

	if dueAmount == 0 {
		return false, errors.New(errUtil.ErrNoDuePayment)
	}
	if dueAmount != amount {
		return false, fmt.Errorf(errUtil.ErrPaymentAmountNotValid, fmt.Sprintf("Rp %.2f", dueAmount))
	}

	return true, nil
}

func isLate(dueTime time.Time, paidTime *time.Time) bool {
	// Criteria of Late
	// Paid, but Paid time > Due Time
	// Not Paid, Due Time > Today
	if (paidTime != nil && paidTime.After(dueTime)) ||
		(dueTime.Before(time.Now()) && paidTime == nil) {
		return true
	}
	return false
}
