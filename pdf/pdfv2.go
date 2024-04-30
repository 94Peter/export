package pdf

import (
	"io"

	"github.com/94peter/export/pdf/style"
	"github.com/signintech/gopdf"
)

type AddPagePipe interface {
	Before(p PDF)
	After(p PDF)
}

type PDF interface {
	GetX() float64
	GetY() float64
	GetHeight() float64
	GetWidth() float64
	GetPage() uint8
	// 輸出檔案
	Write(w io.Writer) error
	// 直印頁面
	AddDirectPage(pp ...AddPagePipe)
	// 橫印頁面
	AddHorizontalPage(pp ...AddPagePipe)
	// 產生文字
	Text(text string, ts style.TextStyle, align int)
	// 指定位置產生文字
	TextWithPosition(text string, style style.TextStyle, x, y float64)
	// 二列文字說明
	TwoColumnText(text1, text2 string, ts style.TextStyle)
	// 換行
	Br(h float64)
	// 畫線
	Line(width float64)
	// 畫有顔色的綫
	LineWithColor(width float64, Color style.Color)
	// 指定 X,Y 畫線
	LineXY(width, x1, y1, x2, y2 float64)

	ImageReader(imageByte io.Reader)
	ImageReaderPosition(imageByte io.Reader, x, y float64)

	DrawSensorTable(nti *sensorTableIter, ts style.FixRowColumnTableStyle)
	DrawSensorMergeTable(nti *sensorTableIter, pageRows int, mergeRows int, ts style.FixRowColumnTableStyle, pp ...AddPagePipe)
	DrawSensorDynamicHeaderMergeTable(nti *sensorDynamicHeaderTableIter, pageRows int, mergeRows int, ts style.FixRowColumnTableStyle, pp ...AddPagePipe)

	DrawSenStateChartTable(nti *sensorTableIter, ts style.FixRowColumnTableStyle)
	DrawSenStateTable(nti *sensorTableIter, ts style.StateTableStyle, pp ...AddPagePipe)

	// 矩型填滿顏色
	RectFillColor(text string,
		ts style.TextBlockStyle,
		w, h float64,
		align, valign int,
	)
}

func NewPDFv2(fontMap map[string]string, left, right, top, bottom float64) PDF {
	gpdf := gopdf.GoPdf{}
	width, height := 595.28, 841.89
	gpdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: width, H: height}}) //595.28, 841.89 = A4
	var err error

	for key, value := range fontMap {
		err = gpdf.AddTTFFont(key, value)
		if err != nil {
			panic(err)
		}
	}
	gpdf.SetLeftMargin(left)
	gpdf.SetTopMargin(top)
	return &pdfv2{
		GoPdf:        &gpdf,
		width:        width,
		height:       height,
		leftMargin:   left,
		rightMargin:  right,
		topMargin:    top,
		bottomMargin: bottom,
	}
}

type pdfv2 struct {
	*gopdf.GoPdf

	width, height             float64
	leftMargin, topMargin     float64
	rightMargin, bottomMargin float64
	page                      uint8
}

func (p *pdfv2) Line(width float64) {
	p.SetLineWidth(width)
	p.GoPdf.Line(p.leftMargin, p.GetY(), p.width-p.rightMargin, p.GetY())
}

func (p *pdfv2) LineWithColor(width float64, Color style.Color) {
	p.SetStrokeColor(Color.R, Color.G, Color.B)
	p.SetLineWidth(width)
	p.GoPdf.Line(p.leftMargin, p.GetY(), p.width-p.rightMargin, p.GetY())
}

func (p *pdfv2) LineXY(width, x1, y1, x2, y2 float64) {
	p.SetLineWidth(width)
	p.GoPdf.Line(x1, y1, x2, y2)
}

func (p *pdfv2) GetWidth() float64 {
	return p.width - p.leftMargin - p.rightMargin
}

