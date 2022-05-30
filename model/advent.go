package model

// the advent data is in a dynamic flat format where all fields are present on one layer
type AdventData [][]map[string]interface{}

type Advent struct {
	Code                string `json:"Code"`
	Description         string `json:"Description"`
	CodeType            string `json:"Code Type"`
	Type                string `json:"Type"`
	RevCode             string `json:"RevCode"`
	GrossCharge         string `json:"Gross Charge"`
	DiscountedCashPrice string `json:"Discounted Cash Price"`
	Min                 string `json:"Min"`
	Max                 string `json:"Max"`

	ChargeDescription string `json:"ChargeDescription"`
	NDC               string `json:"NDC"`
	Specialty         string `json:"Specialty"`
	Payer             string `json:"Payer"`
	CashPrice         string `json:"CashPrice"`
	DeidentifiedMin   string `json:"Deidentified Min"`
	DeidentifiedMax   string `json:"Deidentified Max"`
	ContractAllowed   string `json:"Contracted Allowed"`

	ChargeCode string `json:"Charge Code"`

	Prices map[string]interface{} `json:"-"`
}
