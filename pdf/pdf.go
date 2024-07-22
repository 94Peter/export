package pdf

import (
	"io"
	"os"

	"github.com/94peter/export/pdf/style"

	"github.com/signintech/gopdf"
)

type pdf struct {
	myPDF                                            *gopdf.GoPdf
	width, height                                    float64
	leftMargin, topMargin, rightMargin, bottomMargin float64
	page                                             uint8
}

func GetA4PDF(fontMap map[string]string, leftMargin, rightMargin, topMargin, bottomMargin float64) pdf {
	gpdf := gopdf.GoPdf{}
	// width, height := 595.28, 841.89
	pageSize := *gopdf.PageSizeA4
	gpdf.Start(gopdf.Config{PageSize: pageSize}) //595.28, 841.89 = A4
	var err error
	for key, value := range fontMap {
		err = gpdf.AddTTFFont(key, value)
		if err != nil {
			panic(err)
		}
	}
	gpdf.SetLeftMargin(leftMargin)
	gpdf.SetTopMargin(topMargin)
	return pdf{
		myPDF:        &gpdf,
		width:        pageSize.W,
		height:       pageSize.H,
		leftMargin:   leftMargin,
		rightMargin:  rightMargin,
		topMargin:    topMargin,
		bottomMargin: bottomMargin,
	}
}

func (p *pdf) WriteToFile(filepath string) error {
	pdf := p.myPDF
	return pdf.WritePdf(filepath)
}

func (p *pdf) Write(w io.Writer) error {
	pdf := p.myPDF
	_, err := pdf.WriteTo(w)
	return err
}

func (p *pdf) AddPage() {
	p.page++
	pdf := p.myPDF
	pdf.AddPage()
}

func (p *pdf) Br(h float64) {
	pdf := p.myPDF
	pdf.Br(h)
	pdf.SetX(p.leftMargin)
}

func (p *pdf) GetWidth() float64 {
	return p.width - p.leftMargin - p.rightMargin
}

func (p *pdf) Line(width float64) {
	pdf := p.myPDF
	pdf.SetLineWidth(width)
	pdf.Line(p.leftMargin, pdf.GetY(), p.width-p.rightMargin, pdf.GetY())
}

func (p *pdf) Text(text string, ts style.TextStyle, align int) {
	pdf := p.myPDF
	pdf.SetFont(ts.Font, "", ts.FontSize)
	color := ts.Color
	pdf.SetTextColor(color.R, color.G, color.B)
	pdf.SetFillColor(color.R, color.G, color.B)
	ox := pdf.GetX()
	if ox < p.leftMargin {
		ox = p.leftMargin
	}
	x := ox
	textw, _ := pdf.MeasureTextWidth(text)
	switch align {
	case style.AlignCenter:
		x = (p.width / 2) - (textw / 2)
	case style.AlignRight:
		x = p.width - textw - p.rightMargin
	}
	pdf.SetX(x)
	pdf.Cell(nil, text)
	pdf.SetX(ox + textw)
}

func (p *pdf) TwoColumnText(text1, text2 string, ts style.TextStyle) {
	pdf := p.myPDF
	pdf.SetFont(ts.Font, "", ts.FontSize)
	color := ts.Color
	pdf.SetTextColor(color.R, color.G, color.B)
	pdf.SetX(p.leftMargin)
	pdf.Cell(nil, text1)
	pdf.SetX(p.width/2 + p.leftMargin)
	pdf.Cell(nil, text2)
}

func (p *pdf) ImageReader(imageByte io.Reader) {
	//use image holder by io.Reader
	imgH2, err := gopdf.ImageHolderByReader(imageByte)
	if err != nil {
		panic(err)
	}
	pdf := p.myPDF
	pdf.ImageByHolder(imgH2, p.leftMargin, pdf.GetY(), nil)
}

func (p *pdf) Image(imagePath string) {
	//use image holder by io.Reader
	file, err := os.Open(imagePath)
	if err != nil {
		panic(err)
	}
	imgH2, err := gopdf.ImageHolderByReader(file)
	if err != nil {
		panic(err)
	}
	pdf := p.myPDF
	pdf.ImageByHolder(imgH2, p.leftMargin, pdf.GetY(), nil)
}

func (p *pdf) RectFillDrawColor(text string,
	font string,
	fontSize int,
	textColor style.Color,
	w, h float64,
	color style.Color,
	align, valign int,
) {
	p.rectColorText(text, font, fontSize, textColor, w, h, color, align, valign, "FD")
}

func (p *pdf) RectFillColor(text string,
	ts style.TextBlockStyle,
	w, h float64,
	align, valign int,
) {
	p.rectColorText(text, ts.Font, ts.FontSize, ts.Color, w, h, ts.BackGround, align, valign, "F")
}

func (p *pdf) rectColorText(text string,
	font string,
	fontSize int,
	textColor style.Color,
	w, h float64,
	color style.Color,
	align, valign int,
	rectType string,
) {
	pdf := p.myPDF
	pdf.SetLineWidth(0.1)
	pdf.SetFont(font, "", fontSize)
	pdf.SetFillColor(color.R, color.G, color.B) //setup fill color
	ox, x := pdf.GetX(), 0.0

	if ox < p.leftMargin {
		ox = p.leftMargin
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
