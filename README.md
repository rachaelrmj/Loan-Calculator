# Loan Calculator

A menu-driven command-line application written in Go that helps users analyze loans and understand borrowing costs through a variety of financial calculations.

The application accepts loan information from the user and provides calculations for monthly payments, interest accumulation, total loan costs, and compound interest scenarios. The project also includes automated unit testing to validate financial calculations.

---

## Overview

This project was developed to demonstrate practical application of Go programming fundamentals while implementing real-world financial formulas used by banks, lenders, and financial institutions.

Users can enter:

* Loan Amount (Principal)
* Annual Percentage Rate (APR)
* Loan Term (Years)

The program then generates loan metrics through an interactive menu system.

---

## Features

### Loan Payment Analysis

* Calculate Monthly Payment Amount
* Calculate Total Number of Payments
* Calculate Total Cost of Loan
* Calculate Total Interest Paid

### Interest Calculations

* Daily Accrued Interest
* Simple Interest
* Compound Interest

### Compound Interest Options

* Daily Compounding
* Weekly Compounding
* Monthly Compounding
* Annual Compounding

### User Experience

* Input Validation
* Interactive Menu System
* Error Handling
* Loan Information Review and Confirmation
* Educational Loan Terminology Reference

### Quality Assurance

* Unit Testing with Go Testing Package
* Validation of Financial Formulas
* Verification of Interest Calculations
* Verification of Payment Calculations

---

## Technologies Used

### Language

* Go (Golang)

### Standard Libraries

* fmt
* math
* os
* bufio
* strings
* strconv
* log
* testing

---

## Project Structure

```text
loan-calculator/
│
├── main.go
│
├── calculations/
│   ├── calculations.go
│
├── calculations_test.go
│
└── README.md
```

---

## Financial Formulas Implemented

### Monthly Loan Payment

The calculator uses the standard amortization formula:

M = P × [r(1+r)^n] / [(1+r)^n − 1]

Where:

* M = Monthly Payment
* P = Principal
* r = Monthly Interest Rate
* n = Total Number of Payments

---

### Simple Interest

I = P × R × T

Where:

* P = Principal
* R = Interest Rate
* T = Time in Years

---

### Compound Interest

A = P(1 + r/n)^(nt)

Where:

* P = Principal
* r = Annual Interest Rate
* n = Number of Compounding Periods
* t = Time

---

### Daily Accrued Interest

Daily Interest = Principal × (APR / 365)

---

## Menu Options

| Option | Function                           |
| ------ | ---------------------------------- |
| 1      | Calculate Monthly Payment Amount   |
| 2      | Calculate Total Number of Payments |
| 3      | Calculate Daily Accrued Interest   |
| 4      | Calculate Simple Interest          |
| 5      | Calculate Compound Interest        |
| 6      | Calculate Total Interest Paid      |
| 7      | Calculate Total Cost of Loan       |
| 8      | Display All Loan Information       |
| 9      | Enter New Loan Information         |
| 10     | Display Important Loan Terms       |
| 11     | Exit Program                       |

---

## Running the Program

### Clone Repository

```bash
git clone https://github.com/yourusername/loan-calculator-go.git
```

### Navigate to Project Folder

```bash
cd loan-calculator-go
```

### Run Program

```bash
go run main.go
```

---

## Running Unit Tests

Execute all tests using:

```bash
go test ./...
```

The test suite validates:

* Simple Interest Calculations
* Compound Interest Calculations
* Daily Interest Calculations
* Monthly Interest Rate Calculations
* Monthly Payment Calculations
* Total Cost Calculations
* Total Interest Calculations

---

## Sample Output

```text
Welcome to the Loan Calculator

Loan Amount: 25000
Interest Rate (APR): 5
Loan Term: 5

You entered the following information:

Loan Amount: $25000.00
Interest Rate: 5.00%
Loan Term: 5 years
```

---

## Skills Demonstrated

### Go Development

* Packages
* Structs
* Methods
* Functions
* Switch Statements
* Loops
* Error Handling
* Input Validation

### Software Engineering

* Modular Design
* Separation of Concerns
* Unit Testing
* User-Centered Development

### Financial Programming

* Loan Amortization
* Interest Calculations
* Mathematical Modeling
* Financial Data Processing

---

## Future Enhancements

* Loan Amortization Schedule
* Payment Breakdown Tables
* CSV Export
* PDF Report Generation
* Graphical User Interface (GUI)
* Web Application Version
* REST API Support
* Mortgage Calculator Features
* Auto Loan Calculator Features
* Investment Calculator Features

---

## Author

Rachael R. Martinez-Jones

Arizona State University
Bachelor of Science – Graphic Information Technology
Concentration: Full Stack Web Development

---

## License

This project is intended for educational, portfolio, and professional demonstration purposes.

