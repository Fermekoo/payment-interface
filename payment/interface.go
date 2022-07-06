package payment

type PaymentInterface interface {
	Pay(payloads CreateVa) (*ResponseVa, error)
	Inquiry() (string, error)
	Callback() (string, error)
}
