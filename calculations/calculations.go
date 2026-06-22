package calculations

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Loan struct {
	Principal float64
	Term      int
	Rate      float64
}

// -----------Get Loan Information from User------------
func GetInformation(info *Loan) {
	reader := bufio.NewReader(os.Stdin)
	for {

		fmt.Printf("\nPlease provide loan information below")
		fmt.Println()

		fmt.Print("\nLoan Amount: ")
		amount, err := reader.ReadString('\n')
		if err != nil || amount == "" {
			log.Println("Nothing entered. Please try again ", err)
		}
		amount = strings.TrimSpace(amount)
		principal, err := strconv.ParseFloat(amount, 64)
		if err != nil {
			log.Print("Invalid input. Please try again. (i.e. 45000) ", err)
			continue
		}
		info.Principal = principal

		fmt.Print("Interest Rate (APR): ")
		interestRate, err := reader.ReadString('\n')
		if err != nil || interestRate == "" {
			log.Println("Nothing entered. Please try again ", err)
		}
		interestRate = strings.TrimSpace(interestRate)
		rate, err := strconv.ParseFloat(interestRate, 64)
		if err != nil {
			log.Print("Invalid input. Please try again. (i.e. 5 for 5%)", err)
			continue
		}
		rateDecimal := rate / 100
		info.Rate = rateDecimal

		fmt.Print("Loan Term: ")
		loanTerm, err := reader.ReadString('\n')
		if err != nil || loanTerm == "" {
			log.Println("Nothing entered. Please try again ", err)
		}
		loanTerm = strings.TrimSpace(loanTerm)
		term, err := strconv.Atoi(loanTerm)
		if err != nil {
			log.Print("Invalid input. Please try again. (i.e. 5 for 5 years)", err)
			continue
		}
		info.Term = term

		fmt.Printf("\nYou entered the following information:")
		fmt.Println()
		info.PrintInfo()

		fmt.Print("\nIs this information correct? [Yes / No]: ")
		confirm, err := reader.ReadString('\n')
		if err != nil || confirm == "" {
			log.Println("Nothing entered. Please try again ", err)
		}
		confirm = strings.TrimSpace(confirm)

		if confirm == "Yes" || confirm == "yes" || confirm == "Y" || confirm == "y" {
			fmt.Println()
			break
		} else if confirm == "No" || confirm == "no" || confirm == "N" || confirm == "n" {
			continue
		} else {
			fmt.Print("\nInvalid entry. Please select either Yes or No.")
		}
	}
}

// ------------------Simple Interest Formula---------------------
func (l Loan) SimpleInterest() float64 {
	var t = float64(l.Term)
	simpleInterest := l.Principal * l.Rate * t
	return simpleInterest
}

// ----------------Compound Interest Formula----------------------
func (l Loan) CompoundInterest(n float64) float64 {
	t := float64(l.Term)
	compoundAmount := l.Principal * math.Pow(1+(l.Rate/n), n*t)
	return compoundAmount - l.Principal
}

// Formula for interest compounded daily
func (l Loan) DailyCompoundInterest() float64 {
	return l.CompoundInterest(365)
}

// Formula for interest compounded weekly
func (l Loan) WeeklyCompoundInterest() float64 {
	return l.CompoundInterest(52)
}

// Formula for interest compounded monthly
func (l Loan) MonthlyCompoundInterest() float64 {
	return l.CompoundInterest(12)
}

// Formula for interest compounded annually
func (l Loan) AnnualCompoundInterest() float64 {
	return l.CompoundInterest(1)
}

// -------------Daily Accrued Interest Formula--------------------
func (l Loan) DailyAccruedInterest() float64 {
	dailyAccruedInterest := (l.DailyInterestRate() * l.Principal)
	return dailyAccruedInterest
}

// --------Interest Rates-------------------------
// Formula for daily interest rate based on APR
func (l Loan) DailyInterestRate() float64 {
	dailyInterestRate := l.Rate / 365
	return dailyInterestRate
}

// Formula for monthly interest rate based on APR
func (l Loan) MonthlyInterestRate() float64 {
	return l.Rate / 12
}

// --------Number of Payments---------------------
func (l Loan) TotalPayments() int {
	return l.Term * 12
}

// -------Monthly Payment Amount----------------
func (l Loan) MonthlyPayment() float64 {
	r := l.MonthlyInterestRate()
	n := float64(l.TotalPayments())
	p := l.Principal

	// If no interest is charged, divide the Loan Amount by the number of payments to get the monthly payment amount
	if r == 0 {
		return p / n
		// If no payments are required, return 0 since there is no loan balance
	} else if n == 0 {
		return 0
	}
	// Formula to calculate the monthly payment amount (if r,n and p have a value greater than 0)
	return p * r * math.Pow(1+r, n) / (math.Pow(1+r, n) - 1)
}

// Formula for calculating the total cost of the loan. The total amount the user will pay back.
func (l Loan) TotalCost() float64 {
	return l.MonthlyPayment() * float64(l.TotalPayments())
}

// Formula for calculating the total interest paid.
func (l Loan) TotalInterestAmount() float64 {
	return l.TotalCost() - l.Principal
}

// Display the loan information the user entered .
func (l Loan) PrintInfo() {
	fmt.Printf("Loan Amount: $%.2f\n", l.Principal)
	fmt.Printf("Interest Rate: %.2f%%\n", l.Rate*100)
	fmt.Printf("Loan Term: %v years\n", l.Term)
}
