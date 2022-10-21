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

func (i *Intrajasa) Pay(payloads *CreateVa) (*ResponseVa, error) {
	customer_intra := intrajasa.CustomerIntra{
		CustName:         strings.ToUpper(payloads.Name),
		CustAddress1:     payloads.Address,
		CustEmail:        payloads.Email,
		CustRegisterDate: payloads.RegisterDate,
	}

	payloads_intra := &intrajasa.IntraCreateVA{
		CustomerData: customer_intra,
		TotalAmount:  strconv.FormatFloat(float64(payloads.Amount), 'f', 2, 64),
		VaType:       1, //one time va
	}

	credential := credentials.NewCredential(payloads.Bank)
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
