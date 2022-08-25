package intrajasa

type IntraCreateVA struct {
	MerchantId      string        `json:"merchantId"`
	MerchantRefCode string        `json:"merchantRefCode"`
	CustomerData    CustomerIntra `json:"customerData"`
	TotalAmount     string        `json:"totalAmount"`
	VaType          int           `json:"vaType"`
	SecureCode      string        `json:"secureCode"`
}
