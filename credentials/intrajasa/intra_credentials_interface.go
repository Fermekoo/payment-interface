package intrajasa

type IntraCredentialsInterface interface {
	GetMerchantId() string
	GetSecretWord() string
}
