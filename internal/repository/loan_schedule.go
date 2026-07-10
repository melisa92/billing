package repository

import (
	"context"

	"github.com/melisa92/billing/internal/model"
)

type LoanScheduleRepository struct {
	SampleData []model.LoanSchedule
}

func NewLoanScheduleRepository(data []model.LoanSchedule) *LoanScheduleRepository {
	return &LoanScheduleRepository{
		SampleData: data,
	}
}

func (r *LoanScheduleRepository) GetLoanScheduleByLoanID(ctx context.Context, loanID int) ([]*model.LoanSchedule, error) {
	if len(r.SampleData) == 0 {
		return nil, nil
	}

	result := []*model.LoanSchedule{}
	for _, val := range r.SampleData {
		if val.LoanID == loanID {
			result = append(result, &val)
		}
	}

	return result, nil
}

func (r *LoanScheduleRepository) GetTotalPaidLoanByLoanID(ctx context.Context, loanID int) (float64, error) {
	if len(r.SampleData) == 0 {
		return 0, nil
	}

	var totalPaid float64
	for _, val := range r.SampleData {
		if val.IsPaid {
			totalPaid += val.Amount
		}
		continue
	}
	return totalPaid, nil
}

// Get outstanding schedule list by today
// won't return the future outstanding schedule
func (r *LoanScheduleRepository) GetOutstandingScheduleByLoanID(ctx context.Context, loanID int) ([]*model.LoanSchedule, error) {
	if len(r.SampleData) == 0 {
		return nil, nil
	}

	var result []*model.LoanSchedule
	for _, val := range r.SampleData {
		if !val.IsPaid {
			result = append(result, &val)
		}
		continue
	}
	return result, nil
}
