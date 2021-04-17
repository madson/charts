package main

import (
	"errors"
	"testing"

	"github.com/madson/charts/csv_helper"
	"github.com/madson/charts/io_helper"
	"github.com/madson/charts/pie_charts"
)

func Test_pieData(t *testing.T) {
	t.Run("should not return an error", func(t *testing.T) {
		filename, _ := getFilename()
		records, _ := csv_helper.Read(filename)
		data := pie_charts.MassagePieDataForTasks(records)
		if data == nil {
			t.Error("data can't be nil")
		}
	})

	t.Run("should load items", func(t *testing.T) {
		filename, _ := getFilename()
		records, _ := csv_helper.Read(filename)
		items := pie_charts.MassagePieDataForTasks(records)

		if len(items) < 1 {
			t.Error("it should return a non-empty slice")
		}
	})
}

func getFilename() (string, error) {
	var filename string

	filenames, err := io_helper.GetFilesSubstr(".", ".csv")
	if err != nil {
		return filename, err
	}
	if len(filenames) < 1 {
		return filename, errors.New("empty filenames slice")
	}

	filename = filenames[0]
	return filename, nil
}
