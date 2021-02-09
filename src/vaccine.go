package covidtracker

import (
	"time"
)

type rawVaccineData struct {
	Timeline map[string]int
}

// GetVaccineData formats the raw data
func GetVaccineData() (*map[time.Time]int, error) {
	var rawData, err = retrieveVaccineData()
	if err != nil {
		return nil, err
	}

	converted := make(map[time.Time]int)
	for key, element := range rawData.Timeline {
		date, err := time.Parse("1/2/06", key)

		if err != nil {
			return &converted, err
		}

		converted[date] = element
	}

	return &converted, nil
}

func retrieveVaccineData() (*rawVaccineData, error) {
	var rawData rawVaccineData
	if err := DownloadData("https://disease.sh/v3/covid-19/vaccine/coverage/countries/usa?lastdays=all", &rawData); err != nil {
		return &rawVaccineData{}, err
	}

	return &rawData, nil
}
