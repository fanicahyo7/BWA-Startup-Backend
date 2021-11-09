package transaction

import "time"

type TransactionFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatTransaction(transaction Transaction) TransactionFormatter {
	formatTransaction := TransactionFormatter{}
	formatTransaction.ID = transaction.ID
	formatTransaction.Name = transaction.User.Name
	formatTransaction.Amount = transaction.Amount
	formatTransaction.CreatedAt = transaction.CreatedAt

	return formatTransaction
}

func FormatTransactions(transactions []Transaction) []TransactionFormatter {
	var formatTransactions []TransactionFormatter

	for _, transaction := range transactions {
		formatTransaction := FormatTransaction(transaction)
		formatTransactions = append(formatTransactions, formatTransaction)
	}
	return formatTransactions
}
