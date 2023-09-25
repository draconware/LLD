package services

import (
	"container/heap"
	"errors"
	"fmt"
	"math"

	algorithm "github.com/mastik5h/LLD/splitwise/algorithms"
	"github.com/mastik5h/LLD/splitwise/models"
)

func GetPaymentGraphForGroupExpenses(expenses []*models.Expense, totalUsers int) (*models.PaymentGraph, error) {
	userBalances := make(map[string](*models.Balance), totalUsers)
	payments := make([]*models.Payment, 0)
	for _, expense := range expenses {
		if expense == nil {
			return nil, errors.New("invalid expense object found while creating payment graph for group: " + expense.GroupId)
		}
		for userid, userBalanceMap := range expense.UserBalance {
			if _, ok := userBalances[userid]; !ok {
				userBalances[userid] = &models.Balance{
					Amount: 0.0,
				}
			}
			userBalances[userid].SetAmount(userBalances[userid].Amount + userBalanceMap.Amount.Amount)
		}
	}

	positiveUserBalances, negativeUserBalances := make([]models.UserBalanceMap, 0), make([]models.UserBalanceMap, 0)
	for userId, balance := range userBalances {
		if balance.Amount > 0.0 {
			positiveUserBalances = append(positiveUserBalances, models.UserBalanceMap{
				UserId: userId,
				Amount: *balance,
			})
		} else {
			negativeUserBalances = append(negativeUserBalances, models.UserBalanceMap{
				UserId: userId,
				Amount: *balance,
			})
		}
	}

	maxHeap := algorithm.GetMaxHeap(positiveUserBalances)
	minHeap := algorithm.GetMinHeap(negativeUserBalances)

	heap.Init(maxHeap)
	heap.Init(minHeap)

	for maxHeap.Len() > 0 && minHeap.Len() > 0 {
		a1 := maxHeap.Pop().(models.UserBalanceMap)
		a2 := minHeap.Pop().(models.UserBalanceMap)
		fmt.Println("Comparing these two: ", a1, ",", a2)
		paymentAmount := math.Min(a1.Amount.Amount, math.Abs(a2.Amount.Amount))

		payment := &models.Payment{
			PayerId: a2.UserId,
			RecieverBalance: models.UserBalanceMap{
				UserId: a1.UserId,
				Amount: models.Balance{
					Currency: a1.Amount.Currency,
					Amount:   paymentAmount,
				},
			},
		}
		payments = append(payments, payment)

		if a1.Amount.Amount > paymentAmount {
			a1.Amount.Amount -= paymentAmount
			maxHeap.Push(a1)
		}
		if math.Abs(a2.Amount.Amount) > paymentAmount {
			a2.Amount.Amount += paymentAmount
			minHeap.Push(a2)
		}
	}

	paymentGraph := &models.PaymentGraph{
		Payment: payments,
	}
	return paymentGraph, nil
}
