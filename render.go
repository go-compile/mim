package mim

import (
	"encoding/hex"
	"io"

	"github.com/ajstarks/gensvg"
)

const (
	// SquareSize is used to render the mozaic
	SquareSize = float64(20)

	// CanvasWidth is used to create a SVG or PNG canvas for the mozaic
	CanvasWidth = 380
	// CanvasHeight is used to create a SVG or PNG canvas for the mozaic
	CanvasHeight = 80
)

// SVG will produce a SVG mozaic the size of 380 by 80px
func (m Mozaic) SVG(w io.Writer) {
	s := gensvg.New(w)
	s.Start(CanvasWidth, CanvasHeight)

	m.Render(func(x, y float64, colour [3]byte) {
		s.Square(x, y, 20, "fill: "+rgbToHEX(colour))
	})

	s.End()
}

// Render allows you to provide a custom plotting function which
// can then be used to produce your own image using any render engine.
func (m Mozaic) Render(plot func(x, y float64, colour [3]byte)) {
	lm := len(m)
	// rowWidth := lm / 8
	rows := lm / 4

	yOffset := float64(0)
	xOffset := float64(0)

	for i := 0; i < lm; i++ {
		if i%rows == 0 && i != 0 {
			// create new row
			yOffset += SquareSize
			xOffset = 0
		} else if i%2 == 1 && i != 0 { // per row, every fourth byte add space and create parallel square
			l, r := splitUint8(m[i])

			plot(xOffset, yOffset, colour[l])
			xOffset += SquareSize
			plot(xOffset+SquareSize, yOffset, colour[r])

			xOffset += SquareSize * 3

			continue

		}

		l, r := splitUint8(m[i])

		plot(xOffset, yOffset, colour[l])
		xOffset += SquareSize
		plot(xOffset+SquareSize, yOffset, colour[r])
	}
}

func rgbToHEX(rgb [3]byte) string {
	return "#" + hex.EncodeToString(rgb[:])
}
