package covidtracker

import (
	"time"
)

// DailyDataPoint contains all the data for a single day
type DailyDataPoint struct {
	Date                       time.Time
	TotalCases                 int
	NewCases                   int
	TotalDeaths                int
	NewDeaths                  int
	TotalTests                 int
	NewTests                   int
	TotalVaccinations          int
	NewVaccinations            int
	TotalPeopleFullyVaccinated int
	NewPeopleFullyVaccinated   int
	TotalPeopleHospitalized    int
	NewPeopleHospitalized      int
}

type DeserializedData struct {
	USA DeserializedCountry
}

type DeserializedCountry struct {
	Data      []DeserializedDaily
	Continent string
}

type DeserializedDaily struct {
	Date                    string
	Total_cases             float32
	New_cases               float32
	Total_deaths            float32
	New_deaths              float32
	New_tests               float32
	Total_tests             float32
	Total_vaccinations      float32
	New_vaccinations        float32
	People_fully_vaccinated float32
	Hosp_patients           float32
}

// GetDailyData gets the raw data
func GetDailyData() (*[]DailyDataPoint, error) {
	var rawData, err = getOwidData()
	if err != nil {
		return nil, err
	}

	var data []DailyDataPoint
	for index, element := range rawData.USA.Data {
		date, err := time.Parse("2006-01-02", element.Date)

		if err != nil {
			return &data, err
		}

		prevFullVaccinations := 0
		currFullVaccinations := int(element.People_fully_vaccinated)

		if index > 0 {
			prevFullVaccinations = data[index-1].TotalPeopleFullyVaccinated
		}

		if currFullVaccinations == 0 {
			currFullVaccinations = prevFullVaccinations
		}

		dataPoint := DailyDataPoint{
			Date:                       date,
			TotalCases:                 int(element.Total_cases),
			NewCases:                   int(element.New_cases),
			TotalDeaths:                int(element.Total_deaths),
			NewDeaths:                  int(element.New_deaths),
			TotalTests:                 int(element.Total_tests),
			NewTests:                   int(element.New_tests),
			TotalVaccinations:          int(element.Total_vaccinations),
			NewVaccinations:            int(element.New_vaccinations),
			TotalPeopleFullyVaccinated: currFullVaccinations,
			NewPeopleFullyVaccinated:   currFullVaccinations - prevFullVaccinations,
		}

		data = append(data, dataPoint)
	}

	return &data, nil
}

func getOwidData() (DeserializedData, error) {
	var data DeserializedData
	if err := DownloadDataJSON("https://raw.githubusercontent.com/owid/covid-19-data/master/public/data/owid-covid-data.json", &data); err != nil {
		return DeserializedData{}, err
	}

	return data, nil
}
