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
	Date   time.Time
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
func (c *Card) TransactionsSumConcurrently() {

	splitTranz := make(map[string][]Transaction)

	for _, v := range c.Transactions {
		splitTranz[v.Date.Format("2006-01")] = append(splitTranz[v.Date.Format("2006-01")], v)
	}

	wg := sync.WaitGroup{}
	wg.Add(len(splitTranz))


	// avoid calculating the total sum due to the race condition (need to use atomic operations)
	for _, tr := range splitTranz {
		part := tr
		go func() {
			var sum int64
			for _, v := range part {
				sum += v.Amount
			}
			fmt.Println(sum)
			wg.Done()
		}()
	}
	wg.Wait()
}

