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
