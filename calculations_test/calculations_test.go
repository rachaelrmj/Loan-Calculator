package main

import (
	"calculations"
	"testing"
)

func TestSimpleInterest(t *testing.T) {
	loan := calculations.Loan{Principal: 1000, Term: 2, Rate: 0.05}
	want := 100
	got := loan.SimpleInterest()
	if got != want {
		t.Errorf("SimpleInterest() equals %v. Expected SimpleInterest() to equal %v", got, want)
	}
}

func TestCompoundInterest(t *testing.T) {
	loan := calculations.Loan{Principal: 1000, Term: 2, Rate: 0.05}
	want := 1104.94
	got := loan.AnnualCompoundInterest()
	if got != want {
		t.Errorf("AnnualCompoundInterest() equals %v. Expected AnnualCompoundInterest() to equal %v", got, want)
	}
}

func TestDailyCompoundInterest(t *testing.T) {
	loan := calculations.Loan{Principal: 1000, Term: 2, Rate: 0.05}
	want := loan.DailyCompoundInterest()
	got := loan.CompoundInterest(365)
	if got != want {
		t.Errorf("DailyCompoundInterest() equals %v. Expected DailyCompoundInterest() to equal %v", got, want)
	}
}

func TestWeeklyCompoundInterest(t *testing.T) {
	loan := calculations.Loan{Principal: 1000, Term: 2, Rate: 0.05}
	want := loan.WeeklyCompoundInterest()
	got := loan.CompoundInterest(52)
	if got != want {
		t.Errorf("WeeklyCompoundInterest() equals %v. Expected WeeklyCompoundInterest() to equal %v", got, want)
	}
}

func TestMonthlyCompoundInterest(t *testing.T) {
	loan := calculations.Loan{Principal: 1000, Term: 2, Rate: 0.05}
	got := loan.MonthlyCompoundInterest()
	want := loan.CompoundInterest(12)
	if got != want {
		t.Errorf("MonthlyCompoundInterest() equals %v. Expected MonthlyCompoundInterest() to equal %v", got, want)
	}
}

func TestAnnualCompoundInterest(t *testing.T) {
	loan := calculations.Loan{Principal: 1000, Term: 2, Rate: 0.05}
	got := 1576.25 // compound interest annually
	want := loan.AnnualCompoundInterest()
	if got != want {
		t.Errorf("AnnualCompoundInterest() equals %v. Expected AnnualCompoundInterest() to equal %v", got, want)
	}
}

func TestDailyAccruedInterest(t *testing.T) {
	loan := calculations.Loan{Principal: 1000, Term: 2, Rate: 0.05}
	got := (0.05 / 365) * 1000
	want := loan.DailyAccruedInterest()
	if got != want {
		t.Errorf("DailyAccruedInterest() equals %v. Expected DailyAccruedInterest() to equal %v", got, want)
	}
}

func TestDailyInterestRate(t *testing.T) {
	loan := calculations.Loan{Principal: 1000, Term: 2, Rate: 0.05}
	want := 0.05 / 365
	got := loan.DailyInterestRate()
	if got != want {
		t.Errorf("DailyInterestRate() equals %v. Expected DailyInterestRate() to equal %v", got, want)
	}
}

func TestMonthlyInterestRate(t *testing.T) {
	loan := calculations.Loan{Principal: 1000, Term: 2, Rate: 0.05}
	want := 0.004167
	got := loan.MonthlyInterestRate()
	if got != want {
		t.Errorf("MonthlyInterestRate() equals %v. Expected MonthlyInterestRate() to equal %v", got, want)
	}
}

func TestTotalPayments(t *testing.T) {
	loan := calculations.Loan{Principal: 1000, Term: 2, Rate: 0.05}
	want := 2 * 12
	got := loan.TotalPayments()
	if got != want {
		t.Errorf("TotalPayments() equals %v. Expected TotalPayments() to equal %v", got, want)
	}
}

func TestMonthlyPayment(t *testing.T) {
	loan := calculations.Loan{Principal: 1000, Term: 2, Rate: 0.05}
	want := 43.87 // approximate monthly payment
	got := loan.MonthlyPayment()
	if got != want {
		t.Errorf("MonthlyPayment() equals %v. Expected MonthlyPayment() to equal %v", got, want)
	}
}

func TestTotalCost(t *testing.T) {
	loan := calculations.Loan{Principal: 1000, Term: 2, Rate: 0.05}
	want := loan.MonthlyPayment() * float64(loan.TotalPayments())
	got := loan.TotalCost()
	if got != want {
		t.Errorf("TotalCost() equals %v. Expected TotalCost() to equal %v", got, want)
	}
}

func TestTotalInterestAmount(t *testing.T) {
	loan := calculations.Loan{Principal: 1000, Term: 2, Rate: 0.05}
	want := loan.TotalCost() - loan.Principal
	got := loan.TotalInterestAmount()
	if got != want {
		t.Errorf("TotalInterestAmount() equals %v. Expected TotalInterestAmount() to equal %v", got, want)
	}
}
