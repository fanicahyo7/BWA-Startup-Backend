package campaign

import "bwastartup/user"

type GetCampaignDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type CampaignInput struct {
	Name             string
	ShortDescription string
	Description      string
	Perks            string
	GoalAmount       int
	User             user.User
}

type CampaignImageInput struct {
	CampaignID int  `form:"campaign_id" binding:"required"`
	IsPrimary  bool `form:"is_primary"`
	User       user.User
}
