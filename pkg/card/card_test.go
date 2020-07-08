package card

import (
	"reflect"
	"testing"
	"time"
)

func TestCard_TransactionsSortBySum(t *testing.T) {
	type fields struct {
		Id           int64
		Issuer       string
		Balance      int64
		Currency     string
		Number       string
		Transactions []Transaction
	}
	tests := []struct {
		name   string
		fields fields
		want   []Transaction
	}{
		{
			name: "Проверяем, что слайс останется неизменным при сортировке по убыванию для 1-й транзакции",
			fields: fields{
				Id:       1,
				Issuer:   "Visa",
				Balance:  150_000_00,
				Currency: "RUB",
				Number:   "4716662464958306916",
				Transactions: []Transaction{
					{
						Id:     1,
						Amount: -12_000_00,
						Date:   time.Date(2020, 3, 1, 0, 0, 0, 0, time.UTC),
						MCC:    "5812",
						Status: "InProgress",
					},
				},
			},
			want: []Transaction{
				{
					Id:     1,
					Amount: -12_000_00,
					Date:   time.Date(2020, 3, 1, 0, 0, 0, 0, time.UTC),
					MCC:    "5812",
					Status: "InProgress",
				},
			},
		},
		{
			name: "Проверяем, что слайс отсортируется по убыванию для 3-х транзакций",
			fields: fields{
				Id:       1,
				Issuer:   "Visa",
				Balance:  150_000_00,
				Currency: "RUB",
				Number:   "4716662464958306916",
				Transactions: []Transaction{
					{
						Id:     1,
						Amount: -12_000_00,
						Date:   time.Date(2020, 7, 1, 0, 0, 0, 0, time.UTC),
						MCC:    "5812",
						Status: "InProgress",
					},
					{
						Id:     1,
						Amount: -124_000_00,
						Date:   time.Date(2020, 7, 1, 0, 0, 0, 0, time.UTC),
						MCC:    "5812",
						Status: "InProgress",
					},
					{
						Id:     1,
						Amount: 12_000_00,
						Date:   time.Date(2020, 7, 1, 0, 0, 0, 0, time.UTC),
						MCC:    "5812",
						Status: "InProgress",
					},
				},
			},
			want: []Transaction{
				{
					Id:     1,
					Amount: 12_000_00,
					Date:   time.Date(2020, 7, 1, 0, 0, 0, 0, time.UTC),
					MCC:    "5812",
					Status: "InProgress",
				},
				{
					Id:     1,
					Amount: -12_000_00,
					Date:   time.Date(2020, 7, 1, 0, 0, 0, 0, time.UTC),
					MCC:    "5812",
					Status: "InProgress",
				},
				{
					Id:     1,
					Amount: -124_000_00,
					Date:   time.Date(2020, 7, 1, 0, 0, 0, 0, time.UTC),
					MCC:    "5812",
					Status: "InProgress",
				},
			},
		},
	}
	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Card{
				Id:           tt.fields.Id,
				Issuer:       tt.fields.Issuer,
				Balance:      tt.fields.Balance,
				Currency:     tt.fields.Currency,
				Number:       tt.fields.Number,
				Transactions: tt.fields.Transactions,
			}
			if got := c.TransactionsSortBySum(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransactionsSortBySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
