package campaign

import "strings"

type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	formatCampaign := CampaignFormatter{}
	formatCampaign.ID = campaign.ID
	formatCampaign.UserID = campaign.UserID
	formatCampaign.Name = campaign.Name
	formatCampaign.ShortDescription = campaign.ShortDescription
	formatCampaign.GoalAmount = campaign.GoalAmount
	formatCampaign.CurrentAmount = campaign.CurrentAmount
	formatCampaign.Slug = campaign.Slug
	formatCampaign.ImageURL = ""

	if len(campaign.CampaignImages) > 0 {
		formatCampaign.ImageURL = campaign.CampaignImages[0].FileName
	}

	return formatCampaign
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {

	campaignsformatter := []CampaignFormatter{}

	for _, campaign := range campaigns {
		campignFormatter := FormatCampaign(campaign)
		campaignsformatter = append(campaignsformatter, campignFormatter)
	}
	return campaignsformatter
}

type CampaignDetailFormatter struct {
	ID               int                      `json:"id"`
	Name             string                   `json:"name"`
	ShortDescription string                   `json:"short_description"`
	ImageURL         string                   `json:"image_url"`
	GoalAmount       int                      `json:"goal_amount"`
	CurrentAmount    int                      `json:"current_amount"`
	UserID           int                      `json:"user_id"`
	Slug             string                   `json:"slug"`
	Description      string                   `json:"description"`
	User             CampaignUserFormatter    `json:"user"`
	Perks            []string                 `json:"perks"`
	Images           []CampaignImageFormatter `json:"images"`
}

type CampaignUserFormatter struct {
	Name           string `json:"name"`
	AvatarFileName string `json:"image_url"`
}

type CampaignImageFormatter struct {
	FileName  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatCampaignDetail(campaign Campaign) CampaignDetailFormatter {
	formatCampaignDetail := CampaignDetailFormatter{}
	formatCampaignDetail.ID = campaign.ID
	formatCampaignDetail.Name = campaign.Name
	formatCampaignDetail.ShortDescription = campaign.ShortDescription
	formatCampaignDetail.ImageURL = ""

	if len(campaign.CampaignImages) > 0 {
		formatCampaignDetail.ImageURL = campaign.CampaignImages[0].FileName
	}

	formatCampaignDetail.GoalAmount = campaign.GoalAmount
	formatCampaignDetail.CurrentAmount = campaign.CurrentAmount
	formatCampaignDetail.UserID = campaign.UserID
	formatCampaignDetail.Slug = campaign.Slug
	formatCampaignDetail.Description = campaign.Description

	user := campaign.User

	formatCampaignUser := CampaignUserFormatter{}
	formatCampaignUser.Name = user.Name
	formatCampaignUser.AvatarFileName = user.AvatarFileName

	formatCampaignDetail.User = formatCampaignUser

	var perks []string

	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, perk)
	}
	formatCampaignDetail.Perks = perks

	images := []CampaignImageFormatter{}

	for _, image := range campaign.CampaignImages {
		formatCampaignImage := CampaignImageFormatter{}
		formatCampaignImage.FileName = image.FileName

		isPrimary := false
		if image.IsPrimary == 1 {
			isPrimary = true
		}
		formatCampaignImage.IsPrimary = isPrimary

		images = append(images, formatCampaignImage)
	}
	formatCampaignDetail.Images = images

	return formatCampaignDetail
}
