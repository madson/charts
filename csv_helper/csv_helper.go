package csv_helper

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
)

func Read(filepath string) ([][]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, errors.New("error trying to open file: " + err.Error())
	}

	reader := csv.NewReader(file)
	var records [][]string

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, errors.New("error trying to read file: " + err.Error())
		}
		records = append(records, record)
	}

	return records, nil
}
