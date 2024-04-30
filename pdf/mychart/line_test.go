package mychart

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/seq"
)

func Test_Chart(t *testing.T) {
	mainSeries := chart.ContinuousSeries{
		Name:    "A test series",
		XValues: seq.Range(1.0, 100.0),
		YValues: seq.New(seq.NewRandom().WithLen(100).WithMax(100).WithMin(50)).Array(),
	}

	minSeries := &chart.MinSeries{
		Style: chart.Style{
			Show:            true,
			StrokeColor:     chart.ColorAlternateGray,
			StrokeDashArray: []float64{5.0, 5.0},
		},
		InnerSeries: mainSeries,
	}

	maxSeries := &chart.MaxSeries{
		Style: chart.Style{
			Show:            true,
			StrokeColor:     chart.ColorAlternateGray,
			StrokeDashArray: []float64{5.0, 5.0},
		},
		InnerSeries: mainSeries,
	}

	graph := chart.Chart{
		Width:  1920,
		Height: 1080,
		YAxis: chart.YAxis{
			Name:      "Random Values",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
			Range: &chart.ContinuousRange{
				Min: 25,
				Max: 175,
			},
		},
		XAxis: chart.XAxis{
			Name:      "Random Other Values",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
		},
		Series: []chart.Series{
			mainSeries,
			minSeries,
			maxSeries,
			chart.LastValueAnnotation(minSeries),
			chart.LastValueAnnotation(maxSeries),
		},
	}

	f, _ := os.Create("output.png")
	defer f.Close()
	graph.Render(chart.PNG, f)
}
func Test_Line(t *testing.T) {
	tlc := TimeLineChart{
		YAxisName:    "庫溫2 (°C)",
		NoUpperLower: false,
		UpperValue:   30,
		LowerValue:   8,
		TimestampList: []int64{
			1539734400, 1539734460, 1539734520,
			1539734580, 1539734640, 1539734700,
			1539734760, 1539734820, 1539734880,
		},
		TimeData: []TimeLine{
			{
				Name: "最大值",
				Data: map[int64]float64{
					1539734400: 25.4,
					1539734460: 23.2,
					1539734520: 23.1,
					1539734580: 22.4,
					1539734640: 24.2,
					1539734700: 19.4,
					1539734760: 17,
					1539734820: 16.1,
					1539734880: 18.2,
				},
			},
			{
				Name: "Avg",
				Data: map[int64]float64{
					1539734400: 25.3,
					1539734460: 23,
					1539734520: 22.6,
					1539734580: 21.4,
					1539734640: 23.7,
					1539734700: 18.9,
					1539734760: 16.5,
					1539734820: 15.6,
					1539734880: 17.6,
				},
			},
		},
	}

	f, _ := os.Create("line.png")
	defer f.Close()

	tlc.Draw("../../resource/TW-Medium.ttf", f)
	assert.True(t, false)
}
