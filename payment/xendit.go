package payment

import (
	"log"
	"payment-interface/utils"

	strRand "github.com/Fermekoo/go-str-random"
	"github.com/gobeam/stringy"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/virtualaccount"
)

type Xendit struct {
	Xendit PaymentInterface
}

func NewXendit() *Xendit {
	xendit.Opt.SecretKey = utils.Conf("PAYMENT.XENDIT_SERVER_KEY")
	return &Xendit{}
}

func (x *Xendit) Pay(payloads CreateVa) (string, error) {

	createVaParams := virtualaccount.CreateFixedVAParams{
		ExternalID: strRand.RandomString(32),
		BankCode:   stringy.New(payloads.Bank).ToUpper(),
		Name:       payloads.Name,
	}

	resp, err := virtualaccount.CreateFixedVA(&createVaParams)
	if err != nil {
		log.Fatal(err)
	}

	return resp.ExternalID, nil
}

func (x *Xendit) Inquiry() (string, error) {
	return "Xendit Inquiry", nil
}

func (x *Xendit) Callback() (string, error) {
	return "Xendit Callback", nil
}
