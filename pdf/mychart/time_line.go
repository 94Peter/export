package mychart

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/94peter/export/pdf/style"

	"github.com/golang/freetype/truetype"
	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/drawing"
)

type TimeLineChart struct {
	TimestampList []int64
	TimeData      []TimeLine
	UpperValue    float64
	LowerValue    float64
	YAxisName     string
	NoUpperLower  bool
	ShowCI        bool
	DP            uint // 小數位數

	Width  int
	Height int

	max float64
	min float64
}

type TimeLine struct {
	Name  string
	Data  map[int64]float64
	Color style.Color
}

func (tl *TimeLine) toDrawColor() drawing.Color {
	return drawing.Color{
		R: tl.Color.R,
		G: tl.Color.G,
		B: tl.Color.B,
		A: tl.Color.A,
	}
}

func (tlc *TimeLineChart) getTimeSeries() []chart.Series {
	timeLen := len(tlc.TimestampList)
	if timeLen <= 1 {
		return nil
	}
	timeAry := make([]time.Time, timeLen)
	for i := 0; i < timeLen; i++ {
		timeAry[i] = time.Unix(tlc.TimestampList[i], 0)
	}

	dataLen := len(tlc.TimeData)
	timeSeries := make([]chart.Series, dataLen)
	var value float64
	var ok bool
	var timestamp int64
	for i := 0; i < dataLen; i++ {
		yValeAry := make([]float64, timeLen)
		for j := 0; j < timeLen; j++ {
			timestamp = tlc.TimestampList[j]
			value, ok = tlc.TimeData[i].Data[timestamp]
			if ok {
				yValeAry[j] = value
			} else {
				yValeAry[j] = 0
			}
			if i == 0 && j == 0 {
				tlc.max = value
				tlc.min = value
			} else {
				if value > tlc.max {
					tlc.max = value
				}
				if value < tlc.min {
					tlc.min = value
				}
			}
		}
		timeSeries[i] = chart.TimeSeries{
			Name:    tlc.TimeData[i].Name,
			XValues: timeAry,
			YValues: yValeAry,
			Style: chart.Style{
				Show:        true,
				StrokeColor: tlc.TimeData[i].toDrawColor(),
			},
		}
	}
	return timeSeries
}

func (tlc *TimeLineChart) getUpperLowerSeries() []chart.Series {
	minSeries := &chart.MinSeries{
		Name: "Lower Limit",
		Style: chart.Style{
			Show:            true,
			StrokeColor:     chart.ColorAlternateGray,
			StrokeDashArray: []float64{5.0, 5.0},
		},
		InnerSeries: tlc.getLowerSeries(),
	}
	maxSeries := &chart.MinSeries{

		Name: "Upper Limit",
		Style: chart.Style{
			Show:            true,
			StrokeColor:     chart.ColorAlternateGray,
			StrokeDashArray: []float64{5.0, 5.0},
		},
		InnerSeries: tlc.getUpperSeries(),
	}
	return []chart.Series{
		minSeries,
		chart.LastValueAnnotation(minSeries),
		maxSeries,
		chart.LastValueAnnotation(maxSeries),
	}
}

func (tlc *TimeLineChart) getCISeries() []chart.Series {
	minSeries := &chart.MinSeries{
		Name: "-1.96D",
		Style: chart.Style{
			Show:            true,
			StrokeColor:     chart.ColorRed,
			StrokeDashArray: []float64{5.0, 5.0},
		},
		InnerSeries: tlc.getLowerSeries(),
	}
	maxSeries := &chart.MinSeries{

		Name: "1.96D",
		Style: chart.Style{
			Show:            true,
			StrokeColor:     chart.ColorRed,
			StrokeDashArray: []float64{5.0, 5.0},
		},
		InnerSeries: tlc.getUpperSeries(),
	}
	return []chart.Series{
		minSeries,
		chart.LastValueAnnotation(minSeries),
		maxSeries,
		chart.LastValueAnnotation(maxSeries),
	}
}

func readFont(fontfile string) *truetype.Font {
	// 讀字體
	fontBytes, err := os.ReadFile(fontfile)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	font, err := truetype.Parse(fontBytes)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return font
}

