package payment

import (
	"log"
	"payment-interface/utils"

	strRand "github.com/Fermekoo/go-str-random"
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

func (x *Xendit) Pay() (string, error) {

	banks, err := virtualaccount.GetAvailableBanks()
	if err != nil {
		log.Fatal(err)
	}

	createVaParams := virtualaccount.CreateFixedVAParams{
		ExternalID: strRand.RandomString(32),
		BankCode:   banks[0].Code,
		Name:       "Dandi Fermeko",
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
