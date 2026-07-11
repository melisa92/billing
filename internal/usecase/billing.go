package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	errUtil "github.com/melisa92/billing/internal/errors"
	iface "github.com/melisa92/billing/internal/interfaces"
)

type BillingUsecase struct {
	loanRepo         iface.LoanRepositoryInterface
	loanScheduleRepo iface.LoanScheduleRepositoryInterface
}

func NewBillingUsecase(loanRepo iface.LoanRepositoryInterface, loanScheduleRepo iface.LoanScheduleRepositoryInterface) *BillingUsecase {
	return &BillingUsecase{
		loanRepo:         loanRepo,
		loanScheduleRepo: loanScheduleRepo,
	}
}

func (u *BillingUsecase) GetOutstanding(ctx context.Context, loanID int, datePoint string) (float64, error) {
	loanSchedule, err := u.loanScheduleRepo.GetLoanScheduleByLoanIDAndDatePoint(ctx, loanID, datePoint)
	if err != nil {
		return 0, err
	}

	var outstanding float64
	for _, dueDet := range loanSchedule {
		if !dueDet.IsPaid() {
			outstanding += dueDet.Amount
		}
	}
	return outstanding, nil
}

func (u *BillingUsecase) IsDelinquent(ctx context.Context, loanID int) (bool, error) {
	// Implementation for checking if the loan is delinquent
	loanSchedule, err := u.loanScheduleRepo.GetLoanScheduleByLoanIDAndDatePoint(ctx, loanID, time.Now().Format("2006-01-02"))
	if err != nil {
		return false, err
	}

	isDeliquent := false
	var consecutiveLate int
	for k := len(loanSchedule) - 1; k >= 0; k-- {
		if !loanSchedule[k].IsPaid() {
			consecutiveLate++
		}
		if consecutiveLate > 1 {
			isDeliquent = true
			break
		}
	}
	return isDeliquent, nil
}

func (u *BillingUsecase) MakePayment(ctx context.Context, loanID int, amount float64) (isPaymentSuccess bool, err error) {
	// Implementation for making a payment on the loan
	loanSchedule, err := u.loanScheduleRepo.GetOutstandingScheduleByLoanID(ctx, loanID)
	if err != nil {
		return false, err
	}

	var dueAmount float64
	var weekNumber []int
	for _, dueDet := range loanSchedule {
		if dueDet.DueTime.After(time.Now().Add(24 * time.Hour)) {
			continue
		}
		weekNumber = append(weekNumber, dueDet.WeekNumber)
		dueAmount += dueDet.Amount
	}

	if dueAmount == 0 {
		return false, errors.New(errUtil.ErrNoDuePayment)
	}
	if dueAmount != amount {
		return false, fmt.Errorf(errUtil.ErrPaymentAmountNotValid, fmt.Sprintf("Rp %.2f", dueAmount))
	}

	// Update payment
	// DB should use transaction
	for i := 0; i < len(weekNumber); i++ {
		err = u.loanScheduleRepo.UpdateLoanSchedulePaidTime(ctx, loanID, weekNumber[i])
		if err != nil {
			return false, nil
		}
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
