package model

type Centinela struct {
	UpdatedDate     string          `json:"UpdatedDate"`
	Hospital        string          `json:"Hospital"`
	EIN             string          `json:"EIN"`
	StandardCharges StandardCharges `json:"StandardCharges"`
}

type StandardCharges struct {
	CDM       []Procedure `json:"CDM"`
	HIMCPT    []Procedure `json:"HIMCPT"`
	DRG_ICD10 []Procedure `json:"DRG-ICD10"`
}

type Procedure struct {
	AltCodes            map[string]string  `json:"AltCodes"`
	Charge              float64            `json:"Charge"`
	DeIdentifiedMinimum float64            `json:"De-Identified Minimum"`
	DeIdentifiedMaximum float64            `json:"De-Identified Maximum"`
	InsuranceRates      map[string]float64 `json:"InsuranceRates"`
	ProcedureCode       string             `json:"ProcedureCode"`
	ProcedureName       string             `json:"ProcedureName"`
}
