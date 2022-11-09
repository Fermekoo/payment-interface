package intrajasa

import "payment-interface/utils"

type SinarmasCredentials struct {
	SinarmasCredentials IntraCredentialsInterface
}

func NewSinarmas() *SinarmasCredentials {
	return &SinarmasCredentials{}
}

func (permata *SinarmasCredentials) GetMerchantId() string {
	merchant_id := utils.Conf("payment.intrajasa.sinarmas.merchant_id")

	return merchant_id
}

func (permata *SinarmasCredentials) GetSecretWord() string {
	secret_word := utils.Conf("payment.intrajasa.sinarmas.secret_word")

	return secret_word
}
