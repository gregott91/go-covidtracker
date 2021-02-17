package covidtracker

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"io"
	"net/http"
)

// DownloadDataJSON downloads JSON data and deserializes it into the passed type
func DownloadDataJSON(url string, dataType interface{}) error {
	respByte, err := httpGetBytes(url)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(respByte, dataType); err != nil {
		return err
	}

	return nil
}

// DownloadDataCsv downloads CSV data and deserializes it into the passed type
func DownloadDataCsv(url string) ([][]string, error) {
	respByte, err := httpGetBytes(url)
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(bytes.NewReader(respByte))

	records := [][]string{}
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		records = append(records, record)
	}

	return records, nil
}

func httpGetBytes(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	respByte := buf.Bytes()

	return respByte, nil
}
