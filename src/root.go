package covidtracker

// RunApp configures the starting and running of the application
func RunApp(fileName string) error {
	dailyData, err := GetDailyData()
	if err != nil {
		return err
	}

	data, err := FormatData(dailyData)
	if err != nil {
		return err
	}

	WriteToJSON(fileName, data)

	return nil
}
