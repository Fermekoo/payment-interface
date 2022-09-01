package credentials

import "payment-interface/utils"

type BcaCredentials struct {
	BcaCredentials IntraCredentialsInterface
}

func NewBCA() *BcaCredentials {
	return &BcaCredentials{}
}

func (bri *BcaCredentials) GetMerchantId() string {
	merchant_id := utils.Conf("payment.intrajasa.bca.merchant_id")

	return merchant_id
}

func (bri *BcaCredentials) GetSecretWord() string {
	secret_word := utils.Conf("payment.intrajasa.bca.secret_word")

	return secret_word
}
