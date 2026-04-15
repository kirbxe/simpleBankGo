package bank

type Account struct {
	ID      int
	Balance int
}

func NewAccount(id int, balance int) *Account {

	return &Account{
		ID:      id,
		Balance: balance,
	}
}
