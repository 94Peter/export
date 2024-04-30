package style

const (
	ValignTop    = 1
	ValignMiddle = 2
	ValignBottom = 3
	AlignLeft    = 4
	AlignCenter  = 5
	AlignRight   = 6
)

var (
	alignMap = map[string]int{
		"left":   AlignLeft,
		"right":  AlignRight,
		"center": AlignCenter,
	}
)

type TableStyle struct {
	Header          TextBlockStyle
	ColumnWidth     float64
	RowColumnNumber int
	PageRowNumber   int
	Data            []TextBlockStyle
}

type FixRowColumnTableStyle struct {
	ChartHeader      TextBlockStyle
	RowHeader        TextBlockStyle
	ColumnHeader     TextBlockStyle
	Content          TextBlockStyle
	HeatAlertContent TextBlockStyle
	CoolAlertContent TextBlockStyle
	BlankContent     TextBlockStyle
}

type StateTableStyle struct {
	ColumnTime       TextBlockStyle
	ColumnState      TextBlockStyle
	HeaderBackground Color
	MaxRowCount      int
}

type TextStyle struct {
	Font     string
	FontSize int
	Color    Color
}

type TextBlockStyle struct {
	TextStyle
	BackGround Color
	W, H       float64
	TextAlign  string
}

func (tbs *TextBlockStyle) GetAlign() int {
	align, ok := alignMap[tbs.TextAlign]
	if !ok {
		return AlignLeft
	}
	return align
}

type SensorReportStyle struct {
	Title        TextStyle      // 報表抬頭
	SubTitle     TextStyle      // 副標題
	SectionBlock TextBlockStyle // 區塊文字
	Content      TextStyle
	TableDesc    TextStyle
	TableDescMax TextStyle
	TableDescMin TextStyle
	TableDescAvg TextStyle
	TableStyle   TableStyle
}

