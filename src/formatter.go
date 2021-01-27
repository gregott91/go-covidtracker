package covidtracker

import (
	"strconv"
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
	Date             time.Time
	Cases            *DataPoint
	Deaths           *DataPoint
	Hospitalizations *DataPoint
	Tests            *DataPoint
}

// CovidData contains all the data across all days
type CovidData struct {
	DailyData     *[]DailyCovidData
	RetrievalTime time.Time
	DataTypes     []DataType
}

// FormatData formats the raw data
func FormatData(rawData *[]RawCovidData) (*CovidData, error) {
	var data []DailyCovidData
	for _, rawDataPoint := range *rawData {
		parsed, err := convertData(rawDataPoint)

		if err != nil {
			return &CovidData{}, err
		}

		data = append(data, parsed)
	}

	return &CovidData{
		DailyData:     &data,
		RetrievalTime: time.Now(),
		DataTypes: []DataType{
			{Name: "Cases", IsPositive: false},
			{Name: "Deaths", IsPositive: false},
			{Name: "Hospitalizations", IsPositive: false},
			{Name: "Tests", IsPositive: true},
		},
	}, nil
}

func convertData(rawData RawCovidData) (DailyCovidData, error) {
	date, err := time.Parse("20060102", strconv.Itoa(rawData.Date))

	if err != nil {
		return DailyCovidData{}, err
	}

	return DailyCovidData{
		Date: date,
		Cases: &DataPoint{
			TotalCount: rawData.Positive,
			NewCount:   rawData.PositiveIncrease,
		},
		Deaths: &DataPoint{
			TotalCount: rawData.Death,
			NewCount:   rawData.DeathIncrease,
		},
		Hospitalizations: &DataPoint{
			TotalCount: rawData.HospitalizedCumulative,
			NewCount:   rawData.HospitalizedIncrease,
		},
		Tests: &DataPoint{
			TotalCount: rawData.TotalTestResults,
			NewCount:   rawData.TotalTestResultsIncrease,
		},
	}, nil
}
