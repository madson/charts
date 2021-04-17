package pie_charts

import (
	"fmt"
	"strconv"

	"github.com/madson/charts/hash_helper"
	"github.com/madson/charts/string_helper"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func PieChartForProjects(title string, records [][]string) *charts.Pie {
	pie := charts.NewPie()
	pie.AddSeries("pie", pieDataForProjects(records))

	titleOpts := opts.Title{
		Title:    "Hours per Project",
		Subtitle: title,
	}

	labelOpts := opts.Label{
		Show:      true,
		Formatter: "{c}",
	}

	tooltipOpts := opts.Tooltip{
		Show: true,
	}

	legendOpts := opts.Legend{
		Show:   true,
		Left:   "left",
		Orient: "vertical",
		Top:    "15%",
		Align:  "left",
		TextStyle: &opts.TextStyle{
			FontSize:   12,
			FontFamily: "monospace",
		},
	}

	pieChartOpts := opts.PieChart{
		Radius: "40%",
	}

	pie.SetGlobalOptions(
		charts.WithTitleOpts(titleOpts),
		charts.WithTooltipOpts(tooltipOpts),
		charts.WithLegendOpts(legendOpts),
	)

	pie.SetSeriesOptions(
		charts.WithLabelOpts(labelOpts),
		charts.WithPieChartOpts(pieChartOpts),
	)

	return pie
}

func pieDataForProjects(records [][]string) []opts.PieData {
	entries := massagePieDataForProjects(records)

	items := make([]opts.PieData, 0)
	for _, key := range hash_helper.RankMapStringFloat(entries) {
		name := string_helper.MaxString(key, 70)
		value := fmt.Sprintf("%.2f", entries[key])
		items = append(items, opts.PieData{Name: name, Value: value})
	}

	return items
}

func massagePieDataForProjects(records [][]string) map[string]float64 {
	results := make(map[string]float64)

	for index, record := range records {
		if index == 0 {
			continue
		}

		noteRaw := record[2]
		hoursRaw := record[6]

		var hours float64
		if h, err := strconv.ParseFloat(hoursRaw, 64); err == nil {
			hours = h
		}

		results[noteRaw] += hours
	}

	return results
}