func (p *pdfv2) GetHeight() float64 {
	return p.height - p.topMargin - p.bottomMargin
}

func (p *pdfv2) GetPage() uint8 {
	return p.page
}

func (pdf *pdfv2) Text(text string, ts style.TextStyle, align int) {
	pdf.SetFont(ts.Font, "", ts.FontSize)
	color := ts.Color
	pdf.SetTextColor(color.R, color.G, color.B)
	pdf.SetFillColor(color.R, color.G, color.B)
	ox := pdf.GetX()
	if ox < pdf.leftMargin {
		ox = pdf.leftMargin
	}
	x := ox
	textw, _ := pdf.MeasureTextWidth(text)
	switch align {
	case style.AlignCenter:
		x = (pdf.width / 2) - (textw / 2)
	case style.AlignRight:
		x = pdf.width - textw - pdf.rightMargin
	}
	pdf.SetX(x)
	pdf.Cell(nil, text)
	pdf.SetX(ox + textw)
}

func (pdf *pdfv2) TextWithPosition(text string, style style.TextStyle, x, y float64) {
	pdf.SetFont(style.Font, "", style.FontSize)
	textw, _ := pdf.MeasureTextWidth(text)
	rightLimit := pdf.width - pdf.rightMargin - textw
	if x < pdf.leftMargin {
		x = pdf.leftMargin
	} else if x > rightLimit {
		x = rightLimit
	}
	oy := pdf.GetY()
	ox := pdf.GetX()
	pdf.SetX(x)
	pdf.SetY(y)

	color := style.Color
	pdf.SetTextColor(color.R, color.G, color.B)
	pdf.SetFillColor(color.R, color.G, color.B)

	pdf.Cell(nil, text)
	pdf.SetX(ox)
	pdf.SetY(oy)
}

func (pdf *pdfv2) TwoColumnText(text1, text2 string, ts style.TextStyle) {
	pdf.SetFont(ts.Font, "", ts.FontSize)
	color := ts.Color
	pdf.SetTextColor(color.R, color.G, color.B)
	pdf.SetFillColor(color.R, color.G, color.B)
	pdf.SetX(pdf.leftMargin)
	pdf.Cell(nil, text1)
	pdf.SetX(pdf.width/2 + pdf.leftMargin)
	pdf.Cell(nil, text2)
}

func (pdf *pdfv2) ImageReader(imageByte io.Reader) {
	//use image holder by io.Reader
	pdf.ImageReaderPosition(imageByte, pdf.leftMargin, pdf.GetY())
}

func (pdf *pdfv2) ImageReaderPosition(imageByte io.Reader, x, y float64) {
	imgH2, err := gopdf.ImageHolderByReader(imageByte)
	if err != nil {
		panic(err)
	}
	pdf.ImageByHolder(imgH2, x, y, nil)
}

func (pdf *pdfv2) Br(h float64) {
	pdf.GoPdf.Br(h)
	pdf.SetX(pdf.leftMargin)
}

func (p *pdfv2) AddDirectPage(pp ...AddPagePipe) {
	for _, t := range pp {
		t.Before(p)
	}
	p.page++
	p.GoPdf.AddPageWithOption(gopdf.PageOption{PageSize: &gopdf.Rect{W: 595.28, H: 841.89}})
	p.height = 841.89
	p.width = 595.28
	for i, t := range pp {
		t.After(p)
		if i == 0 {
			p.Br(10)
		}
	}
}

func (p *pdfv2) AddHorizontalPage(pp ...AddPagePipe) {
	for _, t := range pp {
		t.Before(p)
	}
	p.page++

	p.GoPdf.AddPageWithOption(gopdf.PageOption{PageSize: &gopdf.Rect{W: 841.89, H: 595.28}})
	p.width = 841.89
	p.height = 595.28
	for _, t := range pp {
		t.After(p)
	}
}

