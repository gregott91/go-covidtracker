package covidtracker

import (
	"strconv"
	"time"
)

// CountData contains all the data for a single day
type CountData struct {
	Date   time.Time
	Cases  *DataPoint
	Deaths *DataPoint
}

// GetCountData formats the raw data
func GetCountData() (*[]CountData, error) {
	var rawData, err = retrieveCountData()
	if err != nil {
		return nil, err
	}

	rawData = FilterCsv(rawData, "United States", 1)
	rawData = ReduceCsvColumns(rawData, []int{0, 2, 3, 4, 5})

	var data []CountData
	for _, element := range rawData {
		date, err := time.Parse("2006-01-02", element[0])

		if err != nil {
			return &data, err
		}

		newCases := safeParseIntField(element[1])
		newDeaths := safeParseIntField(element[2])
		totalCases := safeParseIntField(element[3])
		totalDeaths := safeParseIntField(element[4])

		dataPoint := CountData{
			Date: date,
			Cases: &DataPoint{
				TotalCount: totalCases,
				NewCount:   newCases,
			},
			Deaths: &DataPoint{
				TotalCount: totalDeaths,
				NewCount:   newDeaths,
			},
		}

		data = append(data, dataPoint)
	}

	return &data, nil
}

func safeParseIntField(field string) int {
	f, err := strconv.ParseFloat(field, 64)
	if err != nil {
		f = 0
	}

	return int(f)
}

func retrieveCountData() ([][]string, error) {
	records, err := DownloadDataCsv("https://raw.githubusercontent.com/owid/covid-19-data/master/public/data/jhu/full_data.csv")
	if err != nil {
		return nil, err
	}

	return records, nil
}
