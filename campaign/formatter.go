package campaign

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
