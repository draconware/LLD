package client

import (
	"fmt"

	generator "github.com/mastik5h/LLD/splitwise/helpers"
	"github.com/mastik5h/LLD/splitwise/models"
	"github.com/mastik5h/LLD/splitwise/services"
)

func Initialize() {
	generator.InitializeStringGenerator()
}
func CreateExpenseClient(userid, expenseName, title, description string, totalAmount float64, paidAmounts, oweAmounts map[string]float64) string {
	expenseMetaData := &models.ExpenseMetaData{
		Name:        expenseName,
		Description: description,
		Title:       title,
	}

	paObj := createUserBalanceMapList(paidAmounts)
	oaObj := createUserBalanceMapList(oweAmounts)
	resp := services.CreateExpense(userid, "", expenseMetaData, totalAmount, paObj, oaObj)
	if resp.ErrorDetails != "" {
		return resp.ErrorDetails
	}
	return resp.ExpenseId
}

func CreateUserClient(username, email, phoneNum string) string {
	userMetadata := &models.UserMetaData{
		Name:     username,
		Email:    email,
		PhoneNum: phoneNum,
	}

	resp := services.CreateUser(userMetadata)
	if resp.ErrorDetails != "" {
		return resp.ErrorDetails
	}
	return resp.UserId
}

func CreateGroupClient(groupName string, userList []string) string {
	groupMetadata := &models.GroupMetaData{
		Name: groupName,
	}
	resp := services.CreateGroup(groupMetadata, userList)
	if resp.ErrorDetails != "" {
		return resp.ErrorDetails
	}
	return resp.GroupId
}

func CreateGroupExpenseClient(userid, groupid, expenseName, title, description string, totalAmount float64, paidAmounts, oweAmounts map[string]float64) string {
	resp := services.CreateGroupExpense(userid, groupid, &models.ExpenseMetaData{
		Name:        expenseName,
		Title:       title,
		Description: description,
	}, totalAmount, createUserBalanceMapList(paidAmounts), createUserBalanceMapList(oweAmounts))

	if resp.ErrorDetails != "" {
		return resp.ErrorDetails
	}
	return resp.ExpenseId
}

func SettleGroupExpenseClient(userid, groupid string) string {
	resp := services.SettleGroupExpense(userid, groupid)
	if resp.ErrorDetails != "" {
		return resp.ErrorDetails
	}
	result := ""
	for _, p := range resp.Transactions.Payment {
		result += fmt.Sprintf("User: %s needs to pay %f  to User: %s\n", p.PayerId, p.RecieverBalance.Amount.Amount, p.RecieverBalance.UserId)
	}
	return result
}

func createUserBalanceMapList(mp map[string]float64) []*models.UserBalanceMap {
	res := make([]*models.UserBalanceMap, 0)
	for k, v := range mp {
		res = append(res, &models.UserBalanceMap{
			UserId: k,
			Amount: models.Balance{
				Currency: "INR",
				Amount:   v,
			},
		})
	}
	return res
}
