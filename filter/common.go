package filter

// RuneMap drops unwanted runes by returning false.  This
type RuneMap func(r rune) rune
