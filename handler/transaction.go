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

	transactions, err := h.transactionService.GetCampaignTransactionByUserID(userID)
	if err != nil {
		response := helper.ApiResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("List of transactions", 200, "success", transaction.FormatUserTransactions(transactions))

	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) CreateTransaction(c *gin.Context) {
	var input transaction.CreateTransactionIput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errormessagge := gin.H{"errors": errors}

		response := helper.ApiResponse("Create transaction failed", http.StatusUnprocessableEntity, "error", errormessagge)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currenUser := c.MustGet("currentUser").(user.User)
	input.User = currenUser

	newCampaign, err := h.transactionService.CreateTransaction(input)
	if err != nil {
		response := helper.ApiResponse("Create transaction failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("transaction successfully created", 200, "success", newCampaign)

	c.JSON(http.StatusOK, response)
}
