package style

type Color struct {
	R, G, B, A uint8
}

var (
	ColorSenColumnLine = Color{
		R: 153,
		G: 186,
		B: 174,
	}

	ColorTableLine = Color{
		R: 0,
		G: 0,
		B: 0,
	}

	ColorGray = Color{
		R: 211,
		G: 211,
		B: 211,
	}

	ColorHeatAlert = Color{
		R: 255,
		G: 179,
		B: 167,
	}
	ColorCoolAlert = Color{
		R: 152,
		G: 185,
		B: 255,
	}

	ColorWhite = Color{
		R: 255,
		G: 255,
		B: 255,
	}

	ColorBlack = ColorTableLine
)
