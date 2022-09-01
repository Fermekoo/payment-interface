package credentials

type IntraCredentialsInterface interface {
	GetMerchantId() string
	GetSecretWord() string
}
