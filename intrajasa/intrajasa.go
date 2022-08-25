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
	"payment-interface/utils"

	"github.com/google/uuid"
)

type Intra struct {
	Token string
}

var base_url = utils.Conf("payment.intrajasa.base_url_development")

func secretWordHash() string {
	secret_word := utils.Conf("payment.intrajasa.bri.secret_word")
	sha := sha1.New()

	sha.Write([]byte(secret_word))

	encrypted := sha.Sum(nil)
	encrypted_string := fmt.Sprintf("%x", encrypted)
	return encrypted_string
}

func secureCodeToken(ref_code string) string {
	merchant_id := utils.Conf("payment.intrajasa.bri.merchant_id")
	secret_word_hash := secretWordHash()
	code := merchant_id + ref_code + secret_word_hash

	hash_code := sha256.Sum256([]byte(code))

	encrypted_string := fmt.Sprintf("%x", hash_code[:])

	return encrypted_string
}

func SecureCodeVa(ref_code string, amount string, display_name string, token string) string {
	merchant_id := utils.Conf("payment.intrajasa.bri.merchant_id")
	code := merchant_id + token + ref_code + display_name + amount

	hash_code := sha256.Sum256([]byte(code))

	encrypted_string := fmt.Sprintf("%x", hash_code[:])

	return encrypted_string
}

func generateToken(ref_code string) (*Intra, error) {

	var data Intra
	secure_code := secureCodeToken(ref_code)
	merchant_id := utils.Conf("payment.intrajasa.bri.merchant_id")
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

func CreateVa(payloads IntraCreateVA) string {
	ref_code := uuid.NewString()
	token, err := generateToken(ref_code)
	secure_code := SecureCodeVa(ref_code, payloads.TotalAmount, payloads.CustomerData.CustName, token.Token)
	merchant_id := utils.Conf("payment.intrajasa.bri.merchant_id")
	request_payloads := &payloads
	if err != nil {
		log.Fatal(err)
	}
	hash_token := sha256.Sum256([]byte(token.Token))
	ecrypted_token := fmt.Sprintf("%x", hash_token[:])
	request_payloads.SecureCode = secure_code
	request_payloads.MerchantRefCode = ref_code
	request_payloads.MerchantId = merchant_id
	request_payloads.VaType = 1

	url := base_url + "/vaonline/rest/json/generateva?t=" + ecrypted_token
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
