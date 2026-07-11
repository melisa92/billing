# Billing Engine

A simple billing engine implementation for a loan system written in Go.

## Features

* Generate loan repayment schedules (as dummy data)
* Calculate outstanding amounts by a certain date point
* Detect delinquent borrowers (have min 2 consecutive pending payment)
* Process loan repayments 
* Repository abstraction for storage implementation

## Assumptions / Business Logic

Since some business rules were not explicitly defined in the requirements, the following assumptions were made:

1. Borrowers must repay installments in chronological order (FIFO repayment, the oldest one must be pay first).
2. Partial payments are not supported.
3. Future installments cannot be repaid while older installments remain unpaid.
4. Outstanding amount refers only to overdue unpaid installments at the current point in time.
5. A borrower is considered delinquent if they have two or more overdue unpaid installments.
6. Each repayment must exactly match the expected installment amount.

## Project Structure

```text
cmd/
    main.go

internal/
    errors/
    interfaces/
    model/
    repository/
    usecase/
```

* `errors` contains list of error message.
* `interfaces` interface for usecase & repository
* `repository` contains repository implementations.
* `usecase` contains usecase implementation and business logic.

## Available Commands

### Generate JSON Data

```bash
go run ./cmd/main.go generate-json
```

This command generates dummy loan and repayment schedule data.

### Run Billing Simulation

```bash
go run ./cmd/main.go test
```

This command executes the predefined test scenarios and prints the results to the console.

## Current Limitation

This project intentionally uses predefined test scenarios and dummy data in order to focus on the billing logic and abstraction layer.

To test different scenarios, the reviewer may modify:

* the dummy data generation logic
* the generated JSON files
* the test cases inside the simulation function

## Example Output

```text
-- Start Running TestCase --
TestCaseOutstanding For LoanID: 1 got outstanding: 27500.000000
TestCaseIsDeliquent For LoanID: 1 the customer is bad / deliquent
TestCaseMakePayment For LoanID: 1 isSuccess
TestCaseIsDeliquent For LoanID: 1 the customer is safe (not deliquent)
TestCaseOutstanding For LoanID: 1 got outstanding: 0.000000
```
