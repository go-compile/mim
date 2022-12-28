package mim

import (
	"fmt"
	"hash"
	"strconv"

	"golang.org/x/crypto/hkdf"
)

// Mozaic is a Hash Visualisation inspired by randoart.
type Mozaic []byte

// New creates a new Mozaic which can then be converted to ANSI or a image.
//
// A cryptographic hash function such as SHA256 or SHA512 needs to be provided.
// Via the use of a CHF the hash visualisation becomes pre-image resistant.
// Furthermore, the image output becomes fixed regardless of input at 64 bytes
//
// It is recommended to ONLY use SHA256. This allows MIM to be standardised across
// programs.
//
// Recognised MIM configurations:
//
// MIM_SHA256, MIM_SHA512
func New(fingerprint []byte, hash func() hash.Hash) Mozaic {

	kdf := hkdf.New(hash, fingerprint, nil, nil)

	// 256bit buffer
	buf := make([]byte, 32)
	kdf.Read(buf)

	return buf
}

// ANSI will provide the Hash Visualisation as ANSI escape codes
func (m Mozaic) ANSI() (output string) {
	lm := len(m)
	// rowWidth := lm / 8
	rows := lm / 4

	for i := 0; i < lm; i++ {
		if i%rows == 0 && i != 0 {
			output += "\r\n"
		} else if i%2 == 1 && i != 0 { // per row, every fourth byte add space and create parallel square
			l, r := colourANSI(m[i])
			output += l + "  " + r + "  \x1b[0m  "
			continue
		}

		l, r := colourANSI(m[i])
		output += l + "  " + r + "  \x1b[0m"
	}

	return output + "\x1b[0m"
}

// DeprecatedANSITrueColour uses TrueColour RGB to print a short bar.
// DO NOT USE IF SECURITY IS IMPORTANT!
//
// Due to the sRGB colour space, an attacker may find it easier to forge a
// similar looking Mozaic. This is because the bytes for each coloured
// square can be slightly off and our human eye will struggle to notice a
// difference. Furthermore, the colour differences from monitor to monitor
// may skew the colours, resulting in false positives or false negatives.
func (m Mozaic) DeprecatedANSITrueColour() (output string) {
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
				output += colourANSITrueColour([3]byte{m[i], 0, 0}) + "  "
			default:
				output += colourANSITrueColour([3]byte{m[i], m[i+1], 0}) + "  "
			}

			break mainloop
		}

		output += colourANSITrueColour([3]byte{m[i], m[i+1], m[i+2]}) + "  "
	}

	return output + "\x1b[0m"
}

// colourANSI256 takes a byte and returns a 256bit ANSI colour code
func colourANSI256(d byte) string {
	return fmt.Sprintf("\x1b[48;5;%dm", d)
}

// colourANSI takes a byte and returns multiple 8bit colour escape codes
func colourANSI(b byte) (string, string) {
	left, right := splitUint8(b)
	return colourANSITrueColour(colour[left]), colourANSITrueColour(colour[right])
}

// colourANSITrueColour takes three bytes and returns a RGB ANSI colour code
func colourANSITrueColour(rgb [3]byte) string {
	r := strconv.Itoa(int(rgb[0]))
	g := strconv.Itoa(int(rgb[1]))
	b := strconv.Itoa(int(rgb[2]))

	// changed from printf to use strconv for performance
	return "\x1b[38;2;" + r + ";" + g + ";" + b + ";48;2;" + r + ";" + g + ";" + b + "m"
}

// bigEndianByteOrder returns true if the system is using Big Endian byte order
func bigEndianByteOrder() bool {
	// set either: 00000001 or 10000000
	n := byte(1)

	// shift n 7 bits to the right.
	// if Big Endian 00000001 >> 7 = 00000000
	//
	// if Little Endian 10000000 >> 7 = 00000001
	//
	// if big endian a bit shift by 7 on value 1 will always
	// return 0
	return n>>7 == 0
}

// TODO: add PNG output
// TODO: add JPG output
// TODO: add HTML output
