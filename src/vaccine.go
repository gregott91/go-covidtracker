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

	rawData = filterToUnitedStates(rawData)
	rawData = reduceToRelevantData(rawData)

	converted := make(map[time.Time]int)
	for _, element := range rawData {
		date, err := time.Parse("2006-01-02", element[0])

		if err != nil {
			return &converted, err
		}

		f, err := strconv.ParseFloat(element[1], 64)
		if err != nil {
			return &converted, err
		}

		converted[date] = int(f)
	}

	return &converted, nil
}

func filterToUnitedStates(data [][]string) (ret [][]string) {
	for _, s := range data {
		if s[1] == "United States" {
			ret = append(ret, s)
		}
	}
	return
}

func reduceToRelevantData(data [][]string) (ret [][]string) {
	for index, s := range data {
		reduced := []string{s[0], s[2]}
		if len(reduced[1]) <= 0 && index > 0 {
			reduced[1] = ret[index-1][1]
		}
		ret = append(ret, reduced)
	}
	return
}

func retrieveVaccineData() ([][]string, error) {
	records, err := DownloadDataCsv("https://raw.githubusercontent.com/owid/covid-19-data/master/public/data/vaccinations/us_state_vaccinations.csv")
	if err != nil {
		return nil, err
	}

	return records, nil
}
