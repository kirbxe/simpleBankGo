package bank

import (
	"errors"
	"fmt"
	"sync"
)

type Bank struct {
	Accounts []*Account
	mtx      sync.Mutex
	sem chan struct{}
}

func NewBank(accounts []*Account, maxTransfers int) *Bank {
	return &Bank{
		Accounts: accounts,
		sem: make(chan struct{}, maxTransfers),
	}
}

// Получает: id клиента, количество денег
// Ищет клиента в базе
// Прибавляет к балансу клиента количество денег

func (b *Bank) Deposit(id int, amount int) error {
	b.mtx.Lock()
	defer b.mtx.Unlock()

	for _, v := range b.Accounts {
		if v.ID == id {

			v.Balance += amount
			return nil
		}
	}
	return errors.New("Клиент с таким айди не существует")
}

func (b *Bank) Withdraw(id int, amount int) error {
	b.mtx.Lock()
	defer b.mtx.Unlock()
	for _, v := range b.Accounts {
		if v.ID == id {
			if v.Balance >= amount {

				v.Balance -= amount
				return nil
			} else {
				return errors.New("Не хватает баланса на аккаунте")
			}

		}
	}
	return errors.New("Клиента под таким айди не существует")

}

// Найти кто перечисливает
// Проверить баланс отправителя
// Найти того кто получает
// Вычесть из баланса отправителя, добавить к балансу получателя
func (b *Bank) Transfer(fromID int, toID int, amount int) error {
	b.sem <- struct{}{}
	defer func() { <-b.sem }()
	
	b.mtx.Lock()
	defer b.mtx.Unlock()
	var fromAcc, toAcc *Account

	for _, acc := range b.Accounts {
		switch acc.ID {
		case fromID:
			fromAcc = acc
		case toID:
			toAcc = acc
		}
	}

	if fromAcc == nil {
		return errors.New("Айди отправителя не существует")
	}

	if toAcc == nil {

		return errors.New("Айди получателя не существует")

	}

	if fromAcc.Balance < amount {

		return errors.New("Баланса не хватает")

	}

	fromAcc.Balance -= amount
	toAcc.Balance += amount

	return nil

	// for _, from := range b.accounts {

	// 	if from.ID == FromID {
	// 		if from.Balance >= amount {
	// 			for _, to := range b.accounts {
	// 				if to.ID == ToID {
	// 					from.Balance -= amount
	// 					to.Balance += amount
	// 					return nil
	// 				}
	// 			}
	// 			return errors.New("Клиента под таким айди не существует")

	// 		}
	// 		return errors.New("Не хватает баланса")
	// 	}
	// }
	// return errors.New("Клиента под таким айди не существует")
}

func (b *Bank) GetBalance(id int) error {
	b.mtx.Lock()
	defer b.mtx.Unlock()
	for _, acc := range b.Accounts {

		if acc.ID == id {

			fmt.Println("Клиент с айди: ", acc.ID, ". ", "Баланс: ", acc.Balance)
			return nil
		}
	}
	return errors.New("Клиент с таким айди не существует.")

}

func (b *Bank) Total() (int) {
	b.mtx.Lock()
	defer b.mtx.Unlock()
	sum := 0
	for _,v := range b.Accounts {
		sum += v.Balance
	}
	
	return sum
	
}
