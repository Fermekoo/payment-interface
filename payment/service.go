package payment

import (
	"errors"
	vendor "payment-interface/entity"
)

type Payment struct {
	Payment PaymentInterface
}

func NewPayment(vendorName vendor.Vendor) (*Payment, error) {
	var payment PaymentInterface
	var err error
	switch vendorName {
	case vendor.XENDIT:
		payment = NewXendit()
	case vendor.MIDTRANS:
		payment = NewMidtrans()
	case vendor.INTRAJASA:
		payment = NewIntrajasa()
	default:
		err = errors.New("service not available")
	}
	return &Payment{payment}, err
}

func (p Payment) Pay(payloads *CreateVa) (*ResponseVa, error) {
	return p.Payment.Pay(payloads)
}

func (p Payment) Inquiry(order_id string) (*ResponseVa, error) {
	return p.Payment.Inquiry(order_id)
}

func (p Payment) Callback() (string, error) {
	return p.Payment.Callback()
}
