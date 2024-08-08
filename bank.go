package main

import (
	"fmt"
)

var greetUser string = "Hello User! Which operation would you like to make?\n1 - Check balance\n2 - Deposit amount\n3 - Withdraw\n4 - Exit\nEnter: "
var balance float64 = 0
var choice int
var deposit float64
var withdraw float64

func checkBalance() {
	fmt.Printf("Your balance is %.2f\n", balance)
}

func endOperation() {
	fmt.Print("Thanks for using Our Bank\n")
}

func depositAmount() {
	fmt.Print("How much would you like to deposit?: ")
	fmt.Scan(&deposit)
	balance += deposit
	fmt.Printf("You deposited %.2f to your account!\n", deposit)
	fmt.Printf("Your balance is %.2f\n", balance)
}

func withdrawAmount() {
	fmt.Print("How much would you like to withdraw?: ")
	fmt.Scan(&withdraw)
	if withdraw > balance {
		fmt.Println("Insufficient funds.")
	} else if balance == 0 {
		fmt.Printf("There is no amount in your account: %.2f\n", balance)
	} else {
		balance -= withdraw
		fmt.Printf("You have withdrawn %.2f from your account.\n", withdraw)
		fmt.Printf("Your balance is %.2f\n", balance)
	}
}

func main() {
	for {
		fmt.Print(greetUser)
		fmt.Scan(&choice)
		switch choice {
		case 1:
			checkBalance()
		case 2:
			depositAmount()
		case 3:
			withdrawAmount()
		case 4:
			endOperation()
			return // Exit the program
		default:
			fmt.Println("Please enter a valid value.")
		}
	}
}
