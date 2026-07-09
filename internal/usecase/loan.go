package LoanUsecaseInterface

import(
	"/interview/billing/internal/interfaces"
)

type LoanUsecase struct {
	loanRepo interfaces.LoanRepositoryInterface
	loanScheduleRepo interfaces.LoanScheduleRepositoryInterface
}

func NewLoanUc(
	loanRepo interfaces.LoanRepositoryInterface
	loanScheduleRepo interfaces.LoanScheduleRepositoryInterface
) *LoanUsecase{
	return &LoanUsecase {
		
	}
}