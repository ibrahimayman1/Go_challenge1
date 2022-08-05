package Services

import (
	"errors"
	"strconv"
)

func GetAllTransaction() Transactions {
	return transactions
}

func GetTransaction(Id string) (error, *Transaction) {
	var transaction *Transaction
	ID, _ := strconv.Atoi(Id)
	for _, value := range transactions {
		if value.Id == ID {
			transaction = &value
		}
	}
	if transaction == nil {
		return errors.New("Transaction not found!"), nil
	} else {
		return nil, transaction
	}

}
func AddTransaction(trans *Transaction) Transactions {
	transactions := append(Transactions{}, *trans)
	return transactions
}
