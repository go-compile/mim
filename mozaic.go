package mim

import (
	"fmt"
)

type Mozaic []byte

// New creates a new Mozaic which can then be converted to ANSI or a image
func New(fingerprint []byte) Mozaic {
	return fingerprint
}

func (m Mozaic) ANSI256() (output string) {
	for i := 0; i < len(m); i++ {
		if i%8 == 0 && i != 0 {
			output += "\r\n"
		} else if i%4 == 3 && i != 0 { // per row, every fourth byte add space and create parallel square
			output += colourANSI256(m[i]) + "  \x1b[0m  "
			continue
		}

		output += colourANSI256(m[i]) + "  "
	}

	return output + "\x1b[0m"
}

// colourANSI256 takes a byte and returns a 256bit ANSI colour code
func colourANSI256(d byte) string {
	return fmt.Sprintf("\x1b[48;5;%dm", d)
}
