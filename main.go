package main

import (
	"bank/bank"
	"fmt"
	"sync"
	"time"
)

func main() {

	wg := sync.WaitGroup{}

	b1 := bank.NewBank(
		[]*bank.Account{
			bank.NewAccount(1, 0),
			bank.NewAccount(2, 0),
		},
		3,
	)
	module := bank.NewBankModule(b1)

	startTime := time.Now()
	
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {

				_ = module.BankMethod.Deposit(1, 10)
				_ = module.BankMethod.Withdraw(1, 5)
				_ = module.BankMethod.Transfer(1, 2, 1)
			}
		}()

	}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			module.BankMethod.GetBalance(1)
			module.BankMethod.GetBalance(2)
			total := module.BankMethod.Total()
			fmt.Println("Всего денег в банке: ", total)
		}()
	}
	wg.Wait()
	
	fmt.Println("Прошедшее время: ", time.Since(startTime))

}
