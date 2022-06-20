package payment

type Xendit struct {
	Xendit PaymentInterface
}

func NewXendit() *Xendit {
	return &Xendit{}
}

func (x *Xendit) Pay() (string, error) {
	return "Xendit Pay", nil
}

func (x *Xendit) Inquiry() (string, error) {
	return "Xendit Inquiry", nil
}

func (x *Xendit) Callback() (string, error) {
	return "Xendit Callback", nil
}
