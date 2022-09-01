package payment

import (
	"log"
	"payment-interface/utils"

	"github.com/gobeam/stringy"
	"github.com/google/uuid"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/virtualaccount"
)

type Xendit struct {
	Xendit PaymentInterface
}

func NewXendit() *Xendit {
	xendit.Opt.SecretKey = utils.Conf("payment.xendit_server_key")
	return &Xendit{}
}

func (x *Xendit) Pay(payloads CreateVa) (*ResponseVa, error) {

	createVaParams := virtualaccount.CreateFixedVAParams{
		ExternalID: uuid.NewString(),
		BankCode:   stringy.New(payloads.Bank).ToUpper(),
		Name:       payloads.Name,
	}

	response, err := virtualaccount.CreateFixedVA(&createVaParams)
	if err != nil {
		log.Fatal(err)
	}

	responseVa := ResponseVa{
		OrderID:  response.ExternalID,
		VaNumber: response.AccountNumber,
		Status:   response.Status,
	}

	return &responseVa, nil
}

func (x *Xendit) Inquiry(order_id string) (*ResponseVa, error) {
	return &ResponseVa{}, nil
}

func (x *Xendit) Callback() (string, error) {
	return "Xendit Callback", nil
}
