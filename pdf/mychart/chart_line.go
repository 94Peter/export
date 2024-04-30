package mychart

import (
	"math/rand"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

type ChartLine struct {
	ChartA []float64
	ChartB []float64
	Line1  []float64
	Line2  []float64
	Line3  []float64
	Line4  []float64
}

func (cl *ChartLine) CreatePNG() error {
	groupA := plotter.Values(cl.ChartA)
	groupB := plotter.Values(cl.ChartB)

	p := plot.New()

	p.Title.Text = "Bar chart"
	p.Y.Label.Text = "Heights"

	w := vg.Points(20)

	barsA, err := plotter.NewBarChart(groupA, w)
	if err != nil {
		return err
	}
	barsA.LineStyle.Width = vg.Length(0)
	barsA.Color = plotutil.Color(3)
	barsA.Offset = -w / 2

	barsB, err := plotter.NewBarChart(groupB, w)
	if err != nil {
		return err
	}
	barsB.LineStyle.Width = vg.Length(0)
	barsB.Color = plotutil.Color(4)
	barsB.Offset = w / 2

	p.Add(barsA, barsB)
	p.Legend.Add("Group A", barsA)
	p.Legend.Add("Group B", barsB)
	p.Legend.Top = true
	p.NominalX("One", "Two", "Three", "Four", "Five")

	plotutil.AddLinePoints(p,
		"First", randomPoints(5),
		"Second", randomPoints(5))

	return p.Save(5*vg.Inch, 3*vg.Inch, "barchart.png")
}

func randomPoints(n int) plotter.XYs {
	pts := make(plotter.XYs, n)
	for i := range pts {

		pts[i].X = float64(i * 1)

		pts[i].Y = pts[i].X + 10*rand.Float64()
	}
	return pts
}