func (p *pdfv2) RectFillColor(text string,
	ts style.TextBlockStyle,
	w, h float64,
	align, valign int,
) {
	p.rectColorText(text, ts.Font, ts.FontSize, ts.Color, w, h, ts.BackGround, align, valign, "F")
}

func (pdf *pdfv2) rectColorText(text string,
	font string,
	fontSize int,
	textColor style.Color,
	w, h float64,
	color style.Color,
	align, valign int,
	rectType string,
) {
	pdf.SetLineWidth(0.1)
	pdf.SetFont(font, "", fontSize)
	pdf.SetFillColor(color.R, color.G, color.B) //setup fill color
	ox, x := pdf.GetX(), 0.0

	if ox < pdf.leftMargin {
		ox = pdf.leftMargin
	}
	x = ox
	pdf.RectFromUpperLeftWithStyle(x, pdf.GetY(), w, h, rectType)
	pdf.SetFillColor(0, 0, 0)
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
	oy, y := pdf.GetY(), 0.0
	if valign == style.ValignMiddle {
		y = oy + (h / 2) - (float64(fontSize) / 2)
	} else if valign == style.ValignBottom {
		y = oy + h - float64(fontSize)
	}
	pdf.SetY(y)

	pdf.SetTextColor(textColor.R, textColor.G, textColor.B)
	pdf.Cell(nil, text)
	pdf.SetY(oy)
	pdf.SetX(ox + w)
}

func (p *pdfv2) RectFillDrawColor(text string,
	font string,
	fontSize int,
	textColor style.Color,
	w, h float64,
	color style.Color,
	align, valign int,
) {
	p.rectColorText(text, font, fontSize, textColor, w, h, color, align, valign, "FD")
}

func (pdf *pdfv2) DrawSensorTable(nti *sensorTableIter, ts style.FixRowColumnTableStyle) {
	ox, x := pdf.GetX(), 0.0

	if ox < pdf.leftMargin {
		ox = pdf.leftMargin
	}
	x = ox
	pdf.SetX(x)

	//oy, y := pdf.GetY(), pdf.GetY()
	pdf.SetY(pdf.GetY())
	i := 0
	for _, h := range nti.header {
		if i == 0 {
			pdf.RectFillDrawColor(h, ts.ColumnHeader.Font, ts.ColumnHeader.FontSize, ts.ColumnHeader.Color, ts.ColumnHeader.W, 20, ts.ColumnHeader.BackGround, style.AlignCenter, style.ValignMiddle)
		} else {
			pdf.RectFillDrawColor(h, ts.RowHeader.Font, ts.RowHeader.FontSize, ts.RowHeader.Color, ts.RowHeader.W, 20, ts.RowHeader.BackGround, style.AlignCenter, style.ValignMiddle)
		}
		i++
	}
	pdf.Br(20)
	for _, r := range nti.rows {
		i = 0
		for _, h := range r {
			if i == 0 {
				pdf.RectFillDrawColor(h.Value, ts.ColumnHeader.Font, ts.ColumnHeader.FontSize, ts.ColumnHeader.Color, ts.ColumnHeader.W, 20, ts.ColumnHeader.BackGround, style.AlignCenter, style.ValignMiddle)
			} else if h.IsAlert == 1 {
				pdf.RectFillDrawColor(h.Value, ts.RowHeader.Font, ts.RowHeader.FontSize, ts.HeatAlertContent.Color, ts.RowHeader.W, 20, ts.HeatAlertContent.BackGround, style.AlignCenter, style.ValignMiddle)
			} else if h.IsAlert == -1 {
				pdf.RectFillDrawColor(h.Value, ts.RowHeader.Font, ts.RowHeader.FontSize, ts.CoolAlertContent.Color, ts.RowHeader.W, 20, ts.CoolAlertContent.BackGround, style.AlignCenter, style.ValignMiddle)
			} else if h.IsHeader {
				pdf.RectFillDrawColor(h.Value, ts.RowHeader.Font, ts.RowHeader.FontSize, ts.RowHeader.Color, ts.RowHeader.W, 20, ts.RowHeader.BackGround, style.AlignCenter, style.ValignMiddle)
			} else {
				pdf.RectFillDrawColor(h.Value, ts.RowHeader.Font, ts.RowHeader.FontSize, ts.Content.Color, ts.RowHeader.W, 20, ts.Content.BackGround, style.AlignCenter, style.ValignMiddle)
			}
			i++
		}
		pdf.Br(20)
	}
}

