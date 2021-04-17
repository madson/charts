package pie_charts

import (
	"fmt"
	"strconv"

	"github.com/madson/charts/hash_helper"
	"github.com/madson/charts/string_helper"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func PieChartForTasks(title string, records [][]string) *charts.Pie {
	pie := charts.NewPie()
	pie.AddSeries("pie", pieDataForTasks(records))

	titleOpts := opts.Title{
		Title:    "Hours per Tasks",
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
		Center: []string{"78%", "60%"},
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

func pieDataForTasks(records [][]string) []opts.PieData {
	entries := MassagePieDataForTasks(records)

	items := make([]opts.PieData, 0)
	for _, key := range hash_helper.RankMapStringFloat(entries) {
		name := string_helper.MaxString(key, 60)
		value := fmt.Sprintf("%.2f", entries[key])
		items = append(items, opts.PieData{Name: name, Value: value})
	}

	return items
}

func MassagePieDataForTasks(records [][]string) map[string]float64 {
	results := make(map[string]float64)

	for index, record := range records {
		if index == 0 {
			continue
		}

		noteRaw := record[5]
		hoursRaw := record[6]

		var hours float64
		if h, err := strconv.ParseFloat(hoursRaw, 64); err == nil {
			hours = h
		}

		results[noteRaw] += hours
	}

	return results
}
