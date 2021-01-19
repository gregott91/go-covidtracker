package covidtracker

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func DownloadData(url string, dataType interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	respByte := buf.Bytes()
	if err := json.Unmarshal(respByte, dataType); err != nil {
		return err
	}

	return nil
}