func (pdf *pdfv2) DrawSensorMergeTable(nti *sensorTableIter, pageRows int, mergeRows int, ts style.FixRowColumnTableStyle, pp ...AddPagePipe) {
	ox, x := pdf.GetX(), 0.0

	if ox < pdf.leftMargin {
		ox = pdf.leftMargin
	}
	x = ox
	pdf.SetX(x)
	pdf.SetY(pdf.GetY())
	const columnsCount = 12
	i, j, k := 0, 0, 0
	day := 1
	headerLen := len(nti.header)
	headerHeight := float64(mergeRows) * 20
	pageRows++
	for _, r := range nti.rows {
		if day%pageRows == 0 {
			pdf.AddDirectPage(pp...)
		}
		i, k = 0, 0
		for _, h := range r {
			if i == 0 {
				if h.Value == "" {
					break
				}
				pdf.RectFillDrawColor(h.Value, ts.ColumnHeader.Font, ts.ColumnHeader.FontSize, ts.ColumnHeader.Color, ts.ColumnHeader.W, headerHeight, ts.ColumnHeader.BackGround, style.AlignCenter, style.ValignMiddle)
				j, k = 0, 0
			} else {
				innerRowCount := j / 2
				if j%2 == 0 {
					for k = j / 2 * columnsCount; k < headerLen; k++ {
						if k/columnsCount != innerRowCount && k%columnsCount == 0 {
							break
						}
						pdf.RectFillDrawColor(nti.header[k], ts.ColumnHeader.Font, ts.ColumnHeader.FontSize, ts.ColumnHeader.Color, ts.RowHeader.W, 20, ts.ColumnHeader.BackGround, style.AlignCenter, style.ValignMiddle)
					}
					pdf.Br(20)
					pdf.SetX(pdf.leftMargin + ts.ColumnHeader.W)
					j++
				}
				if h.IsAlert == 1 {
					pdf.RectFillDrawColor(h.Value, ts.RowHeader.Font, ts.RowHeader.FontSize, ts.HeatAlertContent.Color, ts.RowHeader.W, 20, ts.HeatAlertContent.BackGround, style.AlignCenter, style.ValignMiddle)
				} else if h.IsAlert == -1 {
					pdf.RectFillDrawColor(h.Value, ts.RowHeader.Font, ts.RowHeader.FontSize, ts.CoolAlertContent.Color, ts.RowHeader.W, 20, ts.CoolAlertContent.BackGround, style.AlignCenter, style.ValignMiddle)
				} else {
					pdf.RectFillDrawColor(h.Value, ts.RowHeader.Font, ts.RowHeader.FontSize, ts.Content.Color, ts.RowHeader.W, 20, ts.Content.BackGround, style.AlignCenter, style.ValignMiddle)
				}
				if i/columnsCount != innerRowCount && i%columnsCount == 0 {
					pdf.Br(20)
					pdf.SetX(pdf.leftMargin + ts.ColumnHeader.W)
					j++
				}
			}
			i++
		}
		day++
		pdf.SetX(pdf.leftMargin)
	}
}

