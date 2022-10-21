package intrajasa

import "payment-interface/utils"

type PermataCredentials struct {
	PermataCredentials IntraCredentialsInterface
}

func NewPermata() *PermataCredentials {
	return &PermataCredentials{}
}

func (permata *PermataCredentials) GetMerchantId() string {
	merchant_id := utils.Conf("payment.intrajasa.permata.merchant_id")

	return merchant_id
}

func (permata *PermataCredentials) GetSecretWord() string {
	secret_word := utils.Conf("payment.intrajasa.permata.secret_word")

	return secret_word
}
