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

type UserTransactionFormatter struct {
	ID        int               `json:"id"`
	Amount    int               `json:"amount"`
	Status    string            `json:"status"`
	CreatedAt time.Time         `json:"created_at"`
	Campaign  CampaignFormatter `json:"campaign"`
}

type CampaignFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func FormatUserTransaction(transaction Transaction) UserTransactionFormatter {
	formatTransaction := UserTransactionFormatter{}
	formatTransaction.ID = transaction.ID
	formatTransaction.Amount = transaction.Amount
	formatTransaction.Status = transaction.Status
	formatTransaction.CreatedAt = transaction.CreatedAt

	formatcampaign := CampaignFormatter{}
	formatcampaign.Name = transaction.Campaign.Name
	formatcampaign.ImageURL = ""

	if len(transaction.Campaign.CampaignImages) > 0 {
		formatcampaign.ImageURL = transaction.Campaign.CampaignImages[0].FileName
	}

	formatTransaction.Campaign = formatcampaign

	return formatTransaction
}

func FormatUserTransactions(transactions []Transaction) []UserTransactionFormatter {
	if len(transactions) == 0 {
		return []UserTransactionFormatter{}
	}
	var formatTransactions []UserTransactionFormatter

	for _, transaction := range transactions {
		formatTransaction := FormatUserTransaction(transaction)
		formatTransactions = append(formatTransactions, formatTransaction)
	}
	return formatTransactions
}
