package covidtracker

import (
	"encoding/json"
	"io/ioutil"
)

// WriteToJSON writes JSON data to a file
func WriteToJSON(fileName string, data *CovidData) {
	file, _ := json.Marshal(data)

	_ = ioutil.WriteFile(fileName, file, 0644)
}
