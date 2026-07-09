package model

import "time"

type Loan struct {
	ID                  int        `json:"id"`
	BorrowerID          int        `json:"borrower_id"`
	LoanAmount          float64    `json:"loan_amount"`
	InterestRatePercent float64    `json:"interest_rate_percent"`
	DurationWeeks       int        `json:"duration_weeks"`
	StartDate           time.Time  `json:"start_date"`
	EndDate             time.Time  `json:"end_date"`
	Created_at          time.Time  `json:"created_at"`
	Updated_at          *time.Time `json:"updated_at"`
}

type LoanSchedule struct {
	ID         int       `json:"id"`
	LoanID     int       `json:"loan_id"`
	WeekNumber int       `json:"week_number"`
	DueAmount  float64   `json:"due_amount"`
	DuePayDate time.Time `json:"due_pay_date"`
	PaidTime   time.Time `json:"paid_time"`
	IsPaid     bool      `json:"is_paid"`
}
