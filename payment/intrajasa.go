package payment

import (
	"log"
	intra_credential "payment-interface/credentials/intrajasa"
	"strings"

	intra_lib "github.com/Fermekoo/intrajasa-go"
	intra_api "github.com/Fermekoo/intrajasa-go/api"
	"github.com/google/uuid"
)

type Intrajasa struct {
	Intrajasa PaymentInterface
}

func NewIntrajasa() *Intrajasa {
	return &Intrajasa{}
}

func (i *Intrajasa) Pay(payloads *CreateVa) (*ResponseVa, error) {

	customer_intra := &intra_lib.CustomerData{
		CustName:           strings.ToUpper(payloads.Name),
		CustAddress1:       payloads.Address,
		CustRegisteredDate: payloads.RegisterDate,
		CustEmail:          payloads.Email,
		CustCountryCode:    "021",
	}

	payload_intra := &intra_lib.CreateVa{
		MerchantRefCode: uuid.NewString(),
		TotalAmount:     payloads.Amount,
		CustomerData:    customer_intra,
		VaType:          2,
	}

	credential := intra_credential.NewCredential(payloads.Bank)
	intra := intra_api.NewClient(credential.GetMerchantId(), credential.GetSecretWord(), intra_lib.Sandbox)
	create_va, err := intra.CreateVa(payload_intra)

	if err != nil {
		log.Fatal(err)
	}

	if create_va.ResponseCode != "200" {
		log.Fatal(create_va.ResponseMsg)
	}

	return &ResponseVa{
		OrderID:  create_va.MerchantRefCode,
		VaNumber: create_va.VaNumber,
		Status:   "SUCCESS",
	}, nil
}

func (i *Intrajasa) Inquiry(order_id string) (*ResponseVa, error) {
	return &ResponseVa{}, nil
}

func (i *Intrajasa) Callback() (string, error) {

	return "", nil
}
