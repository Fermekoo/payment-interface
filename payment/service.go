package payment

import "github.com/gobeam/stringy"

type Payment struct {
	Payment PaymentInterface
}

// func NewPayment(payment PaymentInterface) *Payment {
// 	return &Payment{payment}
// }

func NewPayment(vendor_name string) *Payment {
	var payment PaymentInterface
	switch stringy.New(vendor_name).ToUpper() {
	case "XENDIT":
		payment = NewXendit()
	case "MIDTRANS":
		payment = NewMidtrans()
	case "INTRAJASA":
		payment = NewIntrajasa()

	}
	return &Payment{payment}
}

func (p *Payment) Pay(payloads CreateVa) (*ResponseVa, error) {
	return p.Payment.Pay(payloads)
}

func (p *Payment) Inquiry(order_id string) (*ResponseVa, error) {
	return p.Payment.Inquiry(order_id)
}

func (p *Payment) Callback() (string, error) {
	return p.Payment.Callback()
}
