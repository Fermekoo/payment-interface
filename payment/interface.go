package payment

type PaymentInterface interface {
	Pay() (string, error)
	Inquiry() (string, error)
	Callback() (string, error)
}
