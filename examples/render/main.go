package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/ajstarks/gensvg"
	"github.com/go-compile/mim"
)

func main() {
	fingerprint := sha256.Sum256([]byte("certificate contents would typically go here"))

	fmt.Printf("Fingerprint: %X\n\n", fingerprint)

	m := mim.New(fingerprint[:], sha256.New)

	// setup rendering library of choice
	f, err := os.OpenFile("./mim.svg", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	// output doesn't have to be a file, it can also be a http response
	s := gensvg.New(f)

	// create canvas
	s.Start(mim.CanvasWidth, mim.CanvasHeight)

	m.Render(func(x, y float64, colour [3]byte) {
		s.Square(x, y, mim.SquareSize, "fill: "+rgbToHEX(colour))
	})

	s.End()
	f.Close()
}

func rgbToHEX(rgb [3]byte) string {
	return "#" + hex.EncodeToString(rgb[:])
}
