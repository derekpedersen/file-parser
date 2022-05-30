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
	AltCodes = cent.AltCodes()
	return cent, nil
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

func (c *Centinela) Data() (data Data, err error) {
	for _, v := range c.StandardCharges.CDM {
		if len(v.ProcedureCode) == 0 || v.Charge == 0 || len(v.AltCodes) == 0 {
			continue
		}
		charge := fmt.Sprintf("%v", v.Charge)
		p := model.Data{
			ProcedureCode: v.ProcedureCode,
			// ProcedureCodeType: CodeTypes[v.AltCodes[]],
			ProcedureName: v.ProcedureName,
			GrossCharge:   charge,
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
