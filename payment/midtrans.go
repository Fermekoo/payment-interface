package payment

type Midtrans struct {
	Midtrans PaymentInterface
}

func NewMidtrans() *Midtrans {
	return &Midtrans{}
}

func (m *Midtrans) Pay() (string, error) {
	return "Midtrans Pay", nil
}

func (m *Midtrans) Inquiry() (string, error) {
	return "Midtrans Inquiry", nil
}

func (m *Midtrans) Callback() (string, error) {
	return "Midtrans Callback", nil
}
