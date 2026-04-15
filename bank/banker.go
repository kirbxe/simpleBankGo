package bank

type Banker interface {
	Deposit(id int, amount int) error
	Withdraw(id int, amount int) error
	Transfer(FromID int, ToID int, amount int) error
	GetBalance(id int) error
	Total() int
}

type BankModule struct {
	BankMethod Banker
}

func NewBankModule(module Banker) BankModule {
	return BankModule{
		BankMethod: module,
	}
}
