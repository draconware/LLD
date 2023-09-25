package services

import (
	"errors"

	data_service "github.com/mastik5h/LLD/splitwise/dataServices"
	generator "github.com/mastik5h/LLD/splitwise/helpers"
	"github.com/mastik5h/LLD/splitwise/models"
)

func CreateExpense(userId, groupId string, expenseMetaData *models.ExpenseMetaData, totalAmount float64, paidAmounts []*models.UserBalanceMap, oweAmounts []*models.UserBalanceMap) models.CreateExpenseResponse {
	resp := models.CreateExpenseResponse{}
	if !ValidateUsersId([]string{userId}) {
		resp.ErrorDetails = "invalid user found adding expense."
		return resp
	}

	if err := validateExpenseMetadata(expenseMetaData); err != nil {
		resp.ErrorDetails = err.Error()
		return resp
	}

	if len(paidAmounts) != len(oweAmounts) {
		resp.ErrorDetails = "invalid data set found for balances registered with expense. Kindly check and try again."
		return resp
	}

	userBaMap := make(map[string](*models.UserBalanceMap))
	for _, pa := range paidAmounts {
		bm := &models.UserBalanceMap{
			UserId: pa.UserId,
			Amount: pa.Amount,
		}
		userBaMap[bm.UserId] = bm
	}

	for _, oa := range oweAmounts {
		if v, ok := userBaMap[oa.UserId]; !ok {
			resp.ErrorDetails = "invalid data set found for balances registered with expense. Kindly check and try again."
			return resp
		} else if v != nil {
			v.Amount.Amount = v.Amount.Amount - oa.Amount.Amount
			userBaMap[oa.UserId] = v
		}
	}

	expenseId := generator.GetUniqueIdString()
	expense := &models.Expense{
		Id:          expenseId,
		MetaData:    expenseMetaData,
		GroupId:     groupId,
		TotalAmount: totalAmount,
		UserBalance: userBaMap,
	}
	err := data_service.CreateExpense(expense)
	if err != nil {
		resp.ErrorDetails = "internal error: cannot create expense entry in database"
		return resp
	}
	resp.ExpenseId = expense.Id
	return resp
}

func GetExpenseByGroupId(groupId string) ([]*models.Expense, error) {
	group_expenses_id, err := data_service.GetExpensesByGroupId(groupId)
	if err != nil {
		return nil, err
	}
	var expenses []*models.Expense
	for _, expense_id := range group_expenses_id {
		expense, err := data_service.GetExpenseById(expense_id)
		if err != nil {
			return nil, errors.New("error fetching expenses for this expense_id: " + expense_id)
		}
		expenses = append(expenses, expense)
	}
	return expenses, nil
}

func validateExpenseMetadata(metaData *models.ExpenseMetaData) error {
	//not implemented yet
	return nil
}
