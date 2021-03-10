package covidtracker

import (
	"time"
)

// DataType defines a type of data returned by the JSON
type DataType struct {
	Name       string
	IsPositive bool
}

// DataPoint is a single point of data for a single day
type DataPoint struct {
	NewCount   int
	TotalCount int
}

// DailyCovidData contains all the data for a single day
type DailyCovidData struct {
	Date         time.Time
	Cases        *DataPoint
	Deaths       *DataPoint
	AllVaccines  *DataPoint
	FullVaccines *DataPoint
	Tests        *DataPoint
}

// CovidData contains all the data across all days
type CovidData struct {
	DailyData     *[]DailyCovidData
	RetrievalTime time.Time
	DataTypes     []DataType
}

// FormatData formats the raw data
func FormatData(covidData *[]DailyDataPoint) (*CovidData, error) {
	var data []DailyCovidData
	for _, rawDataPoint := range *covidData {
		parsed, err := convertData(rawDataPoint)

		if err != nil {
			return &CovidData{}, err
		}

		data = append(data, parsed)
	}

	return &CovidData{
		DailyData:     reverse(data),
		RetrievalTime: time.Now(),
		DataTypes: []DataType{
			{Name: "Cases", IsPositive: false},
			{Name: "Deaths", IsPositive: false},
			{Name: "AllVaccines", IsPositive: true},
			{Name: "FullVaccines", IsPositive: true},
			{Name: "Tests", IsPositive: true},
		},
	}, nil
}

func reverse(data []DailyCovidData) *[]DailyCovidData {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
	return &data
}

func convertData(rawData DailyDataPoint) (DailyCovidData, error) {
	return DailyCovidData{
		Date: rawData.Date,
		Cases: &DataPoint{
			TotalCount: rawData.TotalCases,
			NewCount:   rawData.NewCases,
		},
		Deaths: &DataPoint{
			TotalCount: rawData.TotalDeaths,
			NewCount:   rawData.NewDeaths,
		},
		AllVaccines: &DataPoint{
			TotalCount: rawData.TotalVaccinations,
			NewCount:   rawData.NewVaccinations,
		},
		FullVaccines: &DataPoint{
			TotalCount: rawData.TotalPeopleFullyVaccinated,
			NewCount:   rawData.NewPeopleFullyVaccinated,
		},
		Tests: &DataPoint{
			TotalCount: rawData.TotalTests,
			NewCount:   rawData.NewTests,
		},
	}, nil
}
