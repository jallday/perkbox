package utils

import (
	"bytes"
	"encoding/csv"
	"io/ioutil"
)

func readFile(path string) ([]byte, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func ReadCsvFile(path string) ([][]string, error) {
	b, err := readFile(path)
	if err != nil {
		return nil, err
	}

	csv, err := csv.NewReader(bytes.NewReader(b)).ReadAll()
	if err != nil {
		return nil, err
	}
	return csv, nil
}
