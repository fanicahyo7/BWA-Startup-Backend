package handler

import (
	"bwastartup/campaign"
	"bwastartup/helper"
	"bwastartup/user"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	campaignService campaign.Service
}

func NewCampaignHandler(campaignService campaign.Service) *campaignHandler {
	return &campaignHandler{campaignService}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))
	campaigns, err := h.campaignService.GetCampaigns(userID)
	if err != nil {
		response := helper.ApiResponse("error to get campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("list of campaigns", 200, "success", campaign.FormatCampaigns(campaigns))

	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	var input campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ApiResponse("error to get campaign detail", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaigns, err := h.campaignService.GetCampaign(input)
	if err != nil {
		response := helper.ApiResponse("error to get campaign detail", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatCampaigns := campaign.FormatCampaignDetail(campaigns)

	response := helper.ApiResponse("single campaigns", 200, "success", formatCampaigns)

	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	var input campaign.CampaignInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errormessagge := gin.H{"errors": errors}

		response := helper.ApiResponse("Create campaign failed", http.StatusUnprocessableEntity, "error", errormessagge)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currenUser := c.MustGet("currentUser").(user.User)
	input.User = currenUser

	newCampaign, err := h.campaignService.CreateCampaign(input)
	if err != nil {
		response := helper.ApiResponse("Create campaign failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Campaign successfully created", 200, "success", newCampaign)

	c.JSON(http.StatusOK, response)

}

func (h *campaignHandler) UpdateCampaign(c *gin.Context) {
	var input campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ApiResponse("update campaign failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputdata campaign.CampaignInput
	err = c.ShouldBindJSON(&inputdata)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errormessagge := gin.H{"errors": errors}

		response := helper.ApiResponse("update campaign failed", http.StatusUnprocessableEntity, "error", errormessagge)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currenUser := c.MustGet("currentUser").(user.User)
	inputdata.User = currenUser

	updateCampaign, err := h.campaignService.EditCampaign(input, inputdata)
	if err != nil {
		response := helper.ApiResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Campaign successfully updated", 200, "success", updateCampaign)

	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) UploadImage(c *gin.Context) {
	var input campaign.CampaignImageInput
	err := c.ShouldBind(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("error to upload image", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currenUser := c.MustGet("currentUser").(user.User)
	input.User = currenUser
	UserID := currenUser.ID

	file, err := c.FormFile("file")
	if err != nil {
		errormessagge := gin.H{"is_uploaded": false}

		response := helper.ApiResponse("error to upload image", http.StatusUnprocessableEntity, "error", errormessagge)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	path := fmt.Sprintf("images/%d-%s", UserID, file.Filename)
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		errormessagge := gin.H{"is_uploaded": false}

		response := helper.ApiResponse("error to upload image", http.StatusUnprocessableEntity, "error", errormessagge)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.campaignService.CreateImage(input, path)
	if err != nil {
		errormessagge := gin.H{"is_uploaded": false}

		response := helper.ApiResponse("upload image failed", http.StatusUnprocessableEntity, "error", errormessagge)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.ApiResponse("upload success", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}
