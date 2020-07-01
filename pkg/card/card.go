package card

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

type Card struct {
	Id           int64
	Issuer       string
	Balance      int64
	Currency     string
	Number       string
	Transactions []Transaction
}

type Transaction struct {
	Id     int64
	Amount int64
	Date   int64
	MCC    string
	Status string
}

func AddTransaction(card *Card, transaction *Transaction) {
	card.Transactions = append(card.Transactions, *transaction)
}

func (c *Card) TransactionsSortBySum() []Transaction {
	tr := make([]Transaction, len(c.Transactions))
	copy(tr, c.Transactions)
	sort.SliceStable(tr, func(i, j int) bool {return !(tr[i].Amount < tr[j].Amount)})
	return tr
}
// for the sake of simplicity consider only the months of the spring
func (c *Card) TransactionsSumConcurrently(goroutines int) {
	// March, April, May
	tMarch := time.Date(2020, 03, 01, 00, 0, 0, 0, time.UTC) // Time
	unixMarch := tMarch.Unix() // int64

	tApril := time.Date(2020, 04, 01, 00, 0, 0, 0, time.UTC) // Time
	unixApril := tApril.Unix() // int64

	tMay := time.Date(2020, 05, 01, 00, 0, 0, 0, time.UTC) // Time
	unixMay := tMay.Unix()

	tJune := time.Date(2020, 06, 01, 00, 0, 0, 0, time.UTC) // Time
	unixJune := tJune.Unix()

	wg := sync.WaitGroup{}
	wg.Add(goroutines)

	// slice our transactions according to the spring's months
	var trMarch, trApril, trMay []int64
	for _, v := range c.Transactions {
		if v.Date >= unixMarch && v.Date < unixApril {
			trMarch = append(trMarch, v.Amount)
		}
		if v.Date >= unixApril && v.Date < unixMay {
			trApril = append(trApril, v.Amount)
		}
		if v.Date >= unixMay && v.Date < unixJune {
			trMay = append(trMay, v.Amount)
		}
	}
	var transactions [][]int64
	transactions = append(transactions, trMarch)
	transactions = append(transactions, trApril)
	transactions = append(transactions, trMay)

	// avoid calculating the total sum due to the race condition (need to use atomic operations)
	for i := 0; i < goroutines; i++ {
		part := transactions[i]
		go func() {
			var sum int64
			for _, v := range part {
				sum += v
			}
			fmt.Println(sum)
			wg.Done()
		}()
	}
	wg.Wait()
}

