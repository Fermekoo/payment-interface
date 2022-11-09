package intrajasa

import "github.com/gobeam/stringy"

type Credentials struct {
	Credentials IntraCredentialsInterface
}

func NewCredential(bank_name string) *Credentials {

	var bank IntraCredentialsInterface
	switch stringy.New(bank_name).ToUpper() {
	case "BRI":
		bank = NewBRI()
	case "BCA":
		bank = NewBCA()
	case "MANDIRI":
		bank = NewMandiri()
	case "PERMATA":
		bank = NewPermata()
	case "SINARMAS":
		bank = NewSinarmas()
	}

	return &Credentials{Credentials: bank}
}

func (cred *Credentials) GetMerchantId() string {
	return cred.Credentials.GetMerchantId()
}

func (cred *Credentials) GetSecretWord() string {
	return cred.Credentials.GetSecretWord()
}
