package accounts

import (
	"studies/alura/bank/clients"
)

type CheckingAccount struct {
	Holder                                                  clients.Holder
	BankBranchNumber, AccountNumber, AccountDigit, DebitDay int
	balance, MonthlyFee                                     float64
	HasCreditCard                                           bool
}

// GETTERS
func (a CheckingAccount) GetBalance() (string, float64) {
	return "Seu saldo é de R$ ", a.balance
}

// SETTERS
func (a *CheckingAccount) Withdraw(amount float64) (string, float64) {
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

func (a *CheckingAccount) Deposit(amount float64) (string, float64) {
	canDeposit := amount > 0

	if canDeposit {
		a.balance += amount
		return "Depósito realizado com sucesso. Saldo: R$", a.balance
	}

	return "Valor de depósito inválido. R$", amount
}

func (a *CheckingAccount) Transfer(amount float64, destinationAccount *CheckingAccount) (string, float64) {
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

	if amount > a.balance {
		return "Saldo insuficiente. Saldo: R$", a.balance
	}

	return "Valor de transferência inválido. R$", amount
}
