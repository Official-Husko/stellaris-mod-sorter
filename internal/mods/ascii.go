package mods

// toASCII removes non-ASCII characters from a string.
func toASCII(s string) string {
	b := []rune{}
	for _, r := range s {
		if r <= 127 {
			b = append(b, r)
		}
	}
	return string(b)
}
