package payment

import (
	"log"
	"payment-interface/utils"
	"strconv"

	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

var mdCore coreapi.Client

type Midtrans struct {
	Midtrans PaymentInterface
}

func NewMidtrans() *Midtrans {
	mdCore.New(utils.Conf("payment.midtrans_server_key"), midtrans.Sandbox)
	return &Midtrans{}
}

func (m *Midtrans) Pay(payloads *CreateVa) (*ResponseVa, error) {

	var chargeReq *coreapi.ChargeReq

	if payloads.Bank == "gopay" {
		expiry, _ := strconv.Atoi(utils.Conf("payment.gopay.expiry_duration"))
		chargeReq = &coreapi.ChargeReq{
			PaymentType: "gopay",
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  uuid.NewString(),
				GrossAmt: int64(payloads.Amount),
			},
			Gopay: &coreapi.GopayDetails{
				EnableCallback: true,
				CallbackUrl:    utils.Conf("payment.gopay.callback_url"),
			},
			CustomerDetails: &midtrans.CustomerDetails{
				FName: payloads.Name,
				Email: payloads.Email,
				Phone: payloads.Phone,
			},
			CustomExpiry: &coreapi.CustomExpiry{
				ExpiryDuration: expiry,
				Unit:           "minute",
			},
		}
	} else {
		bank := midtrans.Bank(payloads.Bank)
		chargeReq = &coreapi.ChargeReq{
			PaymentType: coreapi.PaymentTypeBankTransfer,
			BankTransfer: &coreapi.BankTransferDetails{
				Bank: bank,
			},
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  uuid.NewString(),
				GrossAmt: int64(payloads.Amount),
			},
		}
	}

	response, err := mdCore.ChargeTransaction(chargeReq)

	if err != nil {
		log.Fatal(err)
	}

	responseVa := ResponseVa{
		OrderID: response.OrderID,
		Status:  response.TransactionStatus,
	}

	return &responseVa, nil
}

func (m *Midtrans) Inquiry(order_id string) (*ResponseVa, error) {

	var vaNum string

	response, err := mdCore.CheckTransaction(order_id)

	if err != nil {
		log.Fatal(err)
	}

	if response.PermataVaNumber != "" {
		vaNum = response.PermataVaNumber

	} else {
		vaNum = response.VaNumbers[0].VANumber
	}

	responseVa := ResponseVa{
		OrderID:  response.OrderID,
		VaNumber: vaNum,
		Status:   response.TransactionStatus,
	}

	return &responseVa, nil

}

func (m *Midtrans) Callback() (string, error) {
	return "Midtrans Callback", nil
}
