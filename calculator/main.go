package main

import (
	"bufio"
	"calculations"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Printf("\nWelcome to the Loan Calculator\n")

	var loan calculations.Loan
	//Request loan information from user
	calculations.GetInformation(&loan)
	calculations.Delay(500)
	//Display main menu with user input values store
	menu(&loan)
}

func menu(loan *calculations.Loan) {
	reader := bufio.NewReader(os.Stdin)

	for {
		//Main menu
		calculations.Delay(500)
		fmt.Println("\n************************************")
		fmt.Println("*            MAIN MENU             *")
		fmt.Println("************************************")
		fmt.Printf(" Choose from the following options\n\n")
		fmt.Println(" 1 - CALCULATE MONTHLY PAYMENT AMOUNT")
		fmt.Println(" 2 - CALCULATE TOTAL NUMBER OF PAYMENTS")
		fmt.Println(" 3 - CALCULATE DAILY ACCRUED INTEREST")
		fmt.Println(" 4 - CALCULATE SIMPLE INTEREST AMOUNT")
		fmt.Println(" 5 - CALCULATE COMPOUND INTEREST AMOUNT")
		fmt.Println(" 6 - CALCULATE TOTAL INTEREST PAID ON LOAN")
		fmt.Println(" 7 - CALCULATE TOTAL COST OF LOAN")
		fmt.Println(" 8 - DISPLAY ALL LOAN INFORMATION")
		fmt.Println(" 9 - ENTER NEW LOAN INFORMATION")
		fmt.Println("10 - IMPORTANT LOAN TERMS")
		fmt.Println("11 - EXIT CALCULATOR")

		//Prompt user to enter a menu option
		fmt.Print("\nMenu Option: ")

		menu, err := reader.ReadString('\n')
		if err != nil || menu == "" {
			fmt.Println("Invalid input, please enter your selection again.")
			continue
		}
		menuOption := strings.TrimSpace(menu)
		option, err := strconv.Atoi(menuOption)
		if err != nil || option < 1 || option > 11 {
			fmt.Println("Invalid selection, please choose a valid option between 1 and 11.")
			continue
		}
		//Display to user what option they selected
		fmt.Printf("\nYou chose option %d\n\n", option)

		//Main menu controls
		switch option {
		case 1: //Calculate the monthly payment amount
			calculations.Delay(500)
			fmt.Println("-----------------------------------")
			fmt.Printf("Monthly Payment is: $%.2f\n", loan.MonthlyPayment())
			fmt.Printf("-----------------------------------\n")
			calculations.NextStep()
		case 2: //Calculate the total number of payments to make on the loan
			calculations.Delay(500)
			fmt.Println("----------------------------------------------------------------------------------------")
			fmt.Printf("Total number of payments is %d. You will make %d monthly payments to satisfy this loan.\n", loan.TotalPayments(), loan.TotalPayments())
			fmt.Printf("----------------------------------------------------------------------------------------\n\n")
			calculations.NextStep()
		case 3: //Calculate and display the daily accrued interest amount on the loan
			calculations.Delay(500)
			fmt.Println("-----------------------------------")
			fmt.Printf("Daily Accrued Interest is $%.2f\n", loan.DailyAccruedInterest())
			fmt.Printf("-----------------------------------\n\n")
			calculations.NextStep()
		case 4: //Calculate and display the simple interest amount on the loan
			calculations.Delay(500)
			fmt.Println("-----------------------------------")
			fmt.Printf("Simple Interest Amount is $%.2f\n", loan.SimpleInterest())
			fmt.Printf("-----------------------------------\n\n")
			calculations.NextStep()
		case 5: //Display the Compound Interest Menu
			calculations.Delay(500)
			calculations.CompoundInterestMenu(loan)
		case 6:
			//Calculate and display the total interest amount that will be paid on the loan
			calculations.Delay(500)
			fmt.Println("-----------------------------------")
			fmt.Printf("Total Interest Paid on Loan is $%.2f\n\n", loan.TotalInterestAmount())
			fmt.Printf("-----------------------------------\n\n")
			calculations.NextStep()
		case 7:
			//Calculate and display the total cost of the loan
			calculations.Delay(500)
			fmt.Println("-----------------------------------")
			fmt.Printf("Total Cost of Loan is $%.2f\n\n", loan.TotalCost())
			fmt.Printf("-----------------------------------\n\n")
			calculations.NextStep()
		case 8:
			//Calculate and display all loan information
			calculations.Delay(500)
			fmt.Println("         Calculated Loan Information")
			fmt.Println("-----------------------------------------------")
			fmt.Printf("Monthly Loan Payment: $%.2f\n", loan.MonthlyPayment())
			fmt.Printf("Number of Loan Payments: %d\n", loan.TotalPayments())
			fmt.Printf("Total Interest Paid on Loan: $%.2f\n", loan.TotalInterestAmount())
			fmt.Printf("Total Cost of Loan: $%.2f\n\n", loan.TotalCost())
			fmt.Println("Loan Interest Information (Simple vs. Compound)")
			fmt.Println("-----------------------------------------------")
			fmt.Println("\n---------Simple Interest Information---------")
			fmt.Printf("Daily Accrued Interest (Based on APR): $%.2f\n", loan.DailyAccruedInterest())
			fmt.Printf("Simple Interest Loan Amount: $%.2f\n\n", loan.SimpleInterest())

			fmt.Println("---------Compound Interest Information---------")
			fmt.Printf("Daily Compound Interest Loan Amount: $%.2f\n", loan.DailyCompoundInterest())
			fmt.Printf("Weekly Compound Interest Loan Amount: $%.2f\n", loan.WeeklyCompoundInterest())
			fmt.Printf("Monthly Compound Interest Loan Amount: $%.2f\n", loan.MonthlyCompoundInterest())
			fmt.Printf("Annual Compound Interest Loan Amount: $%.2f\n\n", loan.AnnualCompoundInterest())
			calculations.NextStep()
		case 9:
			//Enter new loan information
			calculations.Delay(500)
			calculations.GetInformation(loan)
		case 10:
			//Display important loan terms
			calculations.Delay(500)
			calculations.ImportantTerms()
		case 11:
			//Exit the program
			calculations.Delay(500)
			fmt.Printf("Exiting Loan Calculator. Goodbye!\n\n")
			os.Exit(0)
		}
	}
}
