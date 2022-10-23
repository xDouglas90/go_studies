package main

import (
	"fmt"

	"studies/alura/bank/accounts"
	"studies/alura/bank/clients"
)

func payBankSlip(account accountVerify, amount float64) {
	account.Withdraw(amount)
}

type accountVerify interface {
	Withdraw(amount float64) (string, float64)
}

func main() {
	client := clients.Holder{}
	client.SetHolder("João da Silva", "123.456.789-10", "Programador")

	checkingAccount := accounts.CheckingAccount{Holder: client, BankBranchNumber: 123, AccountNumber: 123456, AccountDigit: 1, MonthlyFee: 10.00, DebitDay: 10, HasCreditCard: true}
	savingAccount := accounts.SavingAccount{Holder: client, BankBranchNumber: 123, AccountNumber: 123456, AccountDigit: 1}

	checkingAccount.Deposit(100)
	savingAccount.Deposit(100)

	payBankSlip(&checkingAccount, 10)
	payBankSlip(&savingAccount, 10)

	fmt.Println(checkingAccount.GetBalance())
	fmt.Println(savingAccount.GetBalance())
	fmt.Println(client.GetHolder())
}
