package main

import (
	"crypto/sha256"
	"fmt"
	"os"

	"github.com/go-compile/mim"
)

func main() {
	fingerprint := sha256.Sum256([]byte("certificate contents would typically go here"))

	fmt.Printf("Fingerprint: %X\n\n", fingerprint)

	m := mim.New(fingerprint[:], sha256.New)

	// create a file to write the svg to
	f, err := os.OpenFile("./mim.svg", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	m.SVG(f)

	f.Close()
}
