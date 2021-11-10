package handler

import (
	"bwastartup/helper"
	"bwastartup/transaction"
	"bwastartup/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	transactionService transaction.Service
}

func NewTransactionHandler(transactionService transaction.Service) *transactionHandler {
	return &transactionHandler{transactionService}
}

func (h *transactionHandler) GetCampaignTransactionByCampaignID(c *gin.Context) {
	var input transaction.GetCampaignTransactionInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ApiResponse("error to get campaign transaction 1", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currenUser := c.MustGet("currentUser").(user.User)
	input.User = currenUser

	transactions, err := h.transactionService.GetCampaignTransactionByCampaignID(input)
	if err != nil {
		response := helper.ApiResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("List of transactions", 200, "success", transaction.FormatTransactions(transactions))

	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) GetCampaignTransactionByUserID(c *gin.Context) {
	currenUser := c.MustGet("currentUser").(user.User)
	userID := currenUser.ID
	fmt.Println(userID)

	transaction, err := h.transactionService.GetCampaignTransactionByUserID(userID)
	if err != nil {
		response := helper.ApiResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("List of transactions", 200, "success", transaction)

	c.JSON(http.StatusOK, response)
}
