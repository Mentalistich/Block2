package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

type BankAccount struct {
	Owner           string
	Balance         float64
	AccountType     string
	dateOfOperation time.Time
	OwnerHistory    map[time.Time]string
}

func (b *BankAccount) Deposit(amount float64) {
	b.Balance += amount
	b.dateOfOperation = time.Now()
	b.OwnerHistory[b.dateOfOperation] = "The account " + b.Owner + " was topped up on " + strconv.FormatFloat(amount, 'f', 2, 64)
	time.Sleep(2 * time.Second)
}

func (b *BankAccount) Withdraw(amount float64) {
	b.Balance -= amount
	b.dateOfOperation = time.Now()
	b.Commission(amount)
	if b.AccountType == "Cheking" {
		b.OwnerHistory[b.dateOfOperation] = "the account was debited for" + strconv.FormatFloat(amount, 'f', 2, 64) + "+ commision" + strconv.FormatFloat(amount/100, 'f', 2, 64)
		time.Sleep(2 * time.Second)
	} else {
		b.OwnerHistory[b.dateOfOperation] = "the account was debited for" + strconv.FormatFloat(amount, 'f', 2, 64)
		time.Sleep(2 * time.Second)
	}

}
func (b *BankAccount) GetBalance() {
	fmt.Println("The balance of account " + b.Owner + " is " + strconv.FormatFloat(b.Balance, 'f', 2, 64))
	b.dateOfOperation = time.Now()
	b.OwnerHistory[b.dateOfOperation] = "Balance was checked"
	time.Sleep(2 * time.Second)
}

func (b *BankAccount) Transfer(to *BankAccount, amount float64) error {

	if b.AccountType != "Blocked" {
		if b.Balance < amount {
			return errors.New("not enough balance")
		} else {
			b.Balance -= amount
			to.Balance += amount
			b.dateOfOperation = time.Now()
			b.OwnerHistory[b.dateOfOperation] = strconv.FormatFloat(amount, 'f', 2, 64) + " was transferred from " + b.Owner + "`s account to " + to.Owner + "`s account"
			time.Sleep(2 * time.Second)
			return nil
		}
	} else {
		return errors.New("You cannot transfer funds because your account is blocked")
	}
}

func (b *BankAccount) CalculateInterest() {
	if b.AccountType == "Savings" {
		if b.Balance > 10000 {
			b.Balance = b.Balance * 1.05
			b.dateOfOperation = time.Now()
			b.OwnerHistory[b.dateOfOperation] = strconv.FormatFloat(b.Balance*0.05, 'f', 2, 64) + " have been credited to the account (cumulative interest)"
			time.Sleep(2 * time.Second)
		} else {
			b.Balance = b.Balance * 1.1
			b.dateOfOperation = time.Now()
			b.OwnerHistory[b.dateOfOperation] = strconv.FormatFloat(b.Balance*0.1, 'f', 2, 64) + " have been credited to the account (cumulative interest)"
			time.Sleep(2 * time.Second)
		}
	}
}

func (b *BankAccount) Commission(amount float64) {
	if b.AccountType == "Cheking" {
		b.Balance -= amount / 100
	}
}

func (b *BankAccount) GenerateStatement(begin time.Time, end time.Time) {
	for key, value := range b.OwnerHistory {
		if key.After(begin) && key.Before(end) {
			fmt.Println(value + " " + key.Format("02-01-2006 15:04:05"))
		}
	}

}

func main() {
	account1 := BankAccount{
		Owner:        "John",
		Balance:      100.0,
		AccountType:  "Savings",
		OwnerHistory: make(map[time.Time]string),
	}
	account2 := BankAccount{
		Owner:        "Silver",
		Balance:      100.0,
		AccountType:  "Cheking",
		OwnerHistory: make(map[time.Time]string),
	}
	account1.Transfer(&account2, 10)
	time1, err := time.Parse("02-01-2006 15:04:05", "24-10-2024 18:30:00")
	if err != nil {
		errors.New("Invalid time format")
	}
	time2, err := time.Parse("02-01-2006 15:04:05", "25-10-2024 22:30:00")
	if err != nil {
		errors.New("Invalid time format")
	}

	account1.CalculateInterest()
	account1.Deposit(15)
	//account1.GetBalance()
	//time1 := time.Now()
	account1.Deposit(25)
	//time2 := time.Now()
	account1.GenerateStatement(time1, time2)

}
