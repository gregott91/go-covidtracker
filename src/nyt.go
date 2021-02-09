package covidtracker

import "time"

// NytData contains all the data for a single day
type NytData struct {
	Date   time.Time
	Cases  *DataPoint
	Deaths *DataPoint
}

type rawCovidData struct {
	Date   string
	Cases  int
	Deaths int
}

// GetNytData formats the raw data
func GetNytData() (*[]NytData, error) {
	var rawData, err = retrieveNytData()
	if err != nil {
		return nil, err
	}

	var data []NytData
	for index, rawDataPoint := range *rawData {
		var prevRawData = rawCovidData{Cases: 0, Deaths: 0}
		if index > 0 {
			prevRawData = (*rawData)[index-1]
		}

		parsed, err := convertNytData(rawDataPoint, prevRawData)

		if err != nil {
			return &data, err
		}

		data = append(data, parsed)
	}

	return &data, nil
}

func retrieveNytData() (*[]rawCovidData, error) {
	var rawData []rawCovidData
	if err := DownloadData("https://disease.sh/v3/covid-19/nyt/usa", &rawData); err != nil {
		return &[]rawCovidData{}, err
	}

	return &rawData, nil
}

func convertNytData(rawData rawCovidData, prevRawData rawCovidData) (NytData, error) {
	date, err := time.Parse("2006-01-02", rawData.Date)

	if err != nil {
		return NytData{}, err
	}

	return NytData{
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
