package transaction

import (
	"bwastartup/campaign"
	"bwastartup/user"
	"time"

	"github.com/leekchan/accounting"
)

type Transaction struct {
	ID         int
	CampaignID int
	UserID     int
	Amount     int
	Status     string
	Code       string
	User       user.User
	Campaign   campaign.Campaign
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (c Transaction) AmountFormatIDR() string {
	acc := accounting.Accounting{Symbol: "Rp.", Precision: 2, Thousand: ".", Decimal: ","}
	return acc.FormatMoney(c.Amount)
}
