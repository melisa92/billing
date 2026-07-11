package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	iface "github.com/melisa92/billing/internal/interfaces"
	"github.com/melisa92/billing/internal/model"
	repo "github.com/melisa92/billing/internal/repository"
	usecase "github.com/melisa92/billing/internal/usecase"
)

type dummyStruct struct {
	Loan         []*model.Loan         `json:"loan"`
	LoanSchedule []*model.LoanSchedule `json:"loan_schedule"`
}

func main() {
	flagGenerateJSON := flag.Bool("generate-json", false, "Generate dummy JSON files")
	flagRunTest := flag.Bool("run-test", false, "Run billing test scenarios")

	flag.Parse()

	switch {
	case *flagGenerateJSON:
		fmt.Println("-- Start Generate JSON --")
		generateJSON()
		fmt.Println("-- Finish Generate JSON --")
	case *flagRunTest:
		fmt.Println("-- Start Running TestCase --")
		initTestCase()
	default:
		fmt.Println("Please specify a command:")
		fmt.Println("  --generate-json")
		fmt.Println("  --run-test")
	}
}

func initTestCase() {
	// Init Dummy Data
	sampleData, err := os.ReadFile("sampleData.json")
	if err != nil {
		log.Fatal(err)
	}

	if sampleData == nil {
		log.Fatal("Your sample data is empty, please generate dummy data first")
	}

	dummyData := dummyStruct{}
	json.Unmarshal(sampleData, &dummyData)

	// init Repo
	loanScheduleRepo := repo.NewLoanScheduleRepository(dummyData.LoanSchedule)

	// init usecase
	billingUsecase := usecase.NewBillingUsecase(loanScheduleRepo)

	testCases(billingUsecase)
}

func testCases(uc iface.LoanUsecaseInterface) {

	// map[int]float64 : expected the key is loan ID and float64 is the payment that you wanted to test
	testCase := map[int]float64{
		1: 27500,
	}

	ctx := context.Background()

	for loanID, paymentAmount := range testCase {
		// Test Output Outstanding
		outstanding, err := uc.GetOutstanding(ctx, loanID, time.Now().Format("2006-01-02"))
		if err != nil {
			fmt.Printf("TestCaseOutstanding For LoanID: %d got error: %s", loanID, err.Error())
		} else {
			fmt.Printf("\nTestCaseOutstanding For LoanID: %d got outstanding: %f", loanID, outstanding)
		}

		// Test Output IsDeliquent
		isDeliquent, err := uc.IsDelinquent(ctx, loanID)
		if err != nil {
			fmt.Printf("\nTestCaseIsDeliquent For LoanID: %d got error: %s", loanID, err.Error())
		} else {
			if isDeliquent {
				fmt.Printf("\nTestCaseIsDeliquent For LoanID: %d the customer is bad / deliquent", loanID)
			} else {
				fmt.Printf("\nTestCaseIsDeliquent For LoanID: %d the customer is safe (not deliquent)", loanID)
			}
		}

		// Test Output MakePayment
		isSuccess, err := uc.MakePayment(ctx, loanID, paymentAmount)
		if err != nil {
			fmt.Printf("\nTestCaseMakePayment For LoanID: %d got error: %s", loanID, err.Error())
		} else {
			if isSuccess {
				fmt.Printf("\nTestCaseMakePayment For LoanID: %d isSuccess", loanID)
			} else {
				fmt.Printf("\nTestCaseMakePayment For LoanID: %d isFailed", loanID)
			}
		}

		// Test Output IsDeliquent2
		isDeliquent2, err := uc.IsDelinquent(ctx, loanID)
		if err != nil {
			fmt.Printf("\nTestCaseIsDeliquent For LoanID: %d got error: %s", loanID, err.Error())
		} else {
			if isDeliquent2 {
				fmt.Printf("\nTestCaseIsDeliquent For LoanID: %d the customer is bad / deliquent", loanID)
			} else {
				fmt.Printf("\nTestCaseIsDeliquent For LoanID: %d the customer is safe (not deliquent)", loanID)
			}
		}

		// Test Output Outstanding2
		outstanding2, err := uc.GetOutstanding(ctx, loanID, "2029-10-10")
		if err != nil {
			fmt.Printf("TestCaseOutstanding For LoanID: %d got error: %s", loanID, err.Error())
		} else {
			fmt.Printf("\nTestCaseOutstanding For LoanID: %d got outstanding: %f", loanID, outstanding2)
		}
		fmt.Println("\n-----------------------------------------------")
	}
}

/*
*

	Will generate Dummy Data into a JSON file, but you must define the schedule paid date manually

*
*/
func generateJSON() {
	dummyJson := dummyStruct{}
	dataDummy := []*model.Loan{
		{
			ID:            1,
			BorrowerName:  "Jean",
			Amount:        50000,
			InterestRate:  10,
			WeeksLoanTerm: 10,
			StartDate:     "2026-06-01",
		},
	}

	scheduleID := 1

	for _, v := range dataDummy {
		startTime, _ := time.Parse("2006-01-02 15:04:05", fmt.Sprint(v.StartDate+" 23:59:59"))
		finishDate := startTime.AddDate(0, 0, 7*v.WeeksLoanTerm)
		v.EndDate = finishDate.Format("2006-01-02")

		var tempSchedule []*model.LoanSchedule
		amountPerWeek := v.Amount / float64(v.WeeksLoanTerm) * (100 + v.InterestRate) / 100

		for k := 1; k <= v.WeeksLoanTerm; k++ {
			tempSchedule = append(tempSchedule, &model.LoanSchedule{
				ID:         scheduleID,
				LoanID:     v.ID,
				WeekNumber: k,
				Amount:     amountPerWeek,
				DueTime:    startTime.AddDate(0, 0, k*7),
				PaidTime:   nil,
			})
			scheduleID++
		}

		dummyJson.Loan = append(dummyJson.Loan, v)
		dummyJson.LoanSchedule = append(dummyJson.LoanSchedule, tempSchedule...)
	}
	jsonData, _ := json.Marshal(dummyJson)
	os.WriteFile("sampleData.json", jsonData, 0666)
}
