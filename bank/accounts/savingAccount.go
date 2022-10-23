package accounts

import (
	"studies/alura/bank/clients"
)

type SavingAccount struct {
	Holder                                        clients.Holder
	BankBranchNumber, AccountNumber, AccountDigit int
	balance                                       float64
}

// GETTERS
func (a SavingAccount) GetBalance() (string, float64) {
	return "Seu saldo é de R$ ", a.balance
}

func (a SavingAccount) GetAccountNumber() int {
	return a.AccountNumber
}

func (a SavingAccount) GetAccountDigit() int {
	return a.AccountDigit
}

func (a SavingAccount) GetBankBranchNumber() int {
	return a.BankBranchNumber
}

func (a SavingAccount) GetHolder() clients.Holder {
	return a.Holder
}

// SETTERS
func (a *SavingAccount) Withdraw(amount float64) (string, float64) {
	canWithdraw := amount > 0 && amount <= a.balance

	if canWithdraw {
		a.balance -= amount
		return "Saque realizado com sucesso. Saldo: R$", a.balance
	}

	if amount > a.balance {
		return "Saldo insuficiente. Saldo: R$", a.balance
	}

	return "Valor de saque inválido. R$", amount
}

func (a *SavingAccount) Deposit(amount float64) (string, float64) {
	canDeposit := amount > 0

	if canDeposit {
		a.balance += amount
		return "Depósito realizado com sucesso. Saldo: R$", a.balance
	}

	return "Valor de depósito inválido. R$", amount
}

func (a *SavingAccount) Transfer(amount float64, destinationAccount *SavingAccount) (string, float64) {
	if amount > 0 && amount <= a.balance {
		a.balance -= amount
		destinationAccount.Deposit(amount)
		return "Transferência realizada com sucesso. Saldo: R$", a.balance
	}

	if destinationAccount == nil {
		return "Conta de destino inválida. R$", amount
	}

	if &a == &destinationAccount {
		return "Não é possível transferir para a mesma conta. R$", amount
	}

	return "Saldo insuficiente. R$", amount
}
