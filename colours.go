package mim

// colour is a list of colours based on the Macos terminal app colours.
// Colours taken from: https://wikiless.org/wiki/ANSI_escape_code?lang=en
//
// This set of colours was picked for the diversity of colours used, which
// reduces the risk of one colour being mistaken for another.
var colour = [][3]byte{
	{0, 0, 0},       // black
	{194, 54, 33},   // Red
	{37, 188, 36},   // Green
	{173, 173, 39},  // Yellow
	{73, 46, 225},   // Blue
	{211, 56, 211},  // Magenta
	{51, 187, 200},  // Cyan
	{203, 204, 205}, // White
	{129, 131, 131}, // Gray
	{252, 57, 31},   // Bright Red
	{49, 231, 34},   // Bright Green
	{234, 236, 35},  // Bright Yellow
	{88, 51, 255},   // Bright Blue
	{249, 53, 248},  // Bright Magenta
	{20, 240, 240},  // Bright Cyan
	{233, 235, 235}, // Bright Whites
}

// splitUint8 is used to split a uint8 into two "uint4"s.
//
// The variable is corrected at initiation time if the current system is
// not Little Endian.
//
// Finding out the system's byte order is important as it allows for
// a byte to be split correctly and output the appropriate int >= 16.
// If the wrong byte order was used, ints could return bigger than 16,
// thus breaking our colour codes.
var splitUint8 func(b byte) (byte, byte) = expandLittleEndian

func init() {
	if bigEndianByteOrder() {
		splitUint8 = expandBigEndian
		return
	}

}

// expandBigEndian takes one uint8 and expands it into
// two "uint4"
func expandBigEndian(b byte) (byte, byte) {
	// bit shift 4 bits to the right
	left := b >> 4
	// bit shift 4 to the left giving us the last 4 bits
	// then move back 4 bytes to become Big Endian format
	right := (b << 4) >> 4

	return left, right
}

// expandLittleEndian takes one uint8 and expands it into
// two "uint4"
func expandLittleEndian(b byte) (byte, byte) {
	// bit shift 4 bits to the right then back to clear
	// the right most 4 bits
	left := (b >> 4) << 4
	// bit shift 4 to the left giving us the last 4 bits
	// in Little Endian format
	right := b << 4

	return left, right
}
