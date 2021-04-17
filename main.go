package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"

	"github.com/logictake/charts/pie_charts"

	"github.com/logictake/charts/csv_helper"

	"github.com/fsnotify/fsnotify"

	echarts "github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/render"
)

var (
	ErrorCreatingDirectory = errors.New("error while creating output directory")
	ErrorGeneratingPage    = errors.New("error while generating page")
	ErrorWritingPage       = errors.New("error while writing page")
)

const (
	OutputDir = "./output/"
	InputDir  = "/Users/madson/Downloads"
)

func main() {
	if err := watchForever(); err != nil {
		log.Fatal(err)
	}
}

func watchForever() error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}

	defer func() {
		_ = watcher.Close()
	}()

	done := make(chan bool)

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}

				length := len(event.Name)
				fileExt := event.Name[length-4 : length]

				if event.Op == fsnotify.Create && fileExt == ".csv" {
					log.Println("generating report from", event.Name)
					_ = generateReport(event.Name)
					log.Println("done")
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(InputDir)
	if err != nil {
		return err
	}
	<-done

	return nil
}

func generateReport(filename string) error {
	err := os.MkdirAll(OutputDir, 0755)
	if err != nil {
		return ErrorCreatingDirectory
	}

	basename := path.Base(filename)

	page, err := GeneratePage(filename)
	if err != nil {
		return ErrorGeneratingPage
	}

	outputPath := OutputDir + strings.Replace(basename, ".csv", ".html", -1)

	err = WriteRenderer(page, outputPath)
	if err != nil {
		return ErrorWritingPage
	}

	err = exec.Command("/usr/bin/open", outputPath).Run()
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func GeneratePage(filename string) (*components.Page, error) {
	page := components.NewPage()
	page.PageTitle = path.Base(filename)

	charts, err := generateCharts(filename)
	if err != nil {
		return nil, err
	}

	for _, c := range charts {
		page.AddCharts(c)
	}

	return page, nil
}

func WriteRenderer(renderer render.Renderer, filepath string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}

	writer := io.MultiWriter(file)
	err = renderer.Render(writer)
	if err != nil {
		return err
	}

	return nil
}

func generateCharts(filename string) ([]*echarts.Pie, error) {
	chartSlice := make([]*echarts.Pie, 0)
	title := fmt.Sprintf("Generated at: %s - %s", time.Now().String(), path.Base(filename))

	records, err := csv_helper.Read(filename)
	if err != nil {
		return nil, err
	}

	chartSlice = append(chartSlice, pie_charts.PieChartForTotalHours(title, records))
	chartSlice = append(chartSlice, pie_charts.PieChartForEngineers(title, records))
	chartSlice = append(chartSlice, pie_charts.PieChartForProjects(title, records))
	chartSlice = append(chartSlice, pie_charts.PieChartForTaskTypes(title, records))
	chartSlice = append(chartSlice, pie_charts.PieChartForTasks(title, records))

	return chartSlice, nil
}
