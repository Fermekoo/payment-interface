package payment

type PaymentInterface interface {
	Pay(payloads CreateVa) (*ResponseVa, error)
	Inquiry(order_id string) (*ResponseVa, error)
	Callback() (string, error)
}
