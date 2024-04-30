package pdf

import (
	"fmt"

	"github.com/94peter/export/pdf/style"
)

type TableValueAryIter struct {
	index   int
	dataLen int
	Header  string
	data    []TimeValueAry
}

func GetTableValueAryIter(header string, data []TimeValueAry) *TableValueAryIter {
	return &TableValueAryIter{
		index:   0,
		dataLen: len(data),
		Header:  header,
		data:    data,
	}
}

func (vdai *TableValueAryIter) Next(tva *TimeValueAry) bool {
	if vdai.index >= vdai.dataLen {
		return false
	}
	*tva = vdai.data[vdai.index]
	vdai.index++
	return true
}

type TimeValueAry struct {
	Time   string
	Values []float64
}

func (tva *TimeValueAry) toTextAry() []string {
	result := []string{tva.Time}
	for _, f := range tva.Values {
		result = append(result, fmt.Sprintf("%.2f", f))
	}
	return result
}

type columnData struct {
	TextAry  []string
	StyleAry []style.TextBlockStyle
}

func (p *pdf) TimeValueTable(data []*TableValueAryIter, ts style.TableStyle) []*TableValueAryIter {
	splitSize := ts.RowColumnNumber
	drawData := data
	if dLen := len(drawData); dLen > splitSize {
		drawData = data[:splitSize]

	} else {
		splitSize = dLen
	}

	timveValues := TimeValueAry{}
	rowNumber := 0
	for drawData != nil {
		rowNumber++
		columnLen := len(drawData)
		var tvdata *TableValueAryIter
		maxNumber := 1
		for i := 0; i < columnLen; i++ {
			tvdata = drawData[i]
			if dataLen := len(tvdata.data); dataLen > maxNumber {
				maxNumber = dataLen
			}
			p.RectFillDrawColor(tvdata.Header, ts.Header.Font, ts.Header.FontSize, ts.Header.Color, ts.ColumnWidth, 20, ts.Header.BackGround, style.AlignCenter, style.ValignMiddle)
		}
		p.Br(20)
		h := float64(ts.Data[0].FontSize) * float64(maxNumber)
		pdf := p.myPDF
		oy := pdf.GetY()
		for i := 0; i < columnLen; i++ {
			tvdata = drawData[i]
			p.DrawColumn(ts.ColumnWidth, h, style.ColorWhite, "FD")
			for tvdata.Next(&timveValues) {
				p.TextValuesAry(
					columnData{
						TextAry:  timveValues.toTextAry(),
						StyleAry: ts.Data,
					},
					style.ValignTop,
				)
			}
			pdf.SetX(pdf.GetX() + ts.ColumnWidth)
			pdf.SetY(oy)
		}
		p.Br(h)
		data = data[splitSize:]
		if len(data) == 0 {
			break
		} else if len(data) < 5 {
			splitSize = len(data)
		}
		if rowNumber == ts.PageRowNumber {
			return data
		}
		drawData = data[:splitSize]
	}
	return nil
}

func (p *pdf) DrawColumn(w, h float64, color style.Color, rectType string) {
	pdf := p.myPDF
	pdf.SetFillColor(color.R, color.G, color.B)
	pdf.RectFromUpperLeftWithStyle(pdf.GetX(), pdf.GetY(), w, h, rectType)
}

// 文字區塊4個數字
func (p *pdf) TextValuesAry(rows columnData, valign int) {

	pdf := p.myPDF
	ox, x := pdf.GetX(), 0.0

	if ox < p.leftMargin {
		ox = p.leftMargin
	}
	x = ox

	oy, y := pdf.GetY(), pdf.GetY()

	maxFontSize := 0.0
	fs := 0.0
	for i := 0; i < len(rows.TextAry); i++ {
		pdf.SetX(x)
		text := rows.TextAry[i]
		tStyle := rows.StyleAry[i]
		fs = float64(tStyle.TextStyle.FontSize)
		if fs > maxFontSize {
			maxFontSize = fs
		}
		x = p.tableText(text, fs, tStyle.Color, x, y, tStyle.GetAlign(), valign, tStyle.W, tStyle.H)
	}
	y = oy + float64(maxFontSize)
	pdf.SetY(y)
	pdf.SetX(ox)
}

func (p *pdf) tableText(text string, floatFontSize float64, color style.Color, x, y float64, align, valign int, w, h float64) (endX float64) {
	ox := x
	pdf := p.myPDF
	pdf.SetFillColor(color.R, color.G, color.B)
	if align == style.AlignCenter {
		textw, _ := pdf.MeasureTextWidth(text)
		x = x + (w / 2) - (textw / 2)
	} else if align == style.AlignRight {
		textw, _ := pdf.MeasureTextWidth(text)
		x = x + w - textw
	} else {
		x = x + 5
	}

	pdf.SetX(x)

	if valign == style.ValignMiddle {
		y = y + (h / 2) - (floatFontSize / 2)
	} else if valign == style.ValignBottom {
		y = y + h - floatFontSize
	}
	pdf.SetY(y)
	pdf.Cell(nil, text)
	endX = ox + w
	return
}

