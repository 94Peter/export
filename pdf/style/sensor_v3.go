package style

type SensorV3ReportStyle struct {
	Header          TextStyle
	Title           TextStyle      // 報表抬頭
	SubTitle        TextStyle      // 副標題
	SectionBlock    TextBlockStyle // 區塊文字
	Content         TextStyle
	TableDesc       TextStyle
	TableDescMax    TextStyle
	TableDescMin    TextStyle
	TableDescAvg    TextStyle
	StateTableStyle StateTableStyle
	TableStyle      FixRowColumnTableStyle
	SenColumn	TextStyle
	SenColumnLine 	Color
	Page		TextStyle
}

var (
	SensorV3Style = SensorV3ReportStyle{
		Header: TextStyle{
			Font:     "tw-r",
			FontSize: 20,
			Color:    ColorBlack,
		},
		Title: TextStyle{
			Font:     "tw-m",
			FontSize: 16,
			Color:    ColorBlack,
		},
		SubTitle: TextStyle{
			Font:     "tw-r",
			FontSize: 9,
			Color:    Color{
				R: 104,
				G: 104,
				B: 104,
			},
		},
		Page: TextStyle{
			Font:     "tw-r",
			FontSize: 12,
			Color:    Color{
				R: 170,
				G: 170,
				B: 170,
			},
		},
		SectionBlock: TextBlockStyle{
			TextStyle: TextStyle{
				Font:     "tw-m",
				FontSize: 14,
				Color:    ColorWhite,
			},
			BackGround: Color{
				R: 99,
				G: 147,
				B: 141,
			},
		},
		SenColumn: TextStyle{
			Font:     "tw-r",
			FontSize: 14,
			Color:    Color{
				R: 80,
				G: 124,
				B: 118,
			},
		},
		SenColumnLine: ColorSenColumnLine,
		Content: TextStyle{
			Font:     "tw-r",
			FontSize: 12,
			Color:    Color{
				R: 38,
				G: 38,
				B: 38,
			},
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
		StateTableStyle: StateTableStyle{
			MaxRowCount:      35,
			HeaderBackground: ColorGray,
			ColumnState: TextBlockStyle{
				TextStyle: TextStyle{
					Font:     "tw-r",
					FontSize: 10,
					Color:    ColorBlack,
				},
				W:          85,
				BackGround: ColorWhite,
			},
			ColumnTime: TextBlockStyle{
				TextStyle: TextStyle{
					Font:     "tw-r",
					FontSize: 10,
					Color:    ColorBlack,
				},
				W:          100,
				BackGround: ColorWhite,
			},
		},
		TableStyle: FixRowColumnTableStyle{
			ChartHeader: TextBlockStyle{
				TextStyle: TextStyle{
					Font:     "tw-r",
					FontSize: 10,
					Color:    ColorBlack,
				},
				W:          40,
				BackGround: ColorWhite,
			},
			RowHeader: TextBlockStyle{
				TextStyle: TextStyle{
					Font:     "tw-r",
					FontSize: 8,
					Color:    ColorBlack,
				},
				W:          40,
				BackGround: ColorGray,
			},
			ColumnHeader: TextBlockStyle{
				TextStyle: TextStyle{
					Font:     "tw-r",
					FontSize: 8,
					Color:    ColorBlack,
				},
				W:          75,
				BackGround: ColorGray,
			},
			Content: TextBlockStyle{
				TextStyle: TextStyle{
					Font:     "tw-r",
					FontSize: 8,
					Color:    ColorBlack,
				},
				W:          75,
				BackGround: ColorWhite,
			},
			HeatAlertContent: TextBlockStyle{
				TextStyle: TextStyle{
					Font:     "tw-r",
					FontSize: 8,
					Color:    ColorBlack,
				},
				W:          75,
				BackGround: ColorHeatAlert,
			},
			CoolAlertContent: TextBlockStyle{
				TextStyle: TextStyle{
					Font:     "tw-r",
					FontSize: 8,
					Color:    ColorBlack,
				},
				W:          75,
				BackGround: ColorCoolAlert,
			},
			BlankContent: TextBlockStyle{
				TextStyle: TextStyle{
					Font:     "tw-r",
					FontSize: 8,
					Color:    ColorBlack,
				},
				W:          75,
				BackGround: ColorGray,
			},
		},
	}
)
