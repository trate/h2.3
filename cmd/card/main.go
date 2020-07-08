package main

import (
	"fmt"
	"github.com/trate/h2.3/pkg/card"
	"time"
)

func main() {
	t1 := &card.Transaction{
		Id:     1,
		Amount: -735_55,
		Date:   time.Date(2020, 3, 1, 0, 0, 0, 0, time.UTC),
		MCC:    "5411",
		Status: "InProgress",
	}
	t2 := &card.Transaction{
		Id:     2,
		Amount: -736_55,
		Date:   time.Date(2020, 3, 2, 0, 0, 0, 0, time.UTC),
		MCC:    "5411",
		Status: "InProgress",
	}
	t3 := &card.Transaction{
		Id:     3,
		Amount: 2_000_00,
		Date:   time.Date(2020, 4, 1, 0, 0, 0, 0, time.UTC),
		MCC:    "0000",
		Status: "Done",
	}
	t4 := &card.Transaction{
		Id:     4,
		Amount: 2_100_00,
		Date:   time.Date(2020, 4, 2, 0, 0, 0, 0, time.UTC),
		MCC:    "0000",
		Status: "Done",
	}
	t5 := &card.Transaction{
		Id:     5,
		Amount: -1_203_91,
		Date:   time.Date(2020, 5, 1, 0, 0, 0, 0, time.UTC),
		MCC:    "5812",
		Status: "InProgress",
	}
	t6 := &card.Transaction{
		Id:     6,
		Amount: -1_204_91,
		Date:   time.Date(2020, 5, 2, 0, 0, 0, 0, time.UTC),
		MCC:    "5812",
		Status: "InProgress",
	}
	transactions := []card.Transaction{*t1, *t2, *t3, *t4, *t5}

	master := &card.Card{
		Id:           1,
		Issuer:       "MasterCard",
		Balance:      65_000,
		Currency:     "RUB",
		Number:       "5177827685644009",
		Transactions: transactions,
	}

	card.AddTransaction(master, t6)
	fmt.Println("Выводим исходную структуру..")
	fmt.Println(master)

	fmt.Println("Выводим исходную структуру с отсортированными по сумме транзакциями..")
	fmt.Println(master.TransactionsSortBySum())
	fmt.Println("Выводим суммы подсчитанные горутинами, но не выводим общую сумму....")
	master.TransactionsSumConcurrently()

}