package transaction

import "bwastartup/user"

type GetCampaignTransactionInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}

type CreateTransactionIput struct {
	Amount     int `json:"amount" binding:"required"`
	CampaignID int `json:"campaign_id" binding:"required"`
	User       user.User
}
