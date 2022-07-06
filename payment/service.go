package payment

type Payment struct {
	Payment PaymentInterface
}

func NewPayment(payment PaymentInterface) *Payment {
	return &Payment{payment}
}

func (p *Payment) Pay(payloads CreateVa) (string, error) {
	return p.Payment.Pay(payloads)
}

func (p *Payment) Inquiry() (string, error) {
	return p.Payment.Inquiry()
}

func (p *Payment) Callback() (string, error) {
	return p.Payment.Callback()
}
