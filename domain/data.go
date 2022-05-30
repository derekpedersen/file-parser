package domain

import (
	"github.com/derekpedersen/file-parser/model"
)

const (
	ADVENT    = "ADVENT"
	CENTINELA = "CENTINELA"
)

var AltCodes map[string]string
var CodeTypes map[string]string

type Data []model.Data

func (d *Data) GeneratePriceCsv() (csv [][]string) {
	headers := []string{
		"Procedure Code",
		"Procedure Code Type",
		"Procedure Name",
		"Gross Charge",
		"Insurance Payer Namer",
		"Insurance Rate",
	}

	csv = append(csv, headers)

	for _, v := range *d {
		var row []string
		row = append(row, v.ProcedureCode)
		row = append(row, v.ProcedureCodeType)
		row = append(row, v.ProcedureName)
		row = append(row, v.GrossCharge)
		row = append(row, v.InsurancePayerName)
		row = append(row, v.InsuranceRate)
		csv = append(csv, row)
	}

	return csv
}
