package credentials

import "payment-interface/utils"

type BriCredentials struct {
	BriCredentials IntraCredentialsInterface
}

func NewBRI() *BriCredentials {
	return &BriCredentials{}
}

func (bri *BriCredentials) GetMerchantId() string {
	merchant_id := utils.Conf("payment.intrajasa.bri.merchant_id")

	return merchant_id
}

func (bri *BriCredentials) GetSecretWord() string {
	secret_word := utils.Conf("payment.intrajasa.bri.secret_word")

	return secret_word
}
