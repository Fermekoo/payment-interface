package payment

type Payment struct {
	Payment PaymentInterface
}

func NewPayment(payment PaymentInterface) *Payment {
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
