package covidtracker

// RunApp configures the starting and running of the application
func RunApp(fileName string) error {
	data, err := RetrieveData()
	if err != nil {
		return err
	}

	WriteToJson(fileName, data)

	return nil
}