type normalTableIter struct {
	index   int
	dataLen int
	header  []string
	rows    [][]string
}

func GetNormalTableIter(h []string) *normalTableIter {
	return &normalTableIter{
		index:   0,
		dataLen: 0,
		header:  h,
	}
}

func (nti *normalTableIter) AddRow(r []string) {
	nti.rows = append(nti.rows, r)
}

func (nti *normalTableIter) DrawPDF(p *pdf, ts style.TableStyle) {

	pdf := p.myPDF
	ox, x := pdf.GetX(), 0.0

	if ox < p.leftMargin {
		ox = p.leftMargin
	}
	x = ox
	pdf.SetX(x)

	//oy, y := pdf.GetY(), pdf.GetY()
	pdf.SetY(pdf.GetY())

	for _, h := range nti.header {
		p.RectFillDrawColor(h, ts.Header.Font, ts.Header.FontSize, ts.Header.Color, ts.ColumnWidth, 20, ts.Header.BackGround, style.AlignCenter, style.ValignMiddle)
	}
	p.Br(20)
	for _, r := range nti.rows {
		for _, h := range r {
			p.RectFillDrawColor(h, ts.Header.Font, ts.Header.FontSize, ts.Header.Color, ts.ColumnWidth, 20, ts.Header.BackGround, style.AlignCenter, style.ValignMiddle)
		}
		p.Br(20)
	}

}

func (p *pdf) TimeValueTableV2(data []*TableValueAryIter, ts style.TableStyle) []*TableValueAryIter {
	splitSize := ts.RowColumnNumber
	drawData := data
	if dLen := len(drawData); dLen > splitSize {
		drawData = data[:splitSize]
	} else {
		splitSize = dLen
	}

	timveValues := TimeValueAry{}
	rowNumber := 0
	pdf := p.myPDF
	oy := pdf.GetY()
	maxNumber := 1
	dataL := len(data)
	for i := 0; i < dataL; i++ {
		if dataLen := len(data[i].data); dataLen > maxNumber {
			maxNumber = dataLen
		}
	}
	for drawData != nil {
		rowNumber++
		columnLen := len(drawData)
		var tvdata *TableValueAryIter

		ox := pdf.GetX()
		for i := 0; i < columnLen; i++ {
			tvdata = drawData[i]

			if i == 0 {
				h := float64(ts.Data[0].FontSize) * (float64(maxNumber) + 1.5) * float64(columnLen)
				p.DrawColumn(ts.ColumnWidth, h, style.ColorWhite, "FD")
			}
			pdf.SetY(pdf.GetY() + 5)
			//p.Br(5)
			p.Text(tvdata.Header, ts.Header.TextStyle, style.AlignLeft)
			pdf.SetY(pdf.GetY() + 10)
			pdf.SetX(ox)
			for tvdata.Next(&timveValues) {
				p.TextValuesAry(
					columnData{
						TextAry:  timveValues.toTextAry(),
						StyleAry: ts.Data,
					},
					style.ValignTop,
				)
			}
		}
		pdf.SetX(pdf.GetX() + ts.ColumnWidth)
		pdf.SetY(oy)
		data = data[splitSize:]
		if len(data) == 0 {
			break
		} else if len(data) < ts.RowColumnNumber {
			splitSize = len(data)
		}
		if rowNumber == ts.PageRowNumber {
			return data
		}
		drawData = data[:splitSize]
	}
	return nil
}

type SensorCell struct {
	Value   string
	IsAlert int8
	// if overhigh, then 1;
	// if overlow, then -1;
	// if normal, then 0;
	IsHeader bool
}

type sensorTableIter struct {
	index   int
	dataLen int
	header  []string
	rows    [][]SensorCell
}

func GetSensorTableIter(h []string) *sensorTableIter {
	return &sensorTableIter{
		index:   0,
		dataLen: 0,
		header:  h,
	}
}

func (nti *sensorTableIter) AddRow(r []SensorCell) {
	nti.rows = append(nti.rows, r)
}

type sensorDynamicHeaderTableIter struct {
	index   int
	dataLen int
	header  [][]string
	rows    [][]SensorCell
}

func GetSensorDynamicHeaderTableIter(h [][]string) *sensorDynamicHeaderTableIter {
	return &sensorDynamicHeaderTableIter{
		index:   0,
		dataLen: 0,
		header:  h,
	}
}

func (nti *sensorDynamicHeaderTableIter) AddRow(r []SensorCell) {
	nti.rows = append(nti.rows, r)
}
