package mim

import (
	"fmt"
	"hash"

	"golang.org/x/crypto/hkdf"
)

// Mozaic is a Hash Visualisation inspired by randoart.
type Mozaic []byte

// New creates a new Mozaic which can then be converted to ANSI or a image.
//
// A cryptographic hash function such as SHA256 or SHA512 needs to be provided.
// Via the use of a CHF the hash visualisation becomes pre-image resistant.
// Furthermore, the image output becomes fixed regardless of input, however
// based on the user's security model a 256bit CHF can be changed for a 512 CHF.
//
// It is recommended to ONLY use SHA256. This allows MIM to be standardised across
// programs.
//
// Recognised MIM configurations:
//
// MIM_SHA256, MIM_SHA512
func New(fingerprint []byte, hash func() hash.Hash) Mozaic {

	kdf := hkdf.New(hash, fingerprint, nil, nil)

	buf := make([]byte, hash().Size())
	kdf.Read(buf)

	return buf
}

func (m Mozaic) ANSI256() (output string) {
	lm := len(m)
	// rowWidth := lm / 8
	rows := lm / 4

	for i := 0; i < lm; i++ {
		if i%rows == 0 && i != 0 {
			output += "\r\n"
		} else if i%4 == 3 && i != 0 { // per row, every fourth byte add space and create parallel square
			output += colourANSI256(m[i]) + "  \x1b[0m  "
			continue
		}

		output += colourANSI256(m[i]) + "  \x1b[0m"
	}

	return output + "\x1b[0m"
}

func (m Mozaic) ANSITrueColour() (output string) {
	// pre-computed length of m
	lm := len(m)

mainloop:
	for i := 0; i < lm; i += 3 {

		// if there are less than 3 bytes left
		if x := (lm - 1) - i; x < 3 {
			if x == 0 {
				break mainloop
			}

			switch x {
			case 1:
				output += colourANSITrueColour(m[i], 0, 0) + "  "
			default:
				output += colourANSITrueColour(m[i], m[i+1], 0) + "  "
			}

			break mainloop
		}

		output += colourANSITrueColour(m[i], m[i+1], m[i+2]) + "  "
	}

	return output + "\x1b[0m"
}

// colourANSI256 takes a byte and returns a 256bit ANSI colour code
func colourANSI256(d byte) string {
	return fmt.Sprintf("\x1b[48;5;%dm", d)
}

// colourANSITrueColour takes three bytes and returns a RGB ANSI colour code
func colourANSITrueColour(r, g, b byte) string {
	return fmt.Sprintf("\x1b[38;2;%d;%d;%d;48;2;%d;%d;%dm", r, g, b, r, g, b)
}