var (
	ColorMax = Color{
		R: 40,
		G: 116,
		B: 172,
		A: 255,
	}

	ColorMin = Color{
		R: 40,
		G: 172,
		B: 96,
		A: 255,
	}

	ColorAvg = Color{
		R: 172,
		G: 40,
		B: 50,
		A: 255,
	}

	DefaultDialySensorStyle = &SensorReportStyle{
		Title: TextStyle{
			Font:     "tw-m",
			FontSize: 16,
			Color:    ColorBlack,
		},
		SubTitle: TextStyle{
			Font:     "tw-r",
			FontSize: 12,
			Color:    ColorBlack,
		},
		SectionBlock: TextBlockStyle{
			TextStyle: TextStyle{
				Font:     "tw-m",
				FontSize: 14,
				Color:    ColorBlack,
			},
			BackGround: Color{
				R: 223,
				G: 210,
				B: 151,
			},
		},
		Content: TextStyle{
			Font:     "tw-r",
			FontSize: 12,
			Color:    ColorBlack,
		},
		TableDesc: TextStyle{
			Font:     "tw-r",
			FontSize: 8,
			Color:    ColorBlack,
		},
		TableDescMax: TextStyle{
			Font:     "tw-r",
			FontSize: 8,
			Color:    ColorMax,
		},
		TableDescMin: TextStyle{
			Font:     "tw-r",
			FontSize: 8,
			Color:    ColorMin,
		},
		TableDescAvg: TextStyle{
			Font:     "tw-r",
			FontSize: 8,
			Color:    ColorAvg,
		},
		TableStyle: TableStyle{
			Header: TextBlockStyle{
				TextStyle: TextStyle{
					Font:     "tw-r",
					FontSize: 8,
					Color:    ColorBlack,
				},
				BackGround: ColorWhite,
			},
			ColumnWidth:     111,
			RowColumnNumber: 5,
			Data: []TextBlockStyle{
				{
					TextStyle: TextStyle{
						Font:     "tw-r",
						FontSize: 8,
						Color:    ColorBlack,
					},
					W:          12,
					H:          10.0,
					BackGround: ColorWhite,
					TextAlign:  "left",
				},
				{
					TextStyle: TextStyle{
						Font:     "tw-r",
						FontSize: 8,
						Color:    ColorMax,
					},
					W:          32,
					H:          10.0,
					BackGround: ColorWhite,
					TextAlign:  "right",
				},
				{
					TextStyle: TextStyle{
						Font:     "tw-r",
						FontSize: 8,
						Color:    ColorAvg,
					},
					W:          32,
					H:          10.0,
					BackGround: ColorWhite,
					TextAlign:  "right",
				},
				{
					TextStyle: TextStyle{
						Font:     "tw-r",
						FontSize: 8,
						Color:    ColorMin,
					},
					W:          32,
					H:          10.0,
					BackGround: ColorWhite,
					TextAlign:  "right",
				},
			},
		},
	}

	DefaultWeeklySensorStyle = &SensorReportStyle{
		Title: TextStyle{
			Font:     "tw-m",
			FontSize: 16,
			Color:    ColorBlack,
		},
		SubTitle: TextStyle{
			Font:     "tw-r",
			FontSize: 12,
			Color:    ColorBlack,
		},
		SectionBlock: TextBlockStyle{
			TextStyle: TextStyle{
				Font:     "tw-m",
				FontSize: 14,
				Color:    ColorBlack,
			},
			BackGround: Color{
				R: 223,
				G: 210,
				B: 151,
			},
			TextAlign: "right",
		},
		Content: TextStyle{
			Font:     "tw-r",
			FontSize: 12,
			Color:    ColorBlack,
		},
		TableDesc: TextStyle{
			Font:     "tw-r",
			FontSize: 8,
			Color:    ColorBlack,
		},
		TableDescMax: TextStyle{
			Font:     "tw-r",
			FontSize: 8,
			Color:    ColorMax,
		},
		TableDescMin: TextStyle{
			Font:     "tw-r",
			FontSize: 8,
			Color:    ColorMin,
		},
		TableDescAvg: TextStyle{
			Font:     "tw-r",
			FontSize: 8,
			Color:    ColorAvg,
		},
		TableStyle: TableStyle{
			Header: TextBlockStyle{
				TextStyle: TextStyle{
					Font:     "tw-r",
					FontSize: 8,
					Color:    ColorBlack,
				},
				BackGround: ColorWhite,
			},
			ColumnWidth:     138,
			RowColumnNumber: 4,
			PageRowNumber:   3,
			Data: []TextBlockStyle{
				{
					TextStyle: TextStyle{
						Font:     "tw-r",
						FontSize: 8,
						Color:    ColorBlack,
					},
					W:          30,
					H:          10.0,
					BackGround: ColorWhite,
					TextAlign:  "left",
				},
				{
					TextStyle: TextStyle{
						Font:     "tw-r",
						FontSize: 8,
						Color:    ColorMax,
					},
					W:          36,
					H:          10.0,
					BackGround: ColorWhite,
					TextAlign:  "right",
				},
				{
					TextStyle: TextStyle{
						Font:     "tw-r",
						FontSize: 8,
						Color:    ColorAvg,
					},
					W:          36,
					H:          10.0,
					BackGround: ColorWhite,
					TextAlign:  "right",
				},
				{
					TextStyle: TextStyle{
						Font:     "tw-r",
						FontSize: 8,
						Color:    ColorMin,
					},
					W:          36,
					H:          10.0,
					BackGround: ColorWhite,
					TextAlign:  "right",
				},
			},
		},
	}

	DefaultMonthlySensorStyle = &SensorReportStyle{
		Title: TextStyle{
			Font:     "tw-m",
			FontSize: 16,
			Color:    ColorBlack,
		},
		SubTitle: TextStyle{
			Font:     "tw-r",
			FontSize: 12,
			Color:    ColorBlack,
		},
		SectionBlock: TextBlockStyle{
			TextStyle: TextStyle{
				Font:     "tw-m",
				FontSize: 14,
				Color:    ColorBlack,
			},
			BackGround: Color{
				R: 223,
				G: 210,
				B: 151,
			},
			TextAlign: "right",
		},
		Content: TextStyle{
			Font:     "tw-r",
			FontSize: 12,
			Color:    ColorBlack,
		},
		TableDesc: TextStyle{
			Font:     "tw-r",
			FontSize: 8,
			Color:    ColorBlack,
		},
		TableDescMax: TextStyle{
			Font:     "tw-r",
			FontSize: 8,
			Color:    ColorMax,
		},
		TableDescMin: TextStyle{
			Font:     "tw-r",
			FontSize: 8,
			Color:    ColorMin,
		},
		TableDescAvg: TextStyle{
			Font:     "tw-r",
			FontSize: 8,
			Color:    ColorAvg,
		},
		TableStyle: TableStyle{
			Header: TextBlockStyle{
				TextStyle: TextStyle{
					Font:     "tw-r",
					FontSize: 8,
					Color:    ColorBlack,
				},
				BackGround: ColorWhite,
			},
			ColumnWidth:     138,
			RowColumnNumber: 4,
			PageRowNumber:   8,
			Data: []TextBlockStyle{
				{
					TextStyle: TextStyle{
						Font:     "tw-r",
						FontSize: 8,
						Color:    ColorBlack,
					},
					W:          30,
					H:          10.0,
					BackGround: ColorWhite,
					TextAlign:  "left",
				},
				{
					TextStyle: TextStyle{
						Font:     "tw-r",
						FontSize: 8,
						Color:    ColorMax,
					},
					W:          36,
					H:          10.0,
					BackGround: ColorWhite,
					TextAlign:  "right",
				},
				{
					TextStyle: TextStyle{
						Font:     "tw-r",
						FontSize: 8,
						Color:    ColorAvg,
					},
					W:          36,
					H:          10.0,
					BackGround: ColorWhite,
					TextAlign:  "right",
				},
				{
					TextStyle: TextStyle{
						Font:     "tw-r",
						FontSize: 8,
						Color:    ColorMin,
					},
					W:          36,
					H:          10.0,
					BackGround: ColorWhite,
					TextAlign:  "right",
				},
			},
		},
	}

	DefaultYearlySensorStyle = &SensorReportStyle{
		Title: TextStyle{
			Font:     "tw-m",
			FontSize: 16,
			Color:    ColorBlack,
		},
		SubTitle: TextStyle{
			Font:     "tw-r",
			FontSize: 12,
			Color:    ColorBlack,
		},
		SectionBlock: TextBlockStyle{
			TextStyle: TextStyle{
				Font:     "tw-m",
				FontSize: 14,
				Color:    ColorBlack,
			},
			BackGround: Color{
				R: 223,
				G: 210,
				B: 151,
			},
			TextAlign: "right",
		},
		Content: TextStyle{
			Font:     "tw-r",
			FontSize: 12,
			Color:    ColorBlack,
		},
		TableDesc: TextStyle{
			Font:     "tw-r",
			FontSize: 8,
			Color:    ColorBlack,
		},
		TableDescMax: TextStyle{
			Font:     "tw-r",
			FontSize: 8,
			Color:    ColorMax,
		},
		TableDescMin: TextStyle{
			Font:     "tw-r",
			FontSize: 8,
			Color:    ColorMin,
		},
		TableDescAvg: TextStyle{
			Font:     "tw-r",
			FontSize: 8,
			Color:    ColorAvg,
		},
		TableStyle: TableStyle{
			Header: TextBlockStyle{
				TextStyle: TextStyle{
					Font:     "tw-r",
					FontSize: 8,
					Color:    ColorBlack,
				},
				BackGround: ColorWhite,
			},
			ColumnWidth:     185,
			RowColumnNumber: 3,
			PageRowNumber:   2,
			Data: []TextBlockStyle{
				{
					TextStyle: TextStyle{
						Font:     "tw-r",
						FontSize: 8,
						Color:    ColorBlack,
					},
					W:          12,
					H:          10.0,
					BackGround: ColorWhite,
					TextAlign:  "left",
				},
				{
					TextStyle: TextStyle{
						Font:     "tw-r",
						FontSize: 8,
						Color:    ColorMax,
					},
					W:          55,
					H:          10.0,
					BackGround: ColorWhite,
					TextAlign:  "right",
				},
				{
					TextStyle: TextStyle{
						Font:     "tw-r",
						FontSize: 8,
						Color:    ColorAvg,
					},
					W:          55,
					H:          10.0,
					BackGround: ColorWhite,
					TextAlign:  "right",
				},
				{
					TextStyle: TextStyle{
						Font:     "tw-r",
						FontSize: 8,
						Color:    ColorMin,
					},
					W:          55,
					H:          10.0,
					BackGround: ColorWhite,
					TextAlign:  "right",
				},
			},
		},
	}

	DefaultWeeklySensorV2Style = &SensorReportStyle{
		Title: TextStyle{
			Font:     "tw-m",
			FontSize: 16,
			Color:    ColorBlack,
		},
		SubTitle: TextStyle{
			Font:     "tw-r",
			FontSize: 12,
			Color:    ColorBlack,
		},
		SectionBlock: TextBlockStyle{
			TextStyle: TextStyle{
				Font:     "tw-m",
				FontSize: 14,
				Color:    ColorBlack,
			},
			BackGround: Color{
				R: 223,
				G: 210,
				B: 151,
			},
			TextAlign: "right",
		},
		Content: TextStyle{
			Font:     "tw-r",
			FontSize: 12,
			Color:    ColorBlack,
		},
		TableDesc: TextStyle{
			Font:     "tw-r",
			FontSize: 8,
			Color:    ColorBlack,
		},
		TableDescMax: TextStyle{
			Font:     "tw-r",
			FontSize: 8,
			Color:    ColorMax,
		},
		TableDescMin: TextStyle{
			Font:     "tw-r",
			FontSize: 8,
			Color:    ColorMin,
		},
		TableDescAvg: TextStyle{
			Font:     "tw-r",
			FontSize: 8,
			Color:    ColorAvg,
		},
		TableStyle: TableStyle{
			Header: TextBlockStyle{
				TextStyle: TextStyle{
					Font:     "tw-r",
					FontSize: 8,
					Color:    ColorBlack,
				},
				BackGround: ColorWhite,
			},
			ColumnWidth:     138,
			RowColumnNumber: 1,
			PageRowNumber:   4,
			Data: []TextBlockStyle{
				{
					TextStyle: TextStyle{
						Font:     "tw-r",
						FontSize: 10,
						Color:    ColorBlack,
					},
					W:          20,
					H:          10.0,
					BackGround: ColorWhite,
					TextAlign:  "left",
				},
				{
					TextStyle: TextStyle{
						Font:     "tw-r",
						FontSize: 10,
						Color:    ColorMax,
					},
					W:          38,
					H:          10.0,
					BackGround: ColorWhite,
					TextAlign:  "right",
				},
				{
					TextStyle: TextStyle{
						Font:     "tw-r",
						FontSize: 10,
						Color:    ColorAvg,
					},
					W:          38,
					H:          10.0,
					BackGround: ColorWhite,
					TextAlign:  "right",
				},
				{
					TextStyle: TextStyle{
						Font:     "tw-r",
						FontSize: 10,
						Color:    ColorMin,
					},
					W:          38,
					H:          10.0,
					BackGround: ColorWhite,
					TextAlign:  "right",
				},
			},
		},
	}

	DefaultDialySensorV2Style = &SensorReportStyle{
		Title: TextStyle{
			Font:     "tw-m",
			FontSize: 16,
			Color:    ColorBlack,
		},
		SubTitle: TextStyle{
			Font:     "tw-r",
			FontSize: 12,
			Color:    ColorBlack,
		},
		SectionBlock: TextBlockStyle{
			TextStyle: TextStyle{
				Font:     "tw-m",
				FontSize: 14,
				Color:    ColorBlack,
			},
			BackGround: Color{
				R: 223,
				G: 210,
				B: 151,
			},
		},
		Content: TextStyle{
			Font:     "tw-r",
			FontSize: 12,
			Color:    ColorBlack,
		},
		TableDesc: TextStyle{
			Font:     "tw-r",
			FontSize: 8,
			Color:    ColorBlack,
		},
		TableDescMax: TextStyle{
			Font:     "tw-r",
			FontSize: 8,
			Color:    ColorMax,
		},
		TableDescMin: TextStyle{
			Font:     "tw-r",
			FontSize: 8,
			Color:    ColorMin,
		},
		TableDescAvg: TextStyle{
			Font:     "tw-r",
			FontSize: 8,
			Color:    ColorAvg,
		},
		TableStyle: TableStyle{
			Header: TextBlockStyle{
				TextStyle: TextStyle{
					Font:     "tw-r",
					FontSize: 8,
					Color:    ColorBlack,
				},
				BackGround: ColorWhite,
			},
			ColumnWidth:     138,
			RowColumnNumber: 4,
			PageRowNumber:   4,
			Data: []TextBlockStyle{
				{
					TextStyle: TextStyle{
						Font:     "tw-r",
						FontSize: 10,
						Color:    ColorBlack,
					},
					W:          20,
					H:          10.0,
					BackGround: ColorWhite,
					TextAlign:  "left",
				},
				{
					TextStyle: TextStyle{
						Font:     "tw-r",
						FontSize: 10,
						Color:    ColorMax,
					},
					W:          38,
					H:          10.0,
					BackGround: ColorWhite,
					TextAlign:  "right",
				},
				{
					TextStyle: TextStyle{
						Font:     "tw-r",
						FontSize: 10,
						Color:    ColorAvg,
					},
					W:          38,
					H:          10.0,
					BackGround: ColorWhite,
					TextAlign:  "right",
				},
				{
					TextStyle: TextStyle{
						Font:     "tw-r",
						FontSize: 10,
						Color:    ColorMin,
					},
					W:          38,
					H:          10.0,
					BackGround: ColorWhite,
					TextAlign:  "right",
				},
			},
		},
	}

	DefaultYearlySensorV2Style = &SensorReportStyle{
		Title: TextStyle{
			Font:     "tw-m",
			FontSize: 16,
			Color:    ColorBlack,
		},
		SubTitle: TextStyle{
			Font:     "tw-r",
			FontSize: 12,
			Color:    ColorBlack,
		},
		SectionBlock: TextBlockStyle{
			TextStyle: TextStyle{
				Font:     "tw-m",
				FontSize: 14,
				Color:    ColorBlack,
			},
			BackGround: Color{
				R: 223,
				G: 210,
				B: 151,
			},
			TextAlign: "right",
		},
		Content: TextStyle{
			Font:     "tw-r",
			FontSize: 12,
			Color:    ColorBlack,
		},
		TableDesc: TextStyle{
			Font:     "tw-r",
			FontSize: 8,
			Color:    ColorBlack,
		},
		TableDescMax: TextStyle{
			Font:     "tw-r",
			FontSize: 8,
			Color:    ColorMax,
		},
		TableDescMin: TextStyle{
			Font:     "tw-r",
			FontSize: 8,
			Color:    ColorMin,
		},
		TableDescAvg: TextStyle{
			Font:     "tw-r",
			FontSize: 8,
			Color:    ColorAvg,
		},
		TableStyle: TableStyle{
			Header: TextBlockStyle{
				TextStyle: TextStyle{
					Font:     "tw-r",
					FontSize: 8,
					Color:    ColorBlack,
				},
				BackGround: ColorWhite,
			},
			ColumnWidth:     180,
			RowColumnNumber: 2,
			PageRowNumber:   3,
			Data: []TextBlockStyle{
				{
					TextStyle: TextStyle{
						Font:     "tw-r",
						FontSize: 10,
						Color:    ColorBlack,
					},
					W:          35,
					H:          10.0,
					BackGround: ColorWhite,
					TextAlign:  "left",
				},
				{
					TextStyle: TextStyle{
						Font:     "tw-r",
						FontSize: 10,
						Color:    ColorMax,
					},
					W:          45,
					H:          10.0,
					BackGround: ColorWhite,
					TextAlign:  "right",
				},
				{
					TextStyle: TextStyle{
						Font:     "tw-r",
						FontSize: 10,
						Color:    ColorAvg,
					},
					W:          45,
					H:          10.0,
					BackGround: ColorWhite,
					TextAlign:  "right",
				},
				{
					TextStyle: TextStyle{
						Font:     "tw-r",
						FontSize: 8,
						Color:    ColorMin,
					},
					W:          45,
					H:          10.0,
					BackGround: ColorWhite,
					TextAlign:  "right",
				},
			},
		},
	}

	DefaultMonthlySensorV2Style = &SensorReportStyle{
		Title: TextStyle{
			Font:     "tw-m",
			FontSize: 16,
			Color:    ColorBlack,
		},
		SubTitle: TextStyle{
			Font:     "tw-r",
			FontSize: 12,
			Color:    ColorBlack,
		},
		SectionBlock: TextBlockStyle{
			TextStyle: TextStyle{
				Font:     "tw-m",
				FontSize: 14,
				Color:    ColorBlack,
			},
			BackGround: Color{
				R: 223,
				G: 210,
				B: 151,
			},
			TextAlign: "right",
		},
		Content: TextStyle{
			Font:     "tw-r",
			FontSize: 12,
			Color:    ColorBlack,
		},
		TableDesc: TextStyle{
			Font:     "tw-r",
			FontSize: 8,
			Color:    ColorBlack,
		},
		TableDescMax: TextStyle{
			Font:     "tw-r",
			FontSize: 8,
			Color:    ColorMax,
		},
		TableDescMin: TextStyle{
			Font:     "tw-r",
			FontSize: 8,
			Color:    ColorMin,
		},
		TableDescAvg: TextStyle{
			Font:     "tw-r",
			FontSize: 8,
			Color:    ColorAvg,
		},
		TableStyle: TableStyle{
			Header: TextBlockStyle{
				TextStyle: TextStyle{
					Font:     "tw-r",
					FontSize: 8,
					Color:    ColorBlack,
				},
				BackGround: ColorWhite,
			},
			ColumnWidth:     180,
			RowColumnNumber: 8,
			PageRowNumber:   3,
			Data: []TextBlockStyle{
				{
					TextStyle: TextStyle{
						Font:     "tw-r",
						FontSize: 10,
						Color:    ColorBlack,
					},
					W:          20,
					H:          10.0,
					BackGround: ColorWhite,
					TextAlign:  "left",
				},
				{
					TextStyle: TextStyle{
						Font:     "tw-r",
						FontSize: 10,
						Color:    ColorMax,
					},
					W:          50,
					H:          10.0,
					BackGround: ColorWhite,
					TextAlign:  "right",
				},
				{
					TextStyle: TextStyle{
						Font:     "tw-r",
						FontSize: 10,
						Color:    ColorAvg,
					},
					W:          50,
					H:          10.0,
					BackGround: ColorWhite,
					TextAlign:  "right",
				},
				{
					TextStyle: TextStyle{
						Font:     "tw-r",
						FontSize: 10,
						Color:    ColorMin,
					},
					W:          50,
					H:          10.0,
					BackGround: ColorWhite,
					TextAlign:  "right",
				},
			},
		},
	}
)
