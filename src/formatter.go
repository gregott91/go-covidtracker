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
	Date     time.Time
	Cases    *DataPoint
	Deaths   *DataPoint
	Vaccines *DataPoint
}

// CovidData contains all the data across all days
type CovidData struct {
	DailyData     *[]DailyCovidData
	RetrievalTime time.Time
	DataTypes     []DataType
}

// FormatData formats the raw data
func FormatData(covidData *[]NytData, vaccines map[time.Time]int) (*CovidData, error) {
	var data []DailyCovidData
	prevVaccineCount := 0
	for _, rawDataPoint := range *covidData {
		parsed, err := convertData(rawDataPoint, prevVaccineCount, vaccines)

		prevVaccineCount = parsed.Vaccines.TotalCount

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
			{Name: "Vaccines", IsPositive: true},
		},
	}, nil
}

func reverse(data []DailyCovidData) *[]DailyCovidData {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
	return &data
}

func convertData(rawData NytData, prevVaccineCount int, vaccines map[time.Time]int) (DailyCovidData, error) {
	totalVaccines := vaccines[rawData.Date]

	return DailyCovidData{
		Date: rawData.Date,
		Cases: &DataPoint{
			TotalCount: rawData.Cases.TotalCount,
			NewCount:   rawData.Cases.NewCount,
		},
		Deaths: &DataPoint{
			TotalCount: rawData.Deaths.TotalCount,
			NewCount:   rawData.Deaths.NewCount,
		},
		Vaccines: &DataPoint{
			TotalCount: totalVaccines,
			NewCount:   totalVaccines - prevVaccineCount,
		},
	}, nil
}
