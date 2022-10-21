package payment

import (
	"errors"
	c "payment-interface/constants"

	"github.com/gobeam/stringy"
)

type Payment struct {
	Payment PaymentInterface
}

func NewPayment(vendor_name string) (*Payment, error) {
	var payment PaymentInterface
	var err error
	switch stringy.New(vendor_name).ToUpper() {
	case c.XENDIT:
		payment = NewXendit()
	case c.MIDTRANS:
		payment = NewMidtrans()
	case c.INTRAJASA:
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
