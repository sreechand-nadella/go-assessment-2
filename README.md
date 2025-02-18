# Calculator

A simple command-line calculator built in Go.

## Overview

This project is a basic calculator that supports four operations:
- **Addition**
- **Subtraction**
- **Multiplication**
- **Division**

It accepts command-line arguments to perform calculations.

## Features

- **Command-Line Interface:** Run operations directly from your terminal.
- **Error Handling:** Gracefully handles invalid inputs and division by zero.
- **Modular Code:** Separates the main function from the arithmetic functionalities.

## Project Structure

repo/ ├── calculator/ # Contains the main calculator project files. │ ├── main.go # Entry point of the application. │ └── calc.go # Contains the arithmetic functions: Add, Subtract, Multiply, Divide. ├── .gitignore # Specifies intentionally untracked files to ignore. ├── go.mod # Go module file. └── README.md # This file.

## Prerequisites

- [Go](https://golang.org/dl/) installed on your machine.
- A terminal/command prompt for running the commands.

## Getting Started

1. **Clone the Repository:**
   ```bash
   git clone https://github.com/sreechand-nadella/calculator.git
   cd calculator
Initialize the Go Module (if not already done):


go mod tidy
Run the Calculator:

To run the application with command-line arguments, use:


go run ./calculator add 10 5
This example adds 10 and 5.

Other Operations:

Subtraction:
go run ./calculator subtract 10 5
Multiplication:
go run ./calculator multiply 10 5
Division:
go run ./calculator divide 10 5

## Usage

The general command format is:

go run ./calculator <operator> <num1> <num2>
<operator>: One of add, subtract, multiply, or divide.
<num1> and <num2>: Numbers to perform the operation on.
