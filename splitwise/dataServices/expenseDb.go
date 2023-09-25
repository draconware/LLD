package data_service

import (
	"errors"

	"github.com/mastik5h/LLD/splitwise/models"
)

var expenseDb = make(map[string]*models.Expense, 0)
var groupExpenseDb = make(map[string][]string, 0)

func CreateExpense(expense *models.Expense) error {
	if expense == nil {
		return errors.New("invalid expense object found.")
	}
	if _, ok := expenseDb[expense.Id]; ok {
		return errors.New("entry already exists")
	} else {
		expenseDb[expense.Id] = expense
	}

	if expense.GroupId != "" {
		groupExpenseDb[expense.GroupId] = append(groupExpenseDb[expense.GroupId], expense.Id)
	}
	return nil
}

func GetExpenseById(expenseId string) (*models.Expense, error) {
	if expense, ok := expenseDb[expenseId]; !ok {
		return nil, errors.New("no expense exists with this expense id.")
	} else {
		return expense, nil
	}
}

func GetExpensesByGroupId(groupId string) ([]string, error) {
	if expensesList, ok := groupExpenseDb[groupId]; !ok {
		return nil, errors.New("no expenses registered with this group id.")
	} else {
		return expensesList, nil
	}
}
