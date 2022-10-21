package intrajasa

import "payment-interface/utils"

type MandiriCredentials struct {
	MandiriCredentials IntraCredentialsInterface
}

func NewMandiri() *MandiriCredentials {
	return &MandiriCredentials{}
}

func (mandiri *MandiriCredentials) GetMerchantId() string {
	merchant_id := utils.Conf("payment.intrajasa.mandiri.merchant_id")

	return merchant_id
}

func (mandiri *MandiriCredentials) GetSecretWord() string {
	secret_word := utils.Conf("payment.intrajasa.mandiri.secret_word")

	return secret_word
}
