package io_helper

import (
	"os"
	"strings"
)

func GetFilesSubstr(dirname, substr string) ([]string, error) {
	f, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}
	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return nil, err
	}
	var csvFiles []string
	for _, file := range files {
		if strings.Contains(file.Name(), substr) {
			csvFiles = append(csvFiles, file.Name())
		}
	}
	return csvFiles, nil
}
