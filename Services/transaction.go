package Services

import "time"

type Transaction struct {
	Id        int       `json:"Id"`
	Amount    float64   `json:"Amount"`
	Currency  string    `json:"Currency"`
	CreatedAt time.Time `json:"CreatedAt"`
}

type Transactions []Transaction

var transactions = Transactions{
	Transaction{Id: 125602, Amount: 1000, Currency: "MXN", CreatedAt: time.Now().UTC()},
	Transaction{Id: 125778, Amount: 2000, Currency: "COP", CreatedAt: time.Now().UTC()},
	Transaction{Id: 5545565, Amount: 3000, Currency: "CLP", CreatedAt: time.Now().UTC()},
	Transaction{Id: 125778, Amount: 4000, Currency: "USD", CreatedAt: time.Now().UTC()},
}