func (pdf *pdfv2) DrawSensorDynamicHeaderMergeTable(nti *sensorDynamicHeaderTableIter, pageRows int, mergeRows int, ts style.FixRowColumnTableStyle, pp ...AddPagePipe) {
	ox, x := pdf.GetX(), 0.0

	if ox < pdf.leftMargin {
		ox = pdf.leftMargin
	}
	x = ox
	pdf.SetX(x)
	pdf.SetY(pdf.GetY())
	const columnsCount = 12
	i, j, k := 0, 0, 0
	day := 1
	var header []string
	var headerLen int

	headerHeight := float64(mergeRows) * 20
	pageRows++
	addPage := true
	for _, r := range nti.rows {
		header = nti.header[day-1]
		headerLen = len(header)
		if day%pageRows == 0 && addPage {
			pdf.AddDirectPage(pp...)
		}
		i, k = 0, 0
		for index, h := range r {
			if i == 0 {
				if r[index+1].Value == "-" {
					addPage = false
					break
				}
				pdf.RectFillDrawColor(h.Value, ts.ColumnHeader.Font, ts.ColumnHeader.FontSize, ts.ColumnHeader.Color, ts.ColumnHeader.W, headerHeight, ts.ColumnHeader.BackGround, style.AlignCenter, style.ValignMiddle)
				j, k = 0, 0
			} else {
				innerRowCount := j / 2
				if j%2 == 0 {
					for k = j / 2 * columnsCount; k < headerLen; k++ {
						if k/columnsCount != innerRowCount && k%columnsCount == 0 {
							break
						}
						pdf.RectFillDrawColor(header[k], ts.ColumnHeader.Font, ts.ColumnHeader.FontSize, ts.ColumnHeader.Color, ts.RowHeader.W, 20, ts.ColumnHeader.BackGround, style.AlignCenter, style.ValignMiddle)
					}
					pdf.Br(20)
					pdf.SetX(pdf.leftMargin + ts.ColumnHeader.W)
					j++
				}
				if h.IsAlert == 1 {
					pdf.RectFillDrawColor(h.Value, ts.RowHeader.Font, ts.RowHeader.FontSize, ts.HeatAlertContent.Color, ts.RowHeader.W, 20, ts.HeatAlertContent.BackGround, style.AlignCenter, style.ValignMiddle)
				} else if h.IsAlert == -1 {
					pdf.RectFillDrawColor(h.Value, ts.RowHeader.Font, ts.RowHeader.FontSize, ts.CoolAlertContent.Color, ts.RowHeader.W, 20, ts.CoolAlertContent.BackGround, style.AlignCenter, style.ValignMiddle)
				} else {
					pdf.RectFillDrawColor(h.Value, ts.RowHeader.Font, ts.RowHeader.FontSize, ts.Content.Color, ts.RowHeader.W, 20, ts.Content.BackGround, style.AlignCenter, style.ValignMiddle)
				}
				if i/columnsCount != innerRowCount && i%columnsCount == 0 {
					pdf.Br(20)
					pdf.SetX(pdf.leftMargin + ts.ColumnHeader.W)
					j++
				}
			}
			i++
		}
		day++
		pdf.SetX(pdf.leftMargin)
	}
}