func (tlc *TimeLineChart) Draw(fontfile string, ioWriter io.Writer) {
	timeSeries := tlc.getTimeSeries()
	if timeSeries == nil {
		return
	}
	if tlc.min == tlc.max {
		tlc.min--
		tlc.max++
	}
	if !tlc.NoUpperLower {
		timeSeries = append(timeSeries, tlc.getUpperLowerSeries()...)
		if tlc.LowerValue < tlc.min {
			tlc.min = tlc.LowerValue
		}
		if tlc.UpperValue > tlc.max {
			tlc.max = tlc.UpperValue
		}
	} else if tlc.ShowCI {
		timeSeries = append(timeSeries, tlc.getCISeries()...)
		if tlc.LowerValue < tlc.min {
			tlc.min = tlc.LowerValue
		}
		if tlc.UpperValue > tlc.max {
			tlc.max = tlc.UpperValue
		}
	}
	dateFormat := "01-02 15:04"

	if tlc.DP == 0 {
		tlc.DP = 1
	}
	if tlc.Width == 0 {
		tlc.Width = 1000
	}
	if tlc.Height == 0 {
		tlc.Height = 600
	}
	if len(timeSeries) == 0 {
		return
	}

	//ns := chart.StyleShow()
	graph := chart.Chart{
		Width:  tlc.Width,
		Height: tlc.Height,
		Background: chart.Style{
			Padding: chart.Box{
				Top: 50,
			},
		},
		XAxis: chart.XAxis{
			Name: "Time",
			//NameStyle: ns,
			Style: chart.Style{
				Show:     true,
				FontSize: 8,
				TextWrap: chart.TextWrapWord,
			},
			ValueFormatter: func(v interface{}) string {
				if typed, isTyped := v.(time.Time); isTyped {
					return typed.Format(dateFormat)
				}
				if typed, isTyped := v.(int64); isTyped {
					return time.Unix(0, typed).Format(dateFormat)
				}
				if typed, isTyped := v.(float64); isTyped {
					return time.Unix(0, int64(typed)).Format(dateFormat)
				}
				return ""
			},
		},
		YAxis: chart.YAxis{
			Name: tlc.YAxisName,
			//NameStyle: ns,
			Style: chart.Style{
				Show: true,
				Font: readFont(fontfile),
			},
			ValueFormatter: func(v interface{}) string {
				if typed, isTyped := v.(float64); isTyped {
					format := fmt.Sprintf("%%.%df", tlc.DP)
					return fmt.Sprintf(format, typed)
				}
				return ""
			},
			Range: &chart.ContinuousRange{
				Min: tlc.min,
				Max: tlc.max,
			},
			// Range: &chart.ContinuousRange{
			// 	Min: tlc.LowerValue - 10,
			// 	Max: tlc.UpperValue + 10,
			// },
		},
		Series: timeSeries,
	}

	graph.Elements = []chart.Renderable{chart.LegendThin(&graph, chart.Style{
		FontSize: 16,
	})}
	graph.Font = readFont(fontfile)
	err := graph.Render(chart.PNG, ioWriter)
	if err != nil {
		panic(err)
	}

}

func (tlc *TimeLineChart) getLowerSeries() upperLowerSeries {
	return upperLowerSeries{
		TimestampList: tlc.TimestampList,
		Value:         tlc.LowerValue,
	}
}

func (tlc *TimeLineChart) getUpperSeries() upperLowerSeries {
	return upperLowerSeries{
		TimestampList: tlc.TimestampList,
		Value:         tlc.UpperValue,
	}
}

type upperLowerSeries struct {
	TimestampList []int64
	Value         float64
}

func (tls upperLowerSeries) Len() int {
	return len(tls.TimestampList)
}

func (tls upperLowerSeries) GetValues(index int) (x, y float64) {
	if index == -1 {
		return
	}
	if len(tls.TimestampList) <= index {
		return
	}
	x = float64(time.Unix(tls.TimestampList[index], 0).UnixNano())
	y = tls.Value
	return
}
