package models

import (
	"time"
)

type Group struct {
	Id       string
	Metadata *GroupMetaData
	Users    []string
}

type GroupMetaData struct {
	Name       string
	Desc       string
	Title      string
	CreatedAt  time.Time
	ModifiedAt time.Time
}

type CreateGroupResponse struct {
	GroupId      string
	ErrorDetails string
}

type CreateGroupExpenseResponse struct {
	ExpenseId    string
	ErrorDetails string
}

type SettleGroupExpenseResponse struct {
	Transactions *PaymentGraph
	ErrorDetails string
}
