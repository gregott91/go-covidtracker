package covidtracker

import (
	"strconv"
	"time"
)

// GetVaccineData formats the raw data
func GetVaccineData() (*map[time.Time]int, error) {
	var rawData, err = retrieveVaccineData()
	if err != nil {
		return nil, err
	}

	rawData = FilterCsv(rawData, "United States", 1)
	rawData = ReduceCsvColumns(rawData, []int{0, 2})

	previous := 0
	converted := make(map[time.Time]int)
	for _, element := range rawData {
		date, err := time.Parse("2006-01-02", element[0])

		if err != nil {
			return &converted, err
		}

		f, err := strconv.ParseFloat(element[1], 64)
		if err != nil {
			f = 0
		}
		intF := int(f)

		if intF == 0 {
			intF = previous
		}

		converted[date] = intF
		previous = intF
	}

	return &converted, nil
}

func retrieveVaccineData() ([][]string, error) {
	records, err := DownloadDataCsv("https://raw.githubusercontent.com/owid/covid-19-data/master/public/data/vaccinations/us_state_vaccinations.csv")
	if err != nil {
		return nil, err
	}

	return records, nil
}
