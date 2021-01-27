package covidtracker

// RunApp configures the starting and running of the application
func RunApp(fileName string) error {
	rawData, err := RetrieveData()
	if err != nil {
		return err
	}

	data, err := FormatData(rawData)
	if err != nil {
		return err
	}

	WriteToJSON(fileName, data)

	return nil
}
