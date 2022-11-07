/*
I Ching or "Classic of Change" is a Chinese book about telling the future or something. It doesn't matter.

What matters is that I Ching defines a set of 64 "figures", called hexagrams, each consisting of 6 stacked horizontal lines, where each line is either a solid, unbroken line (Yang) or a line with a single gap (Yin).

Each hexagram can therefore be easily used to encode a number up to 6bits of entropy, this library implements a function (`Itoiching`) to convert an arbitrary 64bit unsigned integer into its I Ching representation and a function (`Ichingtoi`) to do the reverse.

Each I Ching character represents a 6 bit value, strings are ordered with the most significant 6 bits of the number first. Withing a I Ching character the top line represents the most significant bit of the sextet.
The unbroken line (Yand) corresponds to 0, the broken line corresponds to 1. We do not use the King Wen sequence.
*/
package iching

import "strconv"

// Converts in to its I Ching encoding
func Itoiching(in uint64) string {
	s := make([]rune, 0, 64/6+1)
	for {
		if in == 0 {
			break
		}

		s = append(s, table[in&077])
		in = in >> 6
	}

	if len(s) == 0 {
		s = append(s, table[0])
	}

	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}

	return string(s)
}

// Parses a number encoded as I Ching characters
func Ichingtoi(in string) (uint64, error) {
	r := uint64(0)
	for _, ch := range []rune(in) {
		idx := int(ch - 19904)
		if idx < 0 || idx >= len(invtable) {
			return 0, &strconv.NumError{"iching.Ichingtoi", in, strconv.ErrSyntax}
		}
		r = r << 6
		r += invtable[idx]
	}
	return r, nil
}

// Pads a string with Qián characters (the character that represents 0 in our encoding) up to a total sz runes
func QianPad(in string, sz int) string {
	return pad(in, sz, 0x4dc0)
}

// Pads a string with space characters up to a total of sz runes
func SpacePad(in string, sz int) string {
	return pad(in, sz, 0x20)
}

func pad(in string, sz int, padch rune) string {
	rin := []rune(in)
	if len(rin) >= sz {
		return in
	}

	r := make([]rune, 0, sz)

	for i := 0; i < sz-len(rin); i++ {
		r = append(r, padch)
	}

	for _, ch := range rin {
		r = append(r, ch)
	}

	return string(r)
}
