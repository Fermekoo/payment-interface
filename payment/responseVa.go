package payment

type ResponseVa struct {
	OrderID  string `json:"order_id"`
	VaNumber string `json:"va_number"`
	Status   string `json:"status"`
}
