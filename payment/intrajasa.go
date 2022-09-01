package payment

import (
	"payment-interface/intrajasa"
	"payment-interface/intrajasa/credentials"
	"strconv"
	"strings"
)

type Intrajasa struct {
	Intrajasa PaymentInterface
}

func NewIntrajasa() *Intrajasa {
	return &Intrajasa{}
}

func (i *Intrajasa) Pay(payloads CreateVa) (*ResponseVa, error) {
	customer_intra := intrajasa.CustomerIntra{
		CustName:         strings.ToUpper("INDODAX Dandi Fermeko"),
		CustAddress1:     "Jakarta",
		CustEmail:        "dandi.fermeko@bitcoin.co.id",
		CustRegisterDate: "2020-01-01",
	}

	payloads_intra := intrajasa.IntraCreateVA{
		CustomerData: customer_intra,
		TotalAmount:  strconv.FormatFloat(100000, 'f', 2, 64),
	}
	bank := credentials.NewBCA()
	credential := credentials.NewCredential(bank)
	intra_lib := intrajasa.NewIntraLib(credential)
	create_va := intra_lib.CreateVa(payloads_intra)
	return &ResponseVa{
		OrderID: create_va,
	}, nil
}

func (i *Intrajasa) Inquiry(order_id string) (*ResponseVa, error) {
	return &ResponseVa{}, nil
}

func (i *Intrajasa) Callback() (string, error) {

	return "", nil
}
