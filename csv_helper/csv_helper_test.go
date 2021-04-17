package csv_helper

import (
	"log"
	"testing"

	"github.com/logictake/charts/io_helper"
)

func Test_GetFilesSubstr(t *testing.T) {

	t.Run("should not return error", func(t *testing.T) {
		if _, err := io_helper.GetFilesSubstr(".", ".csv"); err != nil {
			t.Errorf("retuned an error: %s", err.Error())
		}
	})

	t.Run("should return an empty slice", func(t *testing.T) {
		files, _ := io_helper.GetFilesSubstr(".", ".nonexistent-filename")
		if len(files) > 0 {
			t.Error("it should not return any filename")
		}
	})
}

func Test_getRecordsFromFile(t *testing.T) {
	t.Run("should return an error for an nonexistent file path", func(t *testing.T) {
		if _, err := Read("./nonexistent-filename.csv"); err == nil {
			t.Error("an error should occur")
		}
	})

	t.Run("should return a slice of records", func(t *testing.T) {
		filenames, _ := io_helper.GetFilesSubstr(".", ".csv")
		if len(filenames) <= 0 {
			t.Errorf("couldn't find a file named: %s", filenames[0])
		}

		records, err := Read(filenames[0])
		if err != nil {
			t.Errorf("error happened: %s", err.Error())
		}

		if len(records) <= 0 {
			t.Error("got an empty slice")
		}

		log.Print(records)
	})
}

//func Test_generatePieItems(t *testing.T) {
//	tests := []struct {
//		name string
//		want []opts.PieData
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := pieDataForTasks(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("pieDataForTasks() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_pieWithLabel(t *testing.T) {
//	tests := []struct {
//		name string
//		want *pie_charts.Pie
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := pieChartForTasks(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("pieChartForTasks() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestGeneratePage(t *testing.T) {
//	tests := []struct {
//		name string
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//		})
//	}
//}
