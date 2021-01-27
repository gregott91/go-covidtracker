package covidtracker

// RawCovidData contains the raw covid datapoints
type RawCovidData struct {
	Date                     int
	Positive                 int
	Death                    int
	PositiveIncrease         int
	DeathIncrease            int
	TotalTestResults         int
	TotalTestResultsIncrease int
	HospitalizedCumulative   int
	HospitalizedIncrease     int
}

// RetrieveData retrieves the covid data from the API
func RetrieveData() (*[]RawCovidData, error) {
	var rawData []RawCovidData
	if err := DownloadData("https://api.covidtracking.com/v1/us/daily.json", &rawData); err != nil {
		return &[]RawCovidData{}, err
	}

	return &rawData, nil
}
