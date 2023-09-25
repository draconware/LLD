package models

type Expense struct {
	Id          string
	MetaData    *ExpenseMetaData
	GroupId     string
	TotalAmount float64
	UserBalance map[string]*UserBalanceMap
}

type ExpenseMetaData struct {
	Name        string
	Title       string
	Description string
}

type UserBalanceMap struct {
	UserId string
	Amount Balance
}

type Balance struct {
	Currency string
	Amount   float64
}

type CreateExpenseResponse struct {
	ExpenseId    string
	ErrorDetails string
}

type Payment struct {
	PayerId         string
	RecieverBalance UserBalanceMap
}

type PaymentGraph struct {
	Payment []*Payment
}

func (b *Balance) SetAmount(amount float64) {
	b.Amount = amount
}