func (pdf *pdfv2) DrawSenStateChartTable(nti *sensorTableIter, ts style.FixRowColumnTableStyle) {
	ox, x := pdf.GetX(), 0.0

	if ox < pdf.leftMargin {
		ox = pdf.leftMargin
	}
	x = ox
	pdf.SetX(x)

	pdf.SetY(pdf.GetY())

	i := 0
	for n, r := range nti.rows {

		if n%4 == 0 {
			pdf.Br(2)
			i := 0
			for _, h := range nti.header {
				if i == 0 {
					pdf.RectFillColor("", ts.ChartHeader, ts.ColumnHeader.W-12, 20, style.AlignCenter, style.ValignBottom)
				} else {
					pdf.RectFillColor(h, ts.ChartHeader, ts.RowHeader.W, 20, style.AlignLeft, style.ValignBottom)
				}
				i++
			}
			pdf.RectFillColor("60分", ts.ChartHeader, ts.RowHeader.W, 20, style.AlignLeft, style.ValignBottom)
			pdf.Br(22)
			l := float64(len(nti.header))
			x, y := pdf.GetX()+ts.ColumnHeader.W, pdf.GetY()
			for i := 0.0; i <= l; i++ {
				pdf.LineXY(1, x, y, x, y+10)
				x = x + ts.RowHeader.W
			}
			pdf.Br(12)
		}

		i = 0
		for _, h := range r {
			if i == 0 {
				pdf.RectFillDrawColor(h.Value, ts.ColumnHeader.Font, ts.ColumnHeader.FontSize, ts.ColumnHeader.Color, ts.ColumnHeader.W, 20, ts.ColumnHeader.BackGround, style.AlignCenter, style.ValignMiddle)
			} else if h.IsAlert != 0 {
				pdf.RectFillDrawColor("", ts.RowHeader.Font, ts.RowHeader.FontSize, ts.HeatAlertContent.Color, ts.RowHeader.W/5, 20, ts.HeatAlertContent.BackGround, style.AlignCenter, style.ValignMiddle)
			} else if h.Value == "-" {
				pdf.RectFillDrawColor("", ts.RowHeader.Font, ts.RowHeader.FontSize, ts.HeatAlertContent.Color, ts.RowHeader.W/5, 20, ts.BlankContent.BackGround, style.AlignCenter, style.ValignMiddle)
			} else {
				pdf.RectFillDrawColor("", ts.RowHeader.Font, ts.RowHeader.FontSize, ts.Content.Color, ts.RowHeader.W/5, 20, ts.Content.BackGround, style.AlignCenter, style.ValignMiddle)
			}
			i++
		}
		pdf.Br(20)
	}
}

func (pdf *pdfv2) DrawSenStateTable(nti *sensorTableIter, ts style.StateTableStyle, pp ...AddPagePipe) {

	ox, x := pdf.GetX(), 0.0

	if ox < pdf.leftMargin {
		ox = pdf.leftMargin
	}
	x = ox
	pdf.SetX(x)

	//oy, y := pdf.GetY(), pdf.GetY()
	pdf.SetY(pdf.GetY())
	i := 0
	drawTableHeader := func() {
		for _, h := range nti.header {
			if i%2 == 0 {
				pdf.RectFillDrawColor(h, ts.ColumnTime.Font, ts.ColumnTime.FontSize, ts.ColumnTime.Color, ts.ColumnTime.W, 20, ts.HeaderBackground, style.AlignCenter, style.ValignMiddle)
			} else {
				pdf.RectFillDrawColor(h, ts.ColumnState.Font, ts.ColumnState.FontSize, ts.ColumnState.Color, ts.ColumnState.W, 20, ts.HeaderBackground, style.AlignCenter, style.ValignMiddle)
			}
			i++
		}
		pdf.Br(20)
	}
	drawTableHeader()

	maxRows := ts.MaxRowCount
	rowCount := 0
	for _, r := range nti.rows {
		if rowCount != 0 && rowCount%maxRows == 0 {
			pdf.AddDirectPage(pp...)
			drawTableHeader()
		}
		i = 0
		for _, h := range r {
			if i%2 == 0 {
				pdf.RectFillDrawColor(h.Value, ts.ColumnTime.Font, ts.ColumnTime.FontSize, ts.ColumnTime.Color, ts.ColumnTime.W, 20, ts.ColumnTime.BackGround, style.AlignCenter, style.ValignMiddle)
			} else {
				pdf.RectFillDrawColor(h.Value, ts.ColumnState.Font, ts.ColumnState.FontSize, ts.ColumnState.Color, ts.ColumnState.W, 20, ts.ColumnState.BackGround, style.AlignCenter, style.ValignMiddle)
			}
			i++
		}
		pdf.Br(20)
		rowCount++
	}
}
