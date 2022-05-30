package domain

import (
	"encoding/json"
	"fmt"

	"github.com/derekpedersen/file-parser/model"
	"github.com/derekpedersen/file-parser/utilities"
	log "github.com/sirupsen/logrus"
)

type Centinela model.Centinela

func NewCentinela(key, datasource string) (cent Centinela, err error) {
	data, err := utilities.DownloadFile(key, datasource)
	if err != nil {
		return cent, err
	}
	if err = json.Unmarshal([]byte(data), &cent); err != nil {
		log.Error(err)
		return cent, err
	}
	cent.RemoveDuplicates()
	AltCodes = cent.AltCodes()
	return cent, nil
}

func (c *Centinela) RemoveDuplicates() {
	uniquecdm := make(map[string]bool)
	for k, v := range c.StandardCharges.CDM {
		if _, ok := uniquecdm[string(utilities.Hash(v))]; !ok {
			uniquecdm[string(utilities.Hash(v))] = true
			continue
		}
		RemoveProcedure(c.StandardCharges.CDM, k)
	}

	uniquedrg := make(map[string]bool)
	for k, v := range c.StandardCharges.DRG_ICD10 {
		if _, ok := uniquedrg[string(utilities.Hash(v))]; !ok {
			uniquedrg[string(utilities.Hash(v))] = true
			continue
		}
		RemoveProcedure(c.StandardCharges.DRG_ICD10, k)
	}

	uniquecpt := make(map[string]bool)
	for k, v := range c.StandardCharges.HIMCPT {
		if _, ok := uniquecpt[string(utilities.Hash(v))]; !ok {
			uniquecpt[string(utilities.Hash(v))] = true
			continue
		}
		RemoveProcedure(c.StandardCharges.HIMCPT, k)
	}
}

func RemoveProcedure(s []model.Procedure, i int) []model.Procedure {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func (c *Centinela) AltCodes() (codes map[string]string) {
	codes = make(map[string]string)
	for _, a := range c.StandardCharges.CDM {
		for _, v := range a.AltCodes {
			codes[v] = a.ProcedureCode
		}
	}
	for _, a := range c.StandardCharges.DRG_ICD10 {
		for _, v := range a.AltCodes {
			codes[v] = a.ProcedureCode
		}
	}
	for _, a := range c.StandardCharges.HIMCPT {
		for _, v := range a.AltCodes {
			codes[v] = a.ProcedureCode
		}
	}
	return codes
}

// lots of duplicate behavior should abstact model.Procedure out into it's own domain
func (c *Centinela) Data() (data Data, err error) {
	for _, v := range c.StandardCharges.CDM {
		proc := Procedure(v)
		if len(v.ProcedureCode) == 0 ||
			v.Charge == 0 ||
			len(v.AltCodes) == 0 ||
			len(proc.GetAltCode()) == 0 {
			continue
		}
		charge := fmt.Sprintf("%v", v.Charge)
		p := model.Data{
			ProcedureCode:     v.ProcedureCode,
			ProcedureCodeType: CodeTypes[proc.GetAltCode()],
			ProcedureName:     v.ProcedureName,
			GrossCharge:       charge,
		}

		for insurer, rate := range v.InsuranceRates {
			d := p
			d.InsurancePayerName = insurer
			r := fmt.Sprintf("%v", rate)
			d.InsuranceRate = r
			data = append(data, d)
		}

	}
	for _, v := range c.StandardCharges.DRG_ICD10 {
		proc := Procedure(v)
		if len(v.ProcedureCode) == 0 ||
			v.Charge == 0 ||
			len(v.AltCodes) == 0 ||
			len(proc.GetAltCode()) == 0 {
			continue
		}
		charge := fmt.Sprintf("%v", v.Charge)
		p := model.Data{
			ProcedureCode:     v.ProcedureCode,
			ProcedureCodeType: CodeTypes[proc.GetAltCode()],
			ProcedureName:     v.ProcedureName,
			GrossCharge:       charge,
		}

		for insurer, rate := range v.InsuranceRates {
			d := p
			d.InsurancePayerName = insurer
			r := fmt.Sprintf("%v", rate)
			d.InsuranceRate = r
			data = append(data, d)
		}

	}
	for _, v := range c.StandardCharges.HIMCPT {
		proc := Procedure(v)
		if len(v.ProcedureCode) == 0 ||
			v.Charge == 0 ||
			len(v.AltCodes) == 0 ||
			len(proc.GetAltCode()) == 0 {
			continue
		}
		charge := fmt.Sprintf("%v", v.Charge)
		p := model.Data{
			ProcedureCode:     v.ProcedureCode,
			ProcedureCodeType: CodeTypes[proc.GetAltCode()],
			ProcedureName:     v.ProcedureName,
			GrossCharge:       charge,
		}

		for insurer, rate := range v.InsuranceRates {
			d := p
			d.InsurancePayerName = insurer
			r := fmt.Sprintf("%v", rate)
			d.InsuranceRate = r
			data = append(data, d)
		}

	}

	return data, nil
}

type Procedure model.Procedure

// returns the first alt code entry
func (p *Procedure) GetAltCode() string {
	for _, v := range p.AltCodes {
		return v
	}
	return ""
}
