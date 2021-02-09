package covidtracker

// RunApp configures the starting and running of the application
func RunApp(fileName string) error {
	nytData, err := GetNytData()
	if err != nil {
		return err
	}

	vaccines, err := GetVaccineData()
	if err != nil {
		return err
	}

	data, err := FormatData(nytData, *vaccines)
	if err != nil {
		return err
	}

	WriteToJSON(fileName, data)

	return nil
}
