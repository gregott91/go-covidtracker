package covidtracker

import (
	"encoding/json"
	"io/ioutil"
)

func WriteToJson(fileName string, data *CovidData) {
	file, _ := json.Marshal(data)

	_ = ioutil.WriteFile(fileName, file, 0644)
}
