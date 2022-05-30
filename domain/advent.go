package domain

import (
	"encoding/json"
	"fmt"

	"github.com/derekpedersen/file-parser/model"
	"github.com/derekpedersen/file-parser/utilities"
	log "github.com/sirupsen/logrus"
)

type Advent []model.Advent

// this isn't the cleanest solution, but sometimes brute forcing it just gets the job done and buys you time to revisit the problem later
func parseAdventData(data model.AdventData) (Advent, error) {
	advent := Advent{}
	for _, n := range data {
		for _, m := range n {
			a := model.Advent{}
			for k, v := range m {
				if k == "Code" {
					a.Code = fmt.Sprintf("%v", v)
					delete(m, "Code")
				}
				if k == "Description" {
					a.Description = fmt.Sprintf("%v", v)
					delete(m, "Description")
				}
				if k == "Code Type" {
					a.CodeType = fmt.Sprintf("%v", v)
					delete(m, "Code Type")
				}
				if k == "Type" {
					a.Type = fmt.Sprintf("%v", v)
					delete(m, "Type")
				}
				if k == "RevCode" {
					a.RevCode = fmt.Sprintf("%v", v)
					delete(m, "RevCode")
				}
				if k == "Gross Charge" {
					a.GrossCharge = fmt.Sprintf("%v", v)
					delete(m, "Gross Charge")
				}
				if k == "Discounted Cash Price" {
					a.DiscountedCashPrice = fmt.Sprintf("%v", v)
					delete(m, "Discounted Cash Price")
				}
				if k == "Min" {
					a.Min = fmt.Sprintf("%v", v)
					delete(m, "Min")
				}
				if k == "Max" {
					a.Max = fmt.Sprintf("%v", v)
					delete(m, "Max")
				}
				if k == "Charge Description" {
					a.ChargeDescription = fmt.Sprintf("%v", v)
					delete(m, "Charge Description")
				}
				if k == "NDC" {
					a.NDC = fmt.Sprintf("%v", v)
					delete(m, "NDC")
				}
				if k == "Specialty" {
					a.Specialty = fmt.Sprintf("%v", v)
					delete(m, "Specialty")
				}
				if k == "Payer" {
					a.Payer = fmt.Sprintf("%v", v)
					delete(m, "Payer")
				}
				if k == "Cash Price" {
					a.CashPrice = fmt.Sprintf("%v", v)
					delete(m, "Cash Price")
				}
				if k == "Deidentified Min" {
					a.DeidentifiedMin = fmt.Sprintf("%v", v)
					delete(m, "Deidentified Min")
				}
				if k == "Deidentified Max" {
					a.DeidentifiedMax = fmt.Sprintf("%v", v)
					delete(m, "Deidentified Max")
				}
				if k == "Contracted Allowed" {
					a.ContractAllowed = fmt.Sprintf("%v", v)
					delete(m, "Contracted Allowed")
				}
				if k == "Charge Code" {
					a.ChargeCode = fmt.Sprintf("%v", v)
					delete(m, "Charge Code")
				}
			}
			a.Prices = m
			advent = append(advent, a)
		}
	}
	return advent, nil
}

func NewAdvent(key, datasource string) (advent Advent, err error) {
	data, err := utilities.DownloadFile(key, datasource)
	if err != nil {
		return advent, err
	}
	var ad model.AdventData
	if err = json.Unmarshal([]byte(data), &ad); err != nil {
		log.Error(err)
		return advent, err
	}
	advent, err = parseAdventData(ad)
	if err != nil {
		return advent, err
	}
	advent.RemoveDuplicates()
	CodeTypes = advent.CodeTypes()
	return advent, nil
}

func (a *Advent) RemoveDuplicates() {
	unique := make(map[string]bool)
	for k, v := range *a {
		if _, ok := unique[string(utilities.Hash(v))]; !ok {
			unique[string(utilities.Hash(v))] = true
			continue
		}
		RemoveAdvent(*a, k)
	}
}

func RemoveAdvent(s Advent, i int) Advent {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func (a *Advent) CodeTypes() (codes map[string]string) {
	codes = make(map[string]string)
	for _, v := range *a {
		codes[v.Code] = v.CodeType
	}
	return codes
}

func (a *Advent) Data() (data Data, err error) {
	for _, v := range *a {
		if len(v.Code) == 0 || len(v.GrossCharge) == 0 {
			continue
		}

		altcode := fmt.Sprintf("%v", v.Code)
		code := AltCodes[altcode]
		if len(code) == 0 {
			// unsure if we want to exclude non-matching alt code
			continue
		}
		p := model.Data{
			ProcedureCode:     code,
			ProcedureCodeType: v.CodeType,
			ProcedureName:     v.Description,
			GrossCharge:       v.GrossCharge,
		}

		for insurer, rate := range v.Prices {
			d := p
			// need to possibly update to account for a variety of types
			if _, ok := rate.(string); ok {
				if len(rate.(string)) == 0 ||
					rate.(string) == "N/A" ||
					rate.(string) == "0" {
					continue
				}
				d.InsuranceRate = rate.(string)
			}
			d.InsurancePayerName = insurer
			data = append(data, d)
		}
	}

	return data, nil
}
