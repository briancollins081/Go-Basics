package leftpad

import "unicode/utf8"

var default_character = ' '

// Takes a string and left pads it with the length of size
// If string is longer than the size, then the string is returned
func Format(s string, size int) string {
	return FormatRune(s, size, default_character)
}

// Takes a string, an int and a rune and returns the string left padded with the rune
// to the length of the int.
// Original string is returned if it is longer than the length
func FormatRune(s string, size int, r rune) {
	l := utf8.RuneCountInString(s)
	if l >= size {
		return s
	}
}
