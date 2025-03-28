package models

import "time"

type PaymentStatus string

const (
	PaymentCompleted PaymentStatus = "completed"
	PaymentFailed    PaymentStatus = "failed"
	PaymentPending   PaymentStatus = "pending"
	PaymentCancelled PaymentStatus = "cancelled"
	PaymentUnPaid    PaymentStatus = "unpaid"
)

type PaymentMode string

const (
	CashPaymentMode PaymentMode = "cash"
	CardPaymentMode PaymentMode = "card"
)

type IPayment interface {
	InitiateTransaction() PaymentStatus
}

type Payment struct {
	amount float64
	status PaymentStatus
	date   time.Time
}

type CardPayment struct {
	Payment
}

func (card *CardPayment) InitiateTransaction() PaymentStatus {
	// initiate card payment
	return PaymentCompleted
}

type CashPayment struct {
	Payment
}

func (cash *CashPayment) InitiateTransaction() PaymentStatus {
	// initiate cash payment
	return PaymentCompleted
}

type PaymentFactory struct{}

func NewPaymentFactory() *PaymentFactory {
	return &PaymentFactory{}
}

func (pf *PaymentFactory) GetPaymentMethod(amount float64, paymentMode PaymentMode) IPayment {
	payment := Payment{
		amount: amount,
		status: PaymentPending,
		date:   time.Now(),
	}
	switch paymentMode {
	case CardPaymentMode:
		return &CardPayment{Payment: payment}
	case CashPaymentMode:
		return &CashPayment{Payment: payment}
	default:
		return nil
	}
}
