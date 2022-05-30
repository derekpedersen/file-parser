package main

import (
	"github.com/derekpedersen/file-parser/domain"
	"github.com/derekpedersen/file-parser/utilities"
	log "github.com/sirupsen/logrus"
)

var datasources = map[string]string{
	domain.CENTINELA: "https://www.centinelamed.com/261150758_CentinelaHospitalMedicalCenter_standardcharges.json",
	// "Centinela": "./samples/centinela.json",
	domain.ADVENT: "https://www.adventhealth.com/sites/default/files/CDM/2022/480637331_AdventHealthShawneeMission_standardcharges.json",
}

func main() {
	var centinela domain.Centinela
	var advent domain.Advent
	var err error

	// extract
	for k, v := range datasources {
		if k == domain.ADVENT {
			advent, err = domain.NewAdvent(k, v)
			if err != nil {
				log.Fatal(err)
			}
		}
		if k == domain.CENTINELA {
			centinela, err = domain.NewCentinela(k, v)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	// transform
	adventData, err := advent.Data()
	if err != nil {
		log.Fatal(err)
	}
	centinelaData, err := centinela.Data()
	if err != nil {
		log.Fatal(err)
	}

	// load (save data to file)
	utilities.Write(domain.ADVENT, adventData.GeneratePriceCsv())
	utilities.Write(domain.CENTINELA, centinelaData.GeneratePriceCsv())
}
