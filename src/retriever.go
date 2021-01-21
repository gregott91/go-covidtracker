package covidtracker

import (
	"strconv"
	"time"
)

type rawCovidData struct {
	Date             int
	Positive         int
	Negative         int
	Hospitalized     int
	Death            int
	PositiveIncrease int
	DeathIncrease    int
}

type DataPoint struct {
	NewCount   int
	TotalCount int
}

type DailyCovidData struct {
	Date   time.Time
	Deaths *DataPoint
	Cases  *DataPoint
}

type CovidData struct {
	DailyData     *[]DailyCovidData
	RetrievalTime time.Time
}

func RetrieveData() (*CovidData, error) {
	var rawData []rawCovidData
	if err := DownloadData("https://api.covidtracking.com/v1/us/daily.json", &rawData); err != nil {
		return &CovidData{}, err
	}

	var data []DailyCovidData
	for _, rawDataPoint := range rawData {
		parsed, err := convertData(rawDataPoint)

		if err != nil {
			return &CovidData{}, err
		}

		data = append(data, parsed)
	}

	return &CovidData{
		DailyData:     &data,
		RetrievalTime: time.Now(),
	}, nil
}

func convertData(rawData rawCovidData) (DailyCovidData, error) {
	date, err := time.Parse("20060102", strconv.Itoa(rawData.Date))

	if err != nil {
		return DailyCovidData{}, err
	}

	return DailyCovidData{
		Date: date,
		Deaths: &DataPoint{
			TotalCount: rawData.Death,
			NewCount:   rawData.DeathIncrease,
		},
		Cases: &DataPoint{
			TotalCount: rawData.Positive,
			NewCount:   rawData.PositiveIncrease,
		},
	}, nil
}
