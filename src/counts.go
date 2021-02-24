package covidtracker

import "time"

// CountData contains all the data for a single day
type CountData struct {
	Date   time.Time
	Cases  *DataPoint
	Deaths *DataPoint
}

type rawCovidData struct {
	Date   string
	Cases  int
	Deaths int
}

// GetCountData formats the raw data
func GetCountData() (*[]CountData, error) {
	var rawData, err = retrieveCountData()
	if err != nil {
		return nil, err
	}

	var data []CountData
	for index, rawDataPoint := range *rawData {
		var prevRawData = rawCovidData{Cases: 0, Deaths: 0}
		if index > 0 {
			prevRawData = (*rawData)[index-1]
		}

		parsed, err := convertCountData(rawDataPoint, prevRawData)

		if err != nil {
			return &data, err
		}

		data = append(data, parsed)
	}

	return &data, nil
}

func retrieveCountData() (*[]rawCovidData, error) {
	var rawData []rawCovidData
	if err := DownloadDataJSON("https://disease.sh/v3/covid-19/nyt/usa", &rawData); err != nil {
		return &[]rawCovidData{}, err
	}

	return &rawData, nil
}

func convertCountData(rawData rawCovidData, prevRawData rawCovidData) (CountData, error) {
	date, err := time.Parse("2006-01-02", rawData.Date)

	if err != nil {
		return CountData{}, err
	}

	return CountData{
		Date: date,
		Cases: &DataPoint{
			TotalCount: rawData.Cases,
			NewCount:   rawData.Cases - prevRawData.Cases,
		},
		Deaths: &DataPoint{
			TotalCount: rawData.Deaths,
			NewCount:   rawData.Deaths - prevRawData.Deaths,
		},
	}, nil
}
