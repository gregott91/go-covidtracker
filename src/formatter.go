package covidtracker

import (
	"time"
)

// DataType defines a type of data returned by the JSON
type DataType struct {
	Name         string
	Display      string
	IsPositive   bool
	IsCumulative bool
	ShowDecimals bool
}

// DataPoint is a single point of data for a single day
type DataPoint struct {
	NewCount   float32
	TotalCount float32
}

// DailyCovidData contains all the data for a single day
type DailyCovidData struct {
	Date             time.Time
	Cases            *DataPoint
	Deaths           *DataPoint
	AllVaccines      *DataPoint
	FullVaccines     *DataPoint
	VaccinatedPeople *DataPoint
	Tests            *DataPoint
	Hospitalizations *DataPoint
	Mortality        *DataPoint
	Positivity       *DataPoint
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
			{Name: "Cases", Display: "Cases", IsPositive: false, IsCumulative: true, ShowDecimals: false},
			{Name: "Deaths", Display: "Deaths", IsPositive: false, IsCumulative: true, ShowDecimals: false},
			{Name: "AllVaccines", Display: "All Vaccines", IsPositive: true, IsCumulative: true, ShowDecimals: false},
			{Name: "FullVaccines", Display: "Full Vaccines", IsPositive: true, IsCumulative: true, ShowDecimals: false},
			{Name: "VaccinatedPeople", Display: "People Vaccinated", IsPositive: true, IsCumulative: true, ShowDecimals: false},
			{Name: "Tests", Display: "Tests", IsPositive: true, IsCumulative: true, ShowDecimals: false},
			{Name: "Hospitalizations", Display: "Hospitalizations", IsPositive: false, IsCumulative: false, ShowDecimals: false},
			{Name: "Mortality", Display: "Mortality", IsPositive: false, IsCumulative: false, ShowDecimals: true},
			{Name: "Positivity", Display: "Positivity", IsPositive: false, IsCumulative: false, ShowDecimals: true},
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
			TotalCount: float32(rawData.TotalCases),
			NewCount:   float32(rawData.NewCases),
		},
		Deaths: &DataPoint{
			TotalCount: float32(rawData.TotalDeaths),
			NewCount:   float32(rawData.NewDeaths),
		},
		AllVaccines: &DataPoint{
			TotalCount: float32(rawData.TotalVaccinations),
			NewCount:   float32(rawData.NewVaccinations),
		},
		FullVaccines: &DataPoint{
			TotalCount: float32(rawData.TotalPeopleFullyVaccinated),
			NewCount:   float32(rawData.NewPeopleFullyVaccinated),
		},
		VaccinatedPeople: &DataPoint{
			TotalCount: float32(rawData.TotalPeopleVaccinated),
			NewCount:   float32(rawData.NewPeopleVaccinated),
		},
		Tests: &DataPoint{
			TotalCount: float32(rawData.TotalTests),
			NewCount:   float32(rawData.NewTests),
		},
		Hospitalizations: &DataPoint{
			NewCount: float32(rawData.Hospitalizations),
		},
		Mortality: &DataPoint{
			NewCount: float32(rawData.MortalityRate),
		},
		Positivity: &DataPoint{
			NewCount: float32(rawData.Positivity),
		},
	}, nil
}
