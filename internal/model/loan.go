package model

import "time"

type Loan struct {
	ID            int     `json:"id"`
	BorrowerName  string  `json:"borrower_name"`
	Amount        float64 `json:"amount"`
	InterestRate  float64 `json:"interest_rate"`
	WeeksLoanTerm int     `json:"weeks_loan_term"`
	StartDate     string  `json:"start_date"`
	EndDate       string  `json:"end_date"`
}

type LoanSchedule struct {
	ID         int        `json:"id"`
	LoanID     int        `json:"loan_id"`
	WeekNumber int        `json:"week_number"`
	Amount     float64    `json:"amount"`
	DueTime    time.Time  `json:"due_date"`
	PaidTime   *time.Time `json:"paid_date"`
}

func (m *LoanSchedule) IsPaid() bool {
	return m.PaidTime != nil
}
