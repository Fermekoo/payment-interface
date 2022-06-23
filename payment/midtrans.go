package payment

import (
	"fmt"
	"log"
	"payment-interface/utils"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

var mdCore coreapi.Client

type Midtrans struct {
	Midtrans PaymentInterface
}

func NewMidtrans() *Midtrans {
	mdCore.New(utils.Conf("PAYMENT.MIDTRANS_SERVER_KEY"), midtrans.Sandbox)
	return &Midtrans{}
}

func (m *Midtrans) Pay() (string, error) {
	chargeReq := &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeBankTransfer,
		BankTransfer: &coreapi.BankTransferDetails{
			Bank: "bca",
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "22419223",
			GrossAmt: 200000,
		},
	}
	fmt.Println(chargeReq)
	response, err := mdCore.ChargeTransaction(chargeReq)
	if err != nil {
		log.Fatal(err)
	}

	return response.TransactionID, nil
}

func (m *Midtrans) Inquiry() (string, error) {
	return "Midtrans Inquiry", nil
}

func (m *Midtrans) Callback() (string, error) {
	return "Midtrans Callback", nil
}
