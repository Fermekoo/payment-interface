package credentials

type Credentials struct {
	Credentials IntraCredentialsInterface
}

func NewCredential(credential IntraCredentialsInterface) *Credentials {
	return &Credentials{credential}
}

func (cred *Credentials) GetMerchantId() string {
	return cred.Credentials.GetMerchantId()
}

func (cred *Credentials) GetSecretWord() string {
	return cred.Credentials.GetSecretWord()
}
