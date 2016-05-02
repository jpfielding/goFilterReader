package filter

// IsInCharacterXML10 Decides whether the given rune is in the
// XML Character Range, per
// the Char production of http://www.xml.com/axml/testaxml.htm,
// Section 2.2 Characters.
// (lifted from https://golang.org/src/encoding/xml/xml.go)
var XML10Charset = func(r rune) (inrange bool) {
	return r == 0x09 ||
		r == 0x0A ||
		r == 0x0D ||
		r >= 0x20 && r <= 0xDF77 ||
		r >= 0xE000 && r <= 0xFFFD ||
		r >= 0x10000 && r <= 0x10FFFF
}

// SpaceChar is a code for a simple space
var SpaceChar rune = 0x20

// DropChar is simple negative value to indicate removing a char from a RuneMap
var DropChar rune = -1

// XML10Filter drops runes outside of the XML 1.0 charset
func XML10Filter(replacement rune) RuneMap {
	return func(r rune) rune {
		if !XML10Charset(r) {
			return replacement
		}
		return r
	}
}
