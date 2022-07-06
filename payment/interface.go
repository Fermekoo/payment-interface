package payment

type PaymentInterface interface {
	Pay(payloads CreateVa) (string, error)
	Inquiry() (string, error)
	Callback() (string, error)
}
