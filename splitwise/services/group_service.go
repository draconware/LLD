package services

import (
	"fmt"

	generator "github.com/mastik5h/LLD/splitwise/helpers"
	"github.com/mastik5h/LLD/splitwise/models"
	"golang.org/x/exp/slices"

	data_service "github.com/mastik5h/LLD/splitwise/dataServices"
)

func CreateGroup(metadata *models.GroupMetaData, users []string) models.CreateGroupResponse {
	response := models.CreateGroupResponse{
		GroupId:      "",
		ErrorDetails: "",
	}
	if !ValidateUsersId(users) {
		response.ErrorDetails = "some of the users are not registered"
		return response
	}
	if err := validateGroupMetadata(metadata); err != nil {
		response.ErrorDetails = err.Error()
		return response
	}
	groupId := generator.GetUniqueIdString()
	group := &models.Group{
		Id:       groupId,
		Metadata: metadata,
		Users:    users,
	}
	err := data_service.CreateGroup(group)
	if err != nil {
		response.ErrorDetails = "internal error: failed creating group"
		return response
	}
	response.GroupId = groupId
	fmt.Println("Succesfully created the group. name: ", metadata.Name)
	return response
}

func CreateGroupExpense(userId string, groupId string, expenseMetaData *models.ExpenseMetaData, totalAmount float64, paidAmounts []*models.UserBalanceMap, oweAmounts []*models.UserBalanceMap) models.CreateGroupExpenseResponse {
	response := models.CreateGroupExpenseResponse{}
	if err := validateGroupId(groupId); err != nil {
		response.ErrorDetails = "invalid groupid found"
		return response
	}
	group := &models.Group{}
	group, err := data_service.GetGroup(groupId)
	if err != nil {
		response.ErrorDetails = err.Error()
		return response
	}
	group_users := group.Users

	if len(group_users) == 0 || !slices.Contains(group_users, userId) {
		response.ErrorDetails = "user is not authorized to add expense in this group"
		return response
	}

	groupExpenseResponse := CreateExpense(userId, groupId, expenseMetaData, totalAmount, paidAmounts, oweAmounts)
	if groupExpenseResponse.ErrorDetails != "" {
		response.ErrorDetails = "internal error: cannot create expense." + groupExpenseResponse.ErrorDetails
		return response
	}
	response.ExpenseId = groupExpenseResponse.ExpenseId
	return response
}

func SettleGroupExpense(userId, groupId string) models.SettleGroupExpenseResponse {
	resp := models.SettleGroupExpenseResponse{}
	if err := validateGroupId(groupId); err != nil {
		resp.ErrorDetails = "invalid group id provided."
		return resp
	}
	group, err := data_service.GetGroup(groupId)
	if err != nil {
		resp.ErrorDetails = err.Error()
		return resp
	}
	if !slices.Contains(group.Users, userId) {
		resp.ErrorDetails = "user is not authorized to add expense in this group"
		return resp
	}

	group_expenses, err := GetExpenseByGroupId(groupId)
	if err != nil {
		resp.ErrorDetails = err.Error()
		return resp
	}

	paymentGraph, err := GetPaymentGraphForGroupExpenses(group_expenses, len(group.Users))
	if err != nil {
		resp.ErrorDetails = "error fetching payment graph" + err.Error()
		return resp
	}

	resp.Transactions = paymentGraph
	return resp
}

func validateGroupMetadata(md *models.GroupMetaData) error {
	// need to implement
	return nil
}
func validateGroupId(gid string) error {
	//  need to implement
	return nil
}
