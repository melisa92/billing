package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/melisa92/billing/internal/model"
)

type LoanScheduleRepository struct {
	SampleData []*model.LoanSchedule
}

func NewLoanScheduleRepository(data []*model.LoanSchedule) *LoanScheduleRepository {
	return &LoanScheduleRepository{
		SampleData: data,
	}
}

func (r *LoanScheduleRepository) GetLoanScheduleByLoanIDAndDatePoint(ctx context.Context, loanID int, datePoint string) ([]*model.LoanSchedule, error) {
	if len(r.SampleData) == 0 {
		return nil, nil
	}

	result := []*model.LoanSchedule{}
	datePointTime, _ := time.Parse("2006-01-02 15:04:05", fmt.Sprint(datePoint, " 23:59:59"))
	for _, val := range r.SampleData {
		if val.LoanID == loanID && (val.DueTime.Before(datePointTime) || val.DueTime.Equal(datePointTime)) {
			result = append(result, val)
		}
	}
	return result, nil
}

// Get outstanding schedule list by today
// won't return the future outstanding schedule
func (r *LoanScheduleRepository) GetOutstandingScheduleByLoanIDAndDatePoint(ctx context.Context, loanID int, datePoint string) ([]*model.LoanSchedule, error) {
	if len(r.SampleData) == 0 {
		return nil, nil
	}

	result := []*model.LoanSchedule{}
	datePointTime, _ := time.Parse("2006-01-02 15:04:05", fmt.Sprint(datePoint, " 23:59:59"))
	for _, val := range r.SampleData {
		if val.LoanID == loanID && !val.IsPaid() && (val.DueTime.Before(datePointTime) || val.DueTime.Equal(datePointTime)) {
			result = append(result, val)
		}
	}
	return result, nil
}

func (r *LoanScheduleRepository) UpdateLoanSchedulePaidTime(ctx context.Context, loanID int, weekNumber int) error {
	if len(r.SampleData) == 0 {
		return errors.New("no data to be updated")
	}

	paidTime := time.Now().UTC()
	for _, val := range r.SampleData {
		if val.LoanID == loanID && val.WeekNumber == weekNumber {
			val.PaidTime = &paidTime
		}
	}

	return nil
}
