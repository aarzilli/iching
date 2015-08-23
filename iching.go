package iching

import "strconv"

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

func Ichingtoi(in string) (uint64, error) {
	r := uint64(0)
	for _, ch := range []rune(in) {
		idx := int(ch - 19904)
		if idx < 0 || idx > len(invtable) {
			return 0, &strconv.NumError{"iching.Ichingtoi", in, strconv.ErrSyntax}
		}
		r = r << 6
		r += invtable[idx]
	}
	return r, nil
}

func QianPad(in string, sz int) string {
	return pad(in, sz, 0x4dc0)
}

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
