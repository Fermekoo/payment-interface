package intrajasa

import (
	"bytes"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"payment-interface/intrajasa/credentials"
	"payment-interface/utils"

	"github.com/google/uuid"
)

type IntraLib struct {
	Credential credentials.IntraCredentialsInterface
}
type Intra struct {
	Token string
}

var base_url = utils.Conf("payment.intrajasa.base_url_development")

func NewIntraLib(credential *credentials.Credentials) *IntraLib {
	return &IntraLib{credential}
}

func (lib *IntraLib) secretWordHash() string {
	secret_word := lib.Credential.GetSecretWord()
	sha := sha1.New()

	sha.Write([]byte(secret_word))

	encrypted := sha.Sum(nil)
	encrypted_string := fmt.Sprintf("%x", encrypted)
	return encrypted_string
}

func (lib *IntraLib) secureCodeToken(ref_code string) string {
	merchant_id := lib.Credential.GetMerchantId()
	secret_word_hash := lib.secretWordHash()
	code := merchant_id + ref_code + secret_word_hash

	hash_code := sha256.Sum256([]byte(code))

	encrypted_string := fmt.Sprintf("%x", hash_code[:])

	return encrypted_string
}

func (lib *IntraLib) SecureCodeVa(ref_code string, amount string, display_name string, token string) string {
	merchant_id := lib.Credential.GetMerchantId()
	code := merchant_id + token + ref_code + display_name + amount

	hash_code := sha256.Sum256([]byte(code))

	encrypted_string := fmt.Sprintf("%x", hash_code[:])

	return encrypted_string
}

func (lib *IntraLib) generateToken(ref_code string) (*Intra, error) {

	var data Intra
	secure_code := lib.secureCodeToken(ref_code)
	merchant_id := lib.Credential.GetMerchantId()
	url := base_url + "/vaonline/rest/json/gettoken"

	payloads := map[string]string{
		"merchantId":      merchant_id,
		"merchantRefCode": ref_code,
		"secureCode":      secure_code,
	}

	json_payloads, _ := json.Marshal(payloads)

	response, err := http.Post(url, "application/json", bytes.NewBuffer(json_payloads))

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)

	if err != nil {
		log.Fatal(err)
	}

	return &data, err
}

func (lib *IntraLib) CreateVa(payloads *IntraCreateVA) string {
	ref_code := uuid.NewString()
	token, err := lib.generateToken(ref_code)
	secure_code := lib.SecureCodeVa(ref_code, payloads.TotalAmount, payloads.CustomerData.CustName, token.Token)
	merchant_id := lib.Credential.GetMerchantId()
	request_payloads := payloads
	if err != nil {
		log.Fatal(err)
	}
	hash_token := sha256.Sum256([]byte(token.Token))
	encrypted_token := fmt.Sprintf("%x", hash_token[:])
	request_payloads.SecureCode = secure_code
	request_payloads.MerchantRefCode = ref_code
	request_payloads.MerchantId = merchant_id

	url := base_url + "/vaonline/rest/json/generateva?t=" + encrypted_token
	json_payloads, err := json.Marshal(request_payloads)

	if err != nil {
		log.Fatal(err)
	}
	response, err := http.Post(url, "application/json", bytes.NewBuffer(json_payloads))

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	b, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(b)
}
